FROM ubuntu:latest

RUN mkdir /app
RUN mkdir /app/dict

COPY bidHandler /app
COPY config/config.docker.toml /app
COPY ./dict /app/dict

CMD ["/app/bidHandler"]

#ENTRYPOINT ["tail", "-f", "/dev/null"]
