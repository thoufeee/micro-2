FROM golang:1.25-alpine

WORKDIR /orders

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o orders .

CMD [ "./orders" ]
