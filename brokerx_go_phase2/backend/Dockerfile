
FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/brokerx ./cmd

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/brokerx /app/brokerx
ENV PORT=8000
EXPOSE 8000
USER 65532:65532
ENTRYPOINT ["/app/brokerx"]
