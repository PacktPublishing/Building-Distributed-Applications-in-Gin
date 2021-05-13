FROM golang:1.16
WORKDIR /go/src/github.com/api
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
ARG API_VERSION
ENV API_VERSION=$API_VERSION
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/api/app .
CMD ["./app"]  