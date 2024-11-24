FROM golang:1.18-alphine

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /pricefetcher 

EXPOSE 3000

CMD ["/price"]

