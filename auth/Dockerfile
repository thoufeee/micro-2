FROM golang:1.25-alpine

WORKDIR /auth

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth .

CMD [ "./auth" ]