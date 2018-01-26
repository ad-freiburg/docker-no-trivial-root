No Trivial Root for Docker
==========================

**WARNING THIS IS VERY EXPERIMENTAL WITH NO CLAIM OF ACTUAL SECURITY*

This is a very minimal [docker authorization plugin](https://docs.docker.com/engine/extend/plugins_authorization/) 
designed to prevent trivial root escalation on docker already **running with user namespaces**.

To be absolutely clear *without user namespaces this plugin is useless*

One example of such a trivial way of gaining root would be 

    docker run --userns=host --rm -it -v /:/root/ ubuntu:16.04 /bin/bash 

Building
--------
Make sure you have a Go environment [set up](https://golang.org/doc/install)
then do

    go get github.com/niklas88/docker-no-trivial-root

Setup
-----
Create a startup unit for your init system of choice and make sure
`docker-no-trivial-root` is launched as root on startup

Configuration
-------------
At this time there is absolutely no configuration, if you want to block
anything more than it currently does you must change the code.

Usage
-----
Add `--authorization-plugin=no-trivial-root` to your dockerd command line.
On Ubuntu this an `ExecStart` in `/lib/systemd/system/docker.service`
