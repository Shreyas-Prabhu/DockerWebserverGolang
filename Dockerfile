FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go get
RUN go mod download
RUN go build -o main .
EXPOSE 4000
CMD ["./main"]
