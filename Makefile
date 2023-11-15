VERSION := 0.1.8

ifneq (,$(wildcard ~/.fropenv))
    include ~/.fropenv
    export
endif

include $(CURDIR)/.version

fmt:
	go fmt

lint:
	golint -set_exit_status main.go
	golint -set_exit_status config
	golint -set_exit_status gpg
	golint -set_exit_status tar
	golint -set_exit_status install
	golint -set_exit_status utils
	golint -set_exit_status cmd

test: lint fmt
	go run main.go testing

staging:
	git checkout master
	git pull
	git fetch -p
	git branch -D staging
	git checkout -b staging
	git push --set-upstream origin staging

current:
	git checkout master
	git pull
	git fetch -p
	git branch -D staging
	git checkout staging

patch:
	git aftermerge patch || exit 1

minor:
	git aftermerge minor || exit 1

major:
	git aftermerge major || exit 1

build:
	go build -o fropilr -ldflags '-X fropilr/config.SystemPasswd=${SYSTEM_PASSWD}' main.go

deb: build
	mv fropilr deploy/debian/usr/local/bin
	dpkg-deb --build deploy/debian
	mkdir -p dist
	mv deploy/debian.deb dist/fropilr-${VERSION}.deb
	rm -f deploy/debian/usr/local/bin/fropilr

dist-gzip: build
	mkdir -p dist/${APPNAME}-${VERSION}/
	cp fropilr dist/${APPNAME}-${VERSION}/
	cd dist && tar cvzf ${APPNAME}-${VERSION}.tar.gz ${APPNAME}-${VERSION} && rm -rf ${APPNAME}-${VERSION}

clean:
	rm -rf dist
