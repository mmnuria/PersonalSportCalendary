FROM golang:alpine

RUN apk add --no-cache just

RUN adduser -D -h /home/personalsportcalendary personalsportcalendary

RUN mkdir -p /app/test && chown -R personalsportcalendary /app/test

WORKDIR /app/test

RUN mkdir -p /home/personalsportcalendary/.cache/go && \
    chmod -R a+w /home/personalsportcalendary/.cache/

ENV GOCACHE=/home/personalsportcalendary/.cache/go

USER personalsportcalendary

ENTRYPOINT just test
