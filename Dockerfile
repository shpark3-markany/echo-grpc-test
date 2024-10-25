FROM golang:alpine AS build
COPY . .
RUN go build -o /app


FROM alpine AS app
WORKDIR /fin
COPY --from=build /app .
COPY configs/develop.env .
RUN apk add curl --no-cache