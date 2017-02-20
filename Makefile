PKG=$(shell glide nv)

default: vet test

vet:
	go vet $(PKG)

test:
	go test $(PKG)

