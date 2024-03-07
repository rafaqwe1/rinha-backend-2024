FROM golang:1.21.0 as golang

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /api

FROM scratch

COPY --from=golang /api /api
ENTRYPOINT ["/api"]