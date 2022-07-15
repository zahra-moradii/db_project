FROM golang:1.17

WORKDIR /backend

COPY go.mod ./go.mod
COPY main.go ./main.go
COPY pkg ./pkg
COPY ui ./ui

RUN go get ./...

RUN go build -o main .

ENTRYPOINT ["/backend/main"]