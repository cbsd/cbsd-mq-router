# cbsd-mq-router

CBSD message queue router: subscribe to events and deliver to the CBSD
Client + sample: https://github.com/cbsd/bs_router-client

# Installation

Install dependency first:

```
pkg install -y beanstalkd
service beanstalkd enable
sysrc beanstalkd_flags="-l 127.0.0.1"
service beanstalkd start
```

Build cbsd-mq-router:

```
setenv GOPATH $( realpath . )
go get
go build
pkg update -f
```

