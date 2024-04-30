generate_grpc_customer:
	protoc --go_out=api/core/grpc/customerService --go_opt=paths=source_relative --go-grpc_out=api/core/grpc/customerService --go-grpc_opt=paths=source_relative --proto_path=C:\dev\freelance\customerService\api\ customerService.proto

build_cerebro:
	docker build -t cerebro:[version] .

generate_docker_image:
	sudo docker run -d -p 3000:3000 cerebro:0
	sudo docker run -e PORT=3001 -p 3000:3000 cerebro:0

run_cerebro:
	sudo docker run -d \
	--name cerebro-service \
	-p 3000:3000 \
	--restart unless-stopped \
	cerebro-service:0

test_cerebro:
	sudo docker run -p 3000:3000 -it --rm cerebro:0