FROM golang:alpine

WORKDIR /app

COPY . .

RUN GOOS=linux go mod download
RUN GOOS=linux go build -o main cmd/main.go

# EXPOSE 8080

CMD [ "./main" ]