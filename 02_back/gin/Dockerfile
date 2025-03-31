FROM golang:latest

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]

# RUN go build -o ./tmp/main main.go

# CMD ["go", "build", "-o", "./tmp/main", "main.go", "&&", "./tmp/main"]