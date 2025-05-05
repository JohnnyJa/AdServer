FROM ubuntu:latest

RUN mkdir /app

COPY packageService /app
COPY config/config.docker.toml /app

CMD ["/app/packageService"]
