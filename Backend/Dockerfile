FROM golang:1.21.0

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN build -o -v /usr/local/bin/app ./

CMD [ "app" ]