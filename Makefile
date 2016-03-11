default: vet test

vet:
	go tool vet .

test:
	go test .
