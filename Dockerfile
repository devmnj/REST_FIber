FROM  golang:1.18beta2-bullseye

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS-linux go build -o main .
 
CMD ["/app/main"]