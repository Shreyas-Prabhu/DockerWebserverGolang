FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go get
RUN go mod download
EXPOSE 4000
CMD ["go", "run", "main.go"]
