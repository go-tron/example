# 基础镜像
FROM alpine

##时区问题
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata \
    && ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 工作目录
WORKDIR /app
RUN mkdir /app/logs

COPY *.yml /app/
COPY build/dory-mall /app/
COPY public/ /app/public/

# 启动容器时执行
ENTRYPOINT /app/dory-mall -c ${GO_OPT_APPLICATION_ENV}
# 使用端口
EXPOSE 80