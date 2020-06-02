FROM golang:1.13

# copy source
RUN mkdir -p /go/src/github.com/pmtoan/example-go-service
WORKDIR /go/src/github.com/pmtoan/example-go-service
COPY . .

RUN go build main.go

EXPOSE 8080

CMD ["./main"]

