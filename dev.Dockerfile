ARG GO_VERSION=1.25
ARG PROJECT_NAME=tapokdrive

FROM golang:${GO_VERSION}-alpine

RUN apk add --no-cache git openssh-client
RUN go install github.com/air-verse/air@latest

WORKDIR /app  

COPY go.sum go.mod ./
RUN go mod download 

CMD ["air"]
