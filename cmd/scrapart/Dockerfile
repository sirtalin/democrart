FROM golang:latest

WORKDIR /usr/local/scrapart

ADD . .

RUN go build -o scrapart -v cmd/scrapart/main.go

EXPOSE 3000

CMD ["./scrapart"]