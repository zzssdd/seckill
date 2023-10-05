FROM golang:1.21.1

ENV GO111MODULE=ON \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=am64

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /app/build/seckill .

EXPOSE 8888
ENTRYPOINT ["/app/build/seckill"]