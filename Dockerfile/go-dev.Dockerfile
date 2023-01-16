FROM golang:1.20-rc-alpine
RUN mkdir /app
ADD microservice/ /app/
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

RUN go mod vendor
CMD ["air", "-c", ".air.toml"]