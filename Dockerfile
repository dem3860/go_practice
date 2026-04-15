FROM --platform=linux/amd64 golang:1.25.9-alpine

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN apk update && \
    apk add --no-cache git tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone && \
    go install github.com/air-verse/air@latest

RUN go mod download

COPY . ./

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]