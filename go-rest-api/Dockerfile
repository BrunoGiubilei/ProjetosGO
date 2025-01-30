# Use a imagem oficial do Go como base
FROM golang:1.23

# Defina o diretório de trabalho
WORKDIR /app

# Copie os arquivos go.mod e go.sum e baixe as dependências
COPY go.mod go.sum ./
COPY middleware/middleware.go middleware/
COPY greetings/greetings.go greetings/
COPY handlers/handlers.go handlers/

RUN go mod download

# Copie o código do projeto
COPY . .

# Construa o aplicativo
RUN go build -o main .

# Defina as permissões de execução para o binário
RUN chmod 777 main

# Copiar o script de inicialização
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Defina o ponto de entrada para o script de inicialização
ENTRYPOINT ["/entrypoint.sh"]
