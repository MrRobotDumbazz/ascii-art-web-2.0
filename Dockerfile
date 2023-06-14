# builder image...
FROM golang:1.17-alpine3.16 AS builder
LABEL stage=builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN  apk add build-base && go build -o main ./cmd/main.go

## generate clean, final image...
FROM alpine:3.16 AS runner
LABEL stage=runner
LABEL maintainer="Made by AmayevArtyom && Mr.RobotDumbazz"
LABEL org.label-schema.description="ascii-art-web"
WORKDIR /app
COPY --from=builder /app/main ./
COPY /internal /app/internal
COPY /static /app/static
COPY /templates /app/templates
EXPOSE 8080

# Run the executable
CMD ["./main"]