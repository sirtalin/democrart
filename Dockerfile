FROM golang:latest

WORKDIR /usr/local/democrart

ADD . .

RUN go build -o democrart -v cmd/democrart/main.go

EXPOSE 3000

CMD ["./democrart"]