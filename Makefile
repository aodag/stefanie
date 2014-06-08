.phony: test

test:
	DATABASE_URL=postgres://docker:docker@localhost:8432/docker go test -v .
