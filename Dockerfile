ARG GO_VERSION=1.24.3

# Build frontend assets
FROM node:20-bookworm as frontend-builder

WORKDIR /usr/src/app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

# Build Go application
FROM golang:${GO_VERSION}-bookworm as backend-builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

WORKDIR /app
COPY --from=backend-builder /run-app /usr/local/bin/
COPY --from=frontend-builder /usr/src/app/frontend/dist ./frontend/dist

CMD ["run-app"]
