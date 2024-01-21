FROM golang:1.17.0-alpine as build

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -o /go-app

FROM alpine:latest
RUN mkdir -p /app /logs
WORKDIR /app

COPY --from=build /go-app /app/go-app
COPY --from=build /app/.env /app
ENTRYPOINT [ "./go-app" ]