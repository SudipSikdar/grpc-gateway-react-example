compile:
	@echo "Compiling proto files..."
	@protoc -I ./proto  -I ./third_party \
	--go_out ./pb --go_opt paths=source_relative \
	--go-grpc_out ./pb --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./pb --grpc-gateway_opt paths=source_relative \
	./proto/service.proto
	@echo "Compiling done"