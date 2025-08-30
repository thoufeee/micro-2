FROM golang:1.25-alpine

WORKDIR /products

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o products .

CMD [ "./products" ]