FROM golang:1.22.3-alpine AS build
WORKDIR /src/
RUN go mod init example/hello
COPY app/main.go /src/
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bin/hello

FROM scratch
EXPOSE 8080
COPY --from=build /bin/hello /bin/hello
ENTRYPOINT ["/bin/hello"]

