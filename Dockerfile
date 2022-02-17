FROM  golang:1.18beta2-bullseye

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

CMD ["/app/main"]