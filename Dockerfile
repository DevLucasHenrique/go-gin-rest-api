FROM golang:1.26.1

WORKDIR /go/src/app

COPY . .

EXPOSE 8000

RUN go build -o main cmd/main.go

CMD [ "./main" ]