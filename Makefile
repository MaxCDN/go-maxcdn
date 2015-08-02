# tests without -tabs for go tip

# Run Tests
travis: get
	go test -v

# Run Tests
test: format
	go test

# Go Get Deps
get:
	go get github.com/garyburd/go-oauth/oauth
	go get github.com/jmervine/GoT

# Show Docs
docs: format
	@godoc -ex . | sed -e 's/func /\nfunc /g' | less # add a little spacing for readability

# Go Fmt Source
format:
	@gofmt -s -w -l $(shell find . -type f -name "*.go")

.PHONY: travis test docs format
