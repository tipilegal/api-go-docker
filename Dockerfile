FROM golang:1.23-alpine AS build  

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server . 
EXPOSE 3000

CMD ["./server"]
