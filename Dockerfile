FROM golang:1.20.2-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o book-management-system ./cmd/

FROM alpine:3.17.2

COPY --from=builder /app/book-management-system .
EXPOSE 8181
CMD [ "./book-management-system" ]