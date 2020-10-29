FROM golang:1.15-rc-alpine as builder

WORKDIR /app/api/

ENV GO111MODULE=on
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine

# Security related package, good to have.
RUN apk --no-cache add ca-certificates

RUN mkdir /app/api
WORKDIR /app/api

COPY --from=builder /app/api .

ENTRYPOINT [ "/app/api/maze" ]