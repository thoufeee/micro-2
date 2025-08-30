
FROM golang:1.25-alpine

WORKDIR /dock

COPY . .

RUN go build -o myapp .

CMD [ "./myapp" ]