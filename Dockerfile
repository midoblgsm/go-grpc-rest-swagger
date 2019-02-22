FROM golang as builder
WORKDIR /go/src/github.com/midoblgsm/gogrpcrestswagger-boilerplate
RUN go get github.com/Masterminds/glide
COPY . .
RUN go get github.com/rakyll/statik
RUN statik -src swaggerui

RUN glide up --strip-vendor


RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -v -a --ldflags '-w -linkmode external -extldflags "-static"' -installsuffix cgo -o bin/server server/main.go


FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/midoblgsm/gogrpcrestswagger-boilerplate/bin/server .

CMD ["./server"]

EXPOSE 7788 7789