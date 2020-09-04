FROM golang:1.15-buster
RUN go get github.com/cespare/reflex
COPY .dev/reflex.conf /
ENTRYPOINT ["reflex", "--verbose", "-c", "/reflex.conf"]
