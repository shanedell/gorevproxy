#########################
# Stage 1: Build binary #
#########################

FROM cgr.dev/chainguard/go:latest AS builder

COPY go.mod go.sum /app/
WORKDIR /app

RUN go mod download

COPY ./ /app

RUN go build -o gorevproxy . && chmod +x gorevproxy

#######################
# Stage 2: Run binary #
#######################

FROM cgr.dev/chainguard/go:latest

COPY --from=builder /app/gorevproxy /gorevproxy

ENTRYPOINT ["/gorevproxy"]
