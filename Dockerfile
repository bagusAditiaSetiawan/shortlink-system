FROM golang:1.22.0-alpine

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum untuk pengelolaan dependency
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi
RUN go build

# Expose port (sesuaikan dengan port yang digunakan aplikasi Anda)
EXPOSE 8000

# Jalankan aplikasi
CMD ["./shortlink-system"]