#!/bin/bash

# Imprimir o diretório de trabalho atual e listar os arquivos
echo "Diretório de trabalho atual:"
pwd
echo "Conteúdo do diretório /app:"
ls -la /app
echo "Conteúdo do diretório /:"
ls -la /

# Executar o binário principal
./main
