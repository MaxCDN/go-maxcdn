test: format .PHONY
	# Run Tests
	go test

get:
	# Go Get Deps
	go get -v github.com/garyburd/go-oauth/oauth
	go get -v gopkg.in/jmervine/GoT.v1

docs: format .PHONY
	# Show Docs
	@godoc -ex . | sed -e 's/func /\nfunc /g' | less
	@#                                         ^ add a little spacing for readability

format: .PHONY
	# Go Fmt Source
	@gofmt -w -l $(shell find . -type f -name "*.go")

.PHONY:
