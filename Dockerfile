FROM golang:1.25-alpine AS runtime
WORKDIR /botbot

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o ./server.run ./main.go

# @dev multi-stage bulid for less image size
FROM alpine:3.22 AS runner
WORKDIR /botbot
COPY --from=runtime /botbot/server.run .
EXPOSE 3011

CMD ["./server.run"]
