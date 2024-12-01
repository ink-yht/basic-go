## 使用官方的轻量化 Go 镜像作为构建镜像
#FROM golang:1.23.3 AS builder
#
## 设置工作目录
#WORKDIR /app
#
## 将项目源代码复制到构建镜像中
#COPY . .
#
## 编译二进制文件
#RUN GOOS=linux GOARCH=arm go build -o webook .

# 使用轻量级操作系统作为最终镜像
#FROM ubuntu:22.04
#
## 设置工作目录
#WORKDIR /app
#
## 将编译后的二进制文件从构建镜像复制到最终镜像
#COPY --from=builder /app/webook /app/webook
#
## 设置容器入口点
#ENTRYPOINT ["/app/webook"]

FROM ubuntu:22.04
COPY webook /app/webook
WORKDIR /app
CMD ["/app/webook"]