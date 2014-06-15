# tests without -tabs for go tip
travis: get .PHONY
	# Run Tests
	go test -test.v

test: format .PHONY
	# Run Tests
	go test -test.v

get:
	# Go Get Deps
	go get github.com/jmervine/GoT

docs: format .PHONY
	# Show Docs
	@godoc -ex . | sed -e 's/func /\nfunc /g' | less
	@#                                         ^ add a little spacing for readability

readme: test
	# Generate README
	godoc -ex -v -templates "$(PWD)/_support" . > README.md

format: .PHONY
	# Go Fmt Source
	@gofmt -s -w -l $(shell find . -type f -name "*.go")

.PHONY:
