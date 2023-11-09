FROM golang:1.20.5 AS Builder

LABEL Owner="Habeebullah Ashiru <ashiruhabeeb472gmail.com>"
LABEL Version="0.0.1"

WORKDIR /app
COPY . .

RUN go mod download

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 7273

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=Builder /app/main .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 7273

CMD [ "/app/main" ]
