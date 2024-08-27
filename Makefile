.PHONY: setup
setup:
	make install
	make install-check
	go mod tidy

.PHONY: install
install:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: install-check
install-check:
	which buf grpcurl protoc-gen-connect-go
	# 🚀 If you see four paths, you're good to go!

.PHONY: buf
buf:
	buf lint
	buf generate
