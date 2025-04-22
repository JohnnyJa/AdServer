FROM ubuntu:latest

RUN mkdir /app

COPY eventCollector /app
COPY config/config.docker.toml /app

CMD ["/app/eventCollector"]
