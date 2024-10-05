# Use the lightweight Golang image based on Alpine for development
FROM golang:alpine

RUN go install github.com/gravityblast/fresh@latest

WORKDIR /app

COPY ./src/run.sh .

RUN chmod +x /app/run.sh

EXPOSE 5000
CMD ["./run.sh"]