FROM golang:1.15.11-alpine3.12

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    PORT=3000

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

EXPOSE 3000

ENTRYPOINT [ "./app" ]