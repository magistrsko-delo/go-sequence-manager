FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .
COPY .live.env .
RUN cp .live.env .env

EXPOSE 8010

CMD ["/dist/main"]