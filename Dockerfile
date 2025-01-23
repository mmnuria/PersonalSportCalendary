FROM golang:alpine

RUN apk add --no-cache just

WORKDIR /app/test

RUN adduser -D -h /home/mmnuria mmnuria

RUN mkdir -p /home/mmnuria/.cache/go && \
    chmod -R a+w /home/mmnuria/.cache/

ENV GOCACHE=/home/mmnuria/.cache/go

USER mmnuria

CMD just test
