FROM golang:1.15-alpine

WORKDIR /go/src/notebox

COPY . .

RUN go get  -d -v ./...
RUN go install -v ./...

EXPOSE 80/tcp

CMD ["notebox", "-p", "80"]
