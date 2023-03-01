
FROM golang:latest AS builder
WORKDIR /app
COPY http_server/go.mod ./
COPY http_server/go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN air init
COPY . . 
# RUN go build -o /app/tmp/main
EXPOSE 8000
CMD air ./air.toml
# CMD [ "/duspot_server" ]


# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=builder /duspot_server ./
# CMD ["./duspot_server"]
