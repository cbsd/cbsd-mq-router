all: 
	@./build.sh
clean:
	rm -f bs_router
install: all
	install bs_router /usr/local/sbin
	install rc.d/bs_router /usr/local/etc/rc.d/bs_router
uninstall: 
	rm -f /usr/local/sbin/bs_router /usr/local/etc/rc.d/bs_router

