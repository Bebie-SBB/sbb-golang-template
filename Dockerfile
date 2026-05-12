FROM golang:1.24-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build -o /app/server .
RUN CGO_ENABLED=0 go build -o /app/healthcheck ./cmd/healthcheck.go

FROM alpine:3.20
RUN apk --no-cache add ca-certificates sqlite-libs

WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/healthcheck .
COPY .static/ .static/

ENV PORT=9090
ENV HEALTH_URL=http://localhost:9090/health

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/app/healthcheck"]

EXPOSE 9090
ENTRYPOINT ["/app/server"]
CMD ["--env=production", "--loadEnvFile=false", "--loadConfFile=false"]
