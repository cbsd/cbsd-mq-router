#!/bin/sh
pgm="${0##*/}"				# Program basename
progdir="${0%/*}"			# Program directory
workdir=$( realpath ${progdir} )	# realpath dir
cd ${workdir}

# Check go install
if [ -z "$( which go )" ]; then
	echo "error: Go is not installed. Please install go: pkg install -y lang/go"
	echo "PATH: ${PATH}"
	exit 1
fi

# Check go version
GOVERS="$( go version | cut -d " " -f 3 )"
if [ -z "${GOVERS}" ]; then
	echo "unable to determine: go version"
	exit 1
fi

export GOPATH="${workdir}"
export GOBIN="${workdir}"
#export GO111MODULE=off

set -e
go get
go build -ldflags "${LDFLAGS} -extldflags '-static'" -o "${workdir}/cbsd-mq-router"
