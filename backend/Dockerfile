FROM golang:1.21-alpine as build

ENV HTTPS_PROXY = "fodev.org:8118"
ENV deploy = true
WORKDIR /app

COPY . .

RUN go mod tidy && \
    GOOS=linux GOARCH=amd64 go build -o main && \
    chmod +x ./main && export deploy=true

FROM alpine:latest

RUN apk --no-cache add bash

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]

EXPOSE 3000