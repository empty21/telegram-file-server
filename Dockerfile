FROM golang:1.22-alpine as build

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY ["pkg/", "./pkg/"]
COPY ["cmd/", "./cmd/"]
RUN ls -la /server

RUN go build -o server telegram-file-server/cmd/server

FROM aiogram/telegram-bot-api:latest as RUN
COPY --from=build /server/server /server
COPY docker-entrypoint.sh docker-entrypoint.sh
RUN chmod +x docker-entrypoint.sh
EXPOSE ${PORT:-8080}
