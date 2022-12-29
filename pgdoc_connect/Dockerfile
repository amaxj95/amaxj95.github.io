FROM golang:1.19.0 AS builder

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
RUN curl --request GET && \
      --url http://127.0.0.1:5050/ && \
      --cookie 'pga4_session=104011d7-e91b-4329-a5a9-99e2c834536a!R2Uukm%2Bn7S8GMAMObG%2FO7BxJG%2BSMBoAG5ENXktrCe0U%3D'
