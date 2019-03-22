#源镜像
FROM golang:alpine

#作者
MAINTAINER Pharbers "zyqi@pharbers.com"

#LABEL
LABEL 	PhAuthServer.version="1.0.0" maintainer="ClockQ"

# ??
# RUN apk add --no-cache git mercurial

# 下载依赖
# RUN go get github.com/alfredyang1986/blackmirror
# RUN go get github.com/alfredyang1986/BmServiceDef
# RUN go get github.com/PharbersDeveloper/NtmPods
ADD 	 src/	$GOPATH/src/
ADD      src/github.com/PharbersDeveloper/PhAuthServer/static/  $GOPATH/bin/static/

# 设置工程配置文件的环境变量
ENV     PH_AUTH_HOME $GOPATH/src/github.com/PharbersDeveloper/PhAuthServer/resources


# 构建可执行文件
RUN go install -v github.com/PharbersDeveloper/PhAuthServer

# 暴露端口
EXPOSE 31415

# 设置工作目录
WORKDIR $GOPATH/bin

ENTRYPOINT ["PhAuthServer"]
