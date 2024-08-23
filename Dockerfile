FROM golang:1.23

WORKDIR /app/

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o /user-service

EXPOSE 8080

CMD ["/user-service"]