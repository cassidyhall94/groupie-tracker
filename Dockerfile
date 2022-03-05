FROM golang:1.16-bullseye as build

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 go build -o /groupie-tracker

FROM scratch

WORKDIR /

COPY ./web /web
COPY --from=build /groupie-tracker /groupie-tracker
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/groupie-tracker"]