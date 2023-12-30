    FROM golang:latest
    LABEL maintainer="jzacko <zacko.jason@gmail.com>"
    WORKDIR /app
    COPY . .
    RUN go build -o main .
    CMD ["./main"]