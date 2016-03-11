FROM golang
COPY . /go/src/go-service-test
WORKDIR /go/src/go-service-test
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]
