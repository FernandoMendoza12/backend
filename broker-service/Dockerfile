# Etapa 1: Construcción del binario
FROM golang:1.23 AS builder

WORKDIR /app

# Copiamos los archivos de dependencias y descargamos módulos
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Etapa 2: Imagen final ligera
FROM alpine:latest

WORKDIR /root/

# Copiamos el binario desde la etapa anterior
COPY --from=builder /app/main .

# Expone el puerto que usará la app
EXPOSE 8080

# Ejecuta el binario
CMD ["./main"]
