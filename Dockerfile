FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 5000
CMD ["./main"]


# docker build -t  img .
# docker run -it -p 8000:8080 --name c-1 img
