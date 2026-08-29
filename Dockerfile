# Estágio de Build (atualizado para Go 1.25)
FROM golang:1.25 AS builder

WORKDIR /app

# Copia todo o projeto
COPY . .

# Compila o binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Estágio Final Ultraleve
FROM alpine:latest
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/main .

#EXPOSE 8080
CMD ["./main"]
