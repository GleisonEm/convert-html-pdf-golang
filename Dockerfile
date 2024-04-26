# Use a imagem oficial do Ubuntu como base
FROM ubuntu:22.04

# Defina as variáveis de ambiente para a instalação não interativa
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America/Sao_Paulo

# Instale as dependências
RUN apt-get update && apt-get install -y golang-go wkhtmltopdf git

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código fonte para o diretório de trabalho
COPY . .

# Desative a verificação do certificado SSL para o comando go get
ENV GOPROXY=direct,https://proxy.golang.org

# Compile a aplicação
RUN go build ./cmd/main.go

# Expor a porta configurada
EXPOSE ${PORT}

# Comando para iniciar a aplicação
CMD ["./main"]