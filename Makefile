build-vm:
	@echo "Building VM"
	go build -o bin/flux cmd/vm/main.go