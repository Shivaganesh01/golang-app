FROM golang:latest

LABEL maintainer="Shivaganesh <shivaganesh7970@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 4000

RUN go build

RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./gogin"]


