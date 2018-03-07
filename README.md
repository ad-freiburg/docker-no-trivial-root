No Trivial Root for Docker
==========================

**WARNING THIS IS VERY EXPERIMENTAL WITH NO CLAIM OF ACTUAL SECURITY**

This is a very minimal [docker authorization plugin](https://docs.docker.com/engine/extend/plugins_authorization/) 
designed to prevent trivial root escalation on docker already **running with user namespaces**.

To be absolutely clear *without user namespaces this plugin is useless*

One example of such a trivial way of gaining root would be 

    docker run --userns=host --rm -it -v /:/root/ busybox

Building
--------
Make sure you have a Go environment [set up](https://golang.org/doc/install)
then do

    go get github.com/ad-freiburg/docker-no-trivial-root

Setup
-----
Again **make sure** you have [user namespaces enabled](https://docs.docker.com/engine/security/userns-remap/)

Create a startup unit for your init system of choice and make sure
`docker-no-trivial-root` is launched as root on startup

For systemd (most distributions) this can be done with the following steps

    sudo cp $GOPATH/bin/docker-no-trivial-root /usr/sbin
    sudo cp $GOPATH/src/github.com/ad-freiburg/docker-no-trivial-root/systemd/docker-no-trivial-root.service /lib/systemd/system/
    sudo systemctl enable docker-no-trivial-root.service
    sudo systemctl start docker-no-trivial-root.service

**Enable** the plugin by adding `--authorization-plugin=no-trivial-root` to
your dockerd command line.  On Ubuntu this an `ExecStart` in
`/lib/systemd/system/docker.service`

    sudo $EDITOR /lib/systemd/system/docker.service
    sudo systemctl daemon-reload
    sudo systemctl restart docker.service
   

Test It
-------
The following command should give an error message saying that `--userns=host`
is not allowed

    docker run --userns=host --rm -it -v /:/root/ busybox 

also you should get permission denied running `touch /root/foo` inside the container
created by the following command

    docker run --rm -it -v /:/root/ busybox 

What's Prevented
--------------------------
This authorization plugin currently prevents the following `docker run`
parameters

- `--userns=host`
- `--uts=host`
- `--pid=host`
- `--net=host`
- `--log-driver`
- `--log-opt`
- `--cap-add`
- `--device`
- `--security-opt`
- `--privileged`

Configuration
-------------
At this time there is absolutely no configuration, if you want to block
anything more than it currently does you must change the code.

