# 设置原镜像
FROM golang:latest
#设置维护者信息
MAINTAINER teng<18353366911@163.com>
#设置工作目录
WORKDIR $GOPATH/src/gin-example
#copy二进制gin-example文件到docker$GOPATH/src/gin-example目录下
COPY gin-example $GOPATH/src/gin-example
#程序启动的端口
EXPOSE 8080
#最后执行的命令
ENTRYPOINT ["./gin-example"]