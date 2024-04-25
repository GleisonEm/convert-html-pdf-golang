# Use a imagem oficial do Golang como base
FROM golang:1.18

# Defina a variável de ambiente PORT com um valor padrão
RUN apt-get update && apt-get install -y wkhtmltopdf
# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código fonte para o diretório de trabalho
COPY . .

# Compile a aplicação
RUN go build ./cmd/main.go

# Expor a porta configurada
EXPOSE ${PORT}

# Comando para iniciar a aplicação
CMD ["./main"]