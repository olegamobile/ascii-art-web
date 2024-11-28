FROM golang:1.23-alpine

LABEL description="ASCII-art-web: DOCKERIZE"\
        author="Oleg Balandin, Anastasia Suhareva" \
        version="1.0"

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they changeexit
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/ascii-art-web ./...

EXPOSE 8080

CMD ["ascii-art-web"]