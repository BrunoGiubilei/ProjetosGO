# Use a imagem oficial do Python como base
FROM python:3.9

# Defina o diretório de trabalho
WORKDIR /app

# Copie o código do frontend
COPY . .

# Instale o servidor HTTP necessário (caso não esteja no Python padrão)
RUN pip install httpserver

# Exponha a porta 8080
EXPOSE 8080

# Defina o comando de inicialização
CMD ["python3", "-m", "http.server", "8080"]
