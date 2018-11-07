ARG appname="go-box"

# build go
FROM golang:1.11.1-alpine3.8 as builder0

ARG appname

WORKDIR /go/src/$appname
COPY . .
RUN mkdir -p dist && go build -o dist/app .

# build web
FROM node:10.9.0-alpine as builder1

WORKDIR /root
COPY web .
RUN yarn install && npm run build

# pack
FROM alpine:3.8

ARG appname

WORKDIR /srv
COPY --from=builder0 /go/src/$appname/dist/app .
COPY --from=builder1 /root/dist web/dist
ENTRYPOINT ["/srv/app", "-c", "config.json"]
