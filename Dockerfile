FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/sbdb-semester
ENV GO111MODULE on
WORKDIR /go/src/sbdb-semester
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/sbdb-semester/sbdb-semester /
WORKDIR /
CMD ./sbdb-semester
ENV PORT 8000
EXPOSE 8000