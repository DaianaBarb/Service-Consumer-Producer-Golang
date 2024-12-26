# Etapa 1: Build (compilação da aplicação)
# Usamos uma imagem Golang para compilar o código
FROM golang:1.23.4-alpine AS build

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos go.mod e go.sum para o contêiner (para instalar dependências)
COPY go.mod go.sum ./

# Baixar as dependências (sem copiar o código ainda)
RUN go mod tidy

# Copiar o código da aplicação (diretório cmd/api-server)
COPY . .

# Definir o diretório de trabalho como o local do arquivo main.go
WORKDIR /app/cmd/api-server

# Compilar o código Go e gerar o binário 'api-server'
RUN go build -o /usr/local/bin/api-server .

# Etapa 2: Imagem Final (imagem minimalista)
# Usamos uma imagem leve (Alpine) para rodar o binário
FROM alpine:latest

# Instalar as dependências necessárias para executar a aplicação (se necessário)
RUN apk --no-cache add ca-certificates

# Copiar o binário compilado da etapa anterior
COPY --from=build /usr/local/bin/api-server /usr/local/bin/api-server

# Expor a porta onde a aplicação irá rodar (ajuste conforme sua aplicação)
EXPOSE 8080

# Definir o comando que será executado ao rodar o contêiner
CMD ["/usr/local/bin/api-server"]