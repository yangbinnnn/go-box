# build go
FROM golang:1.11.1-alpine3.8 as builder0

Env NAME go-box

WORKDIR /go/src/$NAME
COPY . .
RUN mkdir -p dist && go build -o dist/$NAME .

# build web
FROM node:10.9.0-alpine as builder1

WORKDIR /root
COPY web .
RUN apk update && apk add yarn
RUN yarn install && npm run build

# pack
FROM alpine:3.8

Env NAME go-box

WORKDIR /srv
COPY --from=builder0 /go/src/$NAME/dist/$NAME .
COPY --from=builder1 /root/dist web
CMD /srv/$NAME
