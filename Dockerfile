FROM golang:1.19-alpine AS build-stage
EXPOSE 6060
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o service main.go

FROM alpine
COPY --from=build-stage /usr/src/app/service ./service
CMD ["./service"]
