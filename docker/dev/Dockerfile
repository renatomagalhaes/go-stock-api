FROM golang:1.22 AS build
ARG APP_ENV
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app cmd/app/main.go

FROM alpine:3.14 AS runtime
RUN apk add --no-cache supervisor tzdata
ENV TZ=America/Sao_Paulo
COPY --from=build /bin/app /bin/app
COPY ./docker/dev/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]