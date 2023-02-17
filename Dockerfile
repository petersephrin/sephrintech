FROM golang:1.19.4-alpine3.17
ENV GO111MODULE=on

RUN cd src
RUN mkdir -p github.com/petersephrin/sephrintech

WORKDIR /go/src/github.com/petersephrin/sephrintech

COPY ./ ./
RUN go build
CMD ["./sephrintech"]
