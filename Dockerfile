FROM golang:1.14.0-alpine3.11 AS builder


db_project

RUN go version
RUN apk add git

COPY ./ /github.com/aydaZaman/jewelry-shop-backend
WORKDIR /github.com/aydaZaman/jewelry-shop-backend

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./main.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/aydaZaman/db_project/.bin/app .
COPY --from=0 /github.com/aydaZaman/db_project/Api/ ./Api/
COPY --from=0 /github.com/aydaZaman/db_project/database/ ./database/
COPY --from=0 /github.com/aydaZaman/db_project/pickbuy/ ./pickbuy/
COPY --from=0 /github.com/aydaZaman/db_project/profile/ ./profile/
COPY --from=0 /github.com/aydaZaman/db_project/signUP_IN/ ./signUP_IN/
COPY --from=0 /github.com/aydaZaman/db_project/struct/ ./struct/



EXPOSE 8000

CMD [ "./app"]
