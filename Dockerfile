FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o "/WebCompressor"

FROM alpine

COPY . .
COPY --from=builder /WebCompressor /

EXPOSE 8080

CMD [ "/WebCompressor" ]
