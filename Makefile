REPO_NAME=https://github.com/lpcruz/strava-updater
SRC=strava-updater


ifndef GOPATH
export GOPATH=$(shell go env "GOPATH")
endif

fmt:
	gofmt -s -w ${SRC}