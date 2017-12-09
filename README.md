# Squirrel

Squirrel is system similar in many ways to ansible with functionality except it
has been designed to closely communicate with the backend service that you use
for provisioning. Right now that is just collins (tumblr/collins) which is what
I use but I will at least try to keep this API sufficiently abstracted to make
switching at a later date "easy".

## Why would you do this?

This is mainly a test of mine to see if I can sufficiently seperate the processes 
of laying down the base OS, hardware configuration, system configuration, and 
application configuration. My hope is to be able to easily port your applications
into a VM environment such as AWS with little work although some conditional logic
will still be needed since for instance you don't configure IPs in AWS.

This is so we can easily setup AWS with all utilities that we need in a typical
DC such as local puppet servers, DNS, and any static asset servers or caches. 
It also aims to let you "burst" into AWS on the scale of minutes.

## Overview

To get a high level idea of what this repo will do lets take a look at an example
plan file and run through it step by step.

```yaml
---

name: Basic bootstrap plan
details: |
  This plan is a basic example that will go through
  configuration of your network and setting up puppet
  to run on the system.

backend: collins
logging: info
codenames:
  - xenial

plan:
  configure_network:
    template:
      src: interfaces.template
      dest: /etc/network/interfaces
      vars:
        iface: eth0
        ip: {{ public_ip }}
  setup_puppet:
    git:
      src: git@github.com:michaeljs1990/puppet.git
      dest: /var/run/puppet_repo
  run_puppet:
    command:
      exec: /usr/bin/puppet agent -t
      retry: 6

success: |
  {{ asset }} has finished provisioning and is ready for use.

failed: |
  {{ asset }} has failed to provision.
```

To start name and details will just be informative information for anyone running
your plan. Backend here is specified as collins which will be what squirrel fetches
configuration variables from such as `public_ip` and `asset` which you see example
use of later in the plan. Logging will set the level at which squirrel will send
messages to your provisioning stack if supported. The different levels are info,
debug, critical, none. Codename can be used to limit what OS this plan is able to
be run on. 

The plan will be walked through in serial and the first level of indentation such
as `configure_network` is only used for logging purposes. template under it is
the action that you will take which will have more info in the go file for it under
the `nuts` directory. 
