FROM golang

WORKDIR /go/src/app

COPY . ./

RUN go mod download 

RUN go build -o /go/bin/app

EXPOSE 3000

CMD ["/go/bin/app"]
