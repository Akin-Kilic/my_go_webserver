# 1. Adım: Go'nun resmi imajını kullanarak bir aşamalı yapı başlatın
FROM golang:1.22-alpine AS builder

# 2. Adım: Çalışma dizinini ayarla
WORKDIR /app

# 3. Adım: Bağımlılık dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# 4. Adım: Proje dosyalarını kopyala
COPY . .

# 5. Adım: Uygulamayı derle
RUN go build -o main ./main.go

# 6. Adım: Küçük bir Alpine Linux imajını kullanarak çalıştırılabilir dosyayı kopyalayın
FROM alpine:latest

# 7. Adım: Çalışma dizinini ayarla
WORKDIR /root/

# 8. Adım: Builder aşamasından uygulama dosyasını ve config.yaml dosyasını kopyala
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

# 9. Adım: Uygulamayı başlat
CMD ["./main"]