FROM ubuntu:latest

RUN mkdir /app

COPY stateService /app
COPY config/config.docker.toml /app

CMD ["/app/stateService"]