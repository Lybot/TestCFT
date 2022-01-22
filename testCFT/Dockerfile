FROM golang:latest
WORKDIR /app
COPY . .
RUN go build /app/cmd/app/apiserver/main.go
EXPOSE 32228
ENTRYPOINT [ "./main" ]