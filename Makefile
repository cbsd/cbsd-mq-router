UNAME_S := $(shell uname -s)

all: 
	@./build.sh

clean:
	rm -f cbsd-mq-router
	rm -rf src

install: all
	install cbsd-mq-router /usr/local/sbin
ifeq ($(UNAME_S),Linux)
	install systemd/cbsd-mq-router.service /lib/systemd/system/cbsd-mq-router.service
	@test -d /var/log/cbsdmq || mkdir -m 0755 /var/log/cbsdmq
	@test -d /var/log/cbsd_mq_router || mkdir -m 0755 /var/log/cbsd_mq_router
	@test -r /etc/cbsd-mq-router.json || sed 's:/dev/stdout:/var/log/cbsd_mq_router/cbsd_mq_router.log:g' etc/cbsd-mq-router.json > /etc/cbsd-mq-router.json
else
	install rc.d/cbsd-mq-router /usr/local/etc/rc.d/cbsd-mq-router
endif

uninstall:
ifeq ($(UNAME_S),Linux)
	rm -f /usr/local/sbin/cbsd-mq-router /lib/systemd/system/cbsd-mq-router.service
else
	rm -f /usr/local/sbin/cbsd-mq-router /usr/local/etc/rc.d/cbsd-mq-router
endif
