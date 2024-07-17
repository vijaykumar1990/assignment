FROM golang:1.22.5-alpine
WORKDIR /app
ADD . /app/
RUN go build -o ./out/go-practise-project .
EXPOSE 8085
ENTRYPOINT ["./out/go-practise-project"]