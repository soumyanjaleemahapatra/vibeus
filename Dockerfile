FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/soumyanjaleemahapatra/vibeus/
WORKDIR /go/src/github.com/soumyanjaleemahapatra/vibeus/
RUN go mod download
COPY . /go/src/github.com/soumyanjaleemahapatra/vibeus/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/vibeus github.com/soumyanjaleemahapatra/vibeus/

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/soumyanjaleemahapatra/vibeus/build/vibeus /usr/bin/vibeus
EXPOSE 8090 8090
ENTRYPOINT ["/usr/bin/vibeus"]
