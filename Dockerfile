FROM golang:alpine

RUN apk add --no-cache just

WORKDIR /app/test

RUN adduser -D -h /home/mmnuria mmnuria

COPY test_config.yaml /app/test/test_config.yaml

RUN touch /app/test/app.log && \
    chown mmnuria:mmnuria /app/test/app.log && \
    chmod 666 /app/test/app.log
    
RUN mkdir -p /home/mmnuria/.cache/go && \
    chmod -R a+w /home/mmnuria/.cache/ && \
    chmod +r /app/test/test_config.yaml 

ENV GOCACHE=/home/mmnuria/.cache/go

USER mmnuria

ENTRYPOINT just test
