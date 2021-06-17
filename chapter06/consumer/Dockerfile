FROM golang:1.16
WORKDIR /go/src/github.com/worker
COPY main.go go.mod go.sum .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/worker/app .
CMD ["./app"]  