FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app cmd/app/main.go

FROM alpine:3.14 AS runtime
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]