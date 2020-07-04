#!/bin/sh
export PATH="/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/sbin:/usr/local/bin"
pgm="${0##*/}"		# Program basename
progdir="${0%/*}"	# Program directory

cd ${proddir}

# Check go install
if [ -z "$( which go )" ]; then
	echo "error: Go is not installed. Please install go: pkg install -y lang/go"
	exit 1
fi

# Check go version
GOVERS="$( go version | cut -d " " -f 3 )"
if [ -z "${GOVERS}" ]; then
	echo "unable to determine: go version"
	exit 1
fi

export GOPATH="${progdir}"
export GOBIN="${progdir}"

set -e
go get
go build -ldflags "${LDFLAGS} -extldflags '-static'" -o "${progdir}/cbsd-mq-router"
