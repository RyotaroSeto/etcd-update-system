FROM golang:1.21.1-bookworm AS dev

ENV CGO_ENABLED=0
ENV PACKAGES="ca-certificates git curl bash zsh"
ENV ROOT /app

RUN apt-get update && apt-get install -y ${PACKAGES} \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR ${ROOT}

COPY ./ ./

RUN go mod download

EXPOSE 8585

CMD ["go", "run", "main.go"]
