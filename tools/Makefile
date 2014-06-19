GOROOT=$(shell go env GOROOT)
PLATFORMS=darwin/386 \
			darwin/amd64 \
			freebsd/386 \
			freebsd/amd64 \
			freebsd/arm \
			linux/386 \
			linux/amd64 \
			linux/arm \
			windows/386 \
			windows/amd64

# build tools
###
build: .PHONY
	# exec tools/build-all.sh
	bash build-all.bash

clean: .PHONY
	# remove previous tools builds
	@find . -type d -name builds -exec rm -vr {} \;

# setup and sub-tasks setup go for crosscompiling
###
setup: $(PLATFORMS) .PHONY

darwin/386:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=darwin GOARCH=386 ./make.bash --no-clean 2>&1'

darwin/amd64:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=darwin GOARCH=amd64 ./make.bash --no-clean 2>&1'

freebsd/386:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=freebsd GOARCH=386 ./make.bash --no-clean 2>&1'

freebsd/amd64:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=freebsd GOARCH=amd64 ./make.bash --no-clean 2>&1'

freebsd/arm:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=freebsd GOARCH=arm ./make.bash --no-clean 2>&1'

linux/386:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=linux GOARCH=386 ./make.bash --no-clean 2>&1'

linux/amd64:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=linux GOARCH=amd64 ./make.bash --no-clean 2>&1'

linux/arm:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=linux GOARCH=arm ./make.bash --no-clean 2>&1'

windows/386:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=windows GOARCH=386 ./make.bash --no-clean 2>&1'

windows/amd64:
	sudo bash -c 'cd $(GOROOT)/src && GOOS=windows GOARCH=amd64 ./make.bash --no-clean 2>&1'

.PHONY:
