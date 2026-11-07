# ---------- Stage 1: build ----------
FROM golang:1.22.5 AS build

WORKDIR /app

# Copy module files first so dependency download is cached
COPY go.mod ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Static binary, no cgo (avoids the stdlib.h / gcc problem too)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# ---------- Stage 2: run ----------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=build /app/main .
COPY --from=build /app/static ./static

EXPOSE 8080

CMD ["./main"]
