FROM ubuntu:latest

RUN mkdir /app

COPY profileMonitor /app
COPY config/config.docker.toml /app

CMD ["/app/profileMonitor"]
