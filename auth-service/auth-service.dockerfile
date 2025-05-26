# FROM golang:1.24-alpine AS builder
# RUN mkdir  /app
# COPY . /app/
# WORKDIR /app
# RUN go build -o authApp ./cmd

FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/authApp /app
COPY ./authApp /app/

CMD [ "/app/authApp" ]
