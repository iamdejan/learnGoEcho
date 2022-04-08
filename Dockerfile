FROM golang:1.18.0-bullseye
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /learn-go-echo
EXPOSE 1111
CMD [ "/learn-go-echo" ]
