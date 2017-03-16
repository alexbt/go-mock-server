.VERSION:=1.0.0
.TIMESTAMP:=`date +%FT%T%z`
.LDFLAGS:=-ldflags "-X bitbucket.org/turbro/swih-version/pkg/version.TIMESTAMP=${.TIMESTAMP} -X bitbucket.org/turbro/swih-version/pkg/version.GIT_HASH=${.GIT_HASH} -X bitbucket.org/turbro/swih-version/pkg/version.VERSION=${.VERSION}"
BIN_NAME:=`basename "$(CURDIR)"`

all:
	${MAKE} getdeps
	${MAKE} clean
	govendor install -a -v -tags heroku ${.LDFLAGS} ./...
	${MAKE} test

install:
	${MAKE} all

heroku:
	${MAKE} all

build:
	${MAKE} getdeps
	${MAKE} clean
	govendor build ${.LDFLAGS} ./...
	${MAKE} test

vendor:
	go get -u github.com/kardianos/govendor
	govendor sync
	govendor update

getdeps:
	go get -u github.com/kardianos/govendor
	go get -d -t ./...

run:
	PORT=8081 ${BIN_NAME}

test:
	govendor test -cover ./...

clean:
	govendor clean

