# syntax=docker/dockerfile:1

FROM golang:1.17.1

WORKDIR /app
COPY . /app

RUN cd /app && go build -o category-svcs

CMD ["/app/category-svcs"]