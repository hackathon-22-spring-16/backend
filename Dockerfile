FROM golang:1.17-alpine AS build

WORKDIR /go/src/github.com/hackathon-22-spring-16/backend
COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go mod tidy

# usermodなどで手元のUIDが変わっている場合は.envに記述する
# RUN chown -R ${UID:-1000}:${GID:-1000} ./

CMD go run main.go
