FROM ubuntu:latest

RUN mkdir /app

COPY bidHandler /app
COPY config/config.docker.toml /app

CMD ["/app/bidHandler"]