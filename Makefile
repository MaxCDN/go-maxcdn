# tests without -tabs for go tip
travis: get .PHONY
	# Run Test Suite
	go test -test.v=true

test: format .PHONY
	go test

get:
	# Go Get Deps
	go get github.com/jmervine/GoT

docs: format .PHONY
	@godoc -ex . | sed -e 's/func /\nfunc /g' | less
	@#                                         ^ add a little spacing for readability

readme: test
	# generating readme
	godoc -ex -v -templates "$(PWD)/_support" . > README.md

format: .PHONY
	# Gofmt Source
	@gofmt -s -w -l $(shell find . -type f -name "*.go")

.PHONY:
