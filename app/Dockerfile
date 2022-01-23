# build stage
FROM golang AS build-env
ADD . /creater
RUN cd /creater && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final stage
FROM scratch
COPY --from=build-env /creater/app /creater/app
COPY --from=build-env etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /creater
ENTRYPOINT ["./app"]