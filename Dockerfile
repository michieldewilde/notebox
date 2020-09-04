FROM golang:1.15-alpine

WORKDIR /go/src/notebox

COPY . .

RUN go get  -d -v ./...
RUN go install -v ./...

EXPOSE 80/tcp

CMD notebox -http 0.0.0.0:80 -datasource test
