# Build Stage
FROM golang:1.23.0-alpine AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a \
    -mod=vendor \
    -installsuffix cgo \
    -o user-management \
    cmd/server/main.go

# Final stage
FROM scratch
COPY --from=build /app/user-management /user-management
ENTRYPOINT ["/user-management"]