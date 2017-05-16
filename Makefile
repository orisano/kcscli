# vim: set noexpandtab:

.PHONY: ci

ci:
	go vet .
	golint .
