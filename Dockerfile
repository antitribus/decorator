FROM golang as builder

WORKDIR /opt
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM scratch

WORKDIR /bin
COPY --from=builder /opt/app .
CMD ["./app"]