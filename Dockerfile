FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./microservice-uporabniki microservice-uporabniki


FROM golang:1.19-alpine
COPY .env /.env
COPY .env /go
COPY --from=build /app/microservice-uporabniki /microservice-uporabniki
EXPOSE 8080
CMD ["/microservice-uporabniki"]

