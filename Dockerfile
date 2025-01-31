FROM golang:alpine

RUN apk add --no-cache just

WORKDIR /app/test

RUN adduser -D -h /home/personalsportcalendary personalsportcalendary

RUN mkdir -p /home/personalsportcalendary/.cache/go && \
    chmod -R a+w /home/personalsportcalendary/.cache/

ENV GOCACHE=/home/personalsportcalendary/.cache/go

USER personalsportcalendary

ENTRYPOINT just test
