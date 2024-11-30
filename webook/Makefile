.PHONY: docker
docker:
	@echo "Cleaning up old build artifacts..."
	@rm -f webook || true # 删除旧的 webook 文件
	@echo "Setting build environment..."
	GOOS=linux GOARCH=arm go build -o webook . # 交叉编译 Linux ARM 架构的可执行文件
	@echo "Removing old Docker image..."
	@docker rmi -f flycash/webook:v0.0.1 || (echo "Docker image removal failed, continuing..." && exit 0)
	@echo "Building new Docker image..."
	@docker build -t flycash/webook:v0.0.1 .