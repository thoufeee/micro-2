

FROM golang:1.25-alpine

WORKDIR /compose/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp .

CMD [ "./myapp" ]