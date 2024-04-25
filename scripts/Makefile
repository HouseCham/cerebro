generate_grpc_customer:
	protoc --go_out=api/core/grpc/customerService --go_opt=paths=source_relative --go-grpc_out=api/core/grpc/customerService --go-grpc_opt=paths=source_relative --proto_path=C:\dev\freelance\customerService\api\ customerService.proto

generate_docker_image:
	docker run -e PORT=3000 -p 3000:3000 cerebro-service:latest