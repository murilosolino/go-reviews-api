FROM golang:1.25

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    default-mysql-client \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/air-verse/air@latest

WORKDIR /workspace

# Delve DAP debug server
EXPOSE 2345

# Common app port
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
