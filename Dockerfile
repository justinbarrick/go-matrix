FROM debian:sid

RUN apt-get update && apt-get install -y libolm-dev
RUN apt-get install -y git golang-go

WORKDIR /app
COPY . .

RUN go build ./cmd/matrixctl

FROM debian:sid

RUN apt-get update && apt-get install -y libolm-dev
RUN apt-get install -y ca-certificates

WORKDIR /app

ENV MATRIX_CONFIG=/app/config.json
ENTRYPOINT ["/app/matrixctl"]
CMD ["slack2matrix"]

COPY --from=0 /app/matrixctl /app/matrixctl
