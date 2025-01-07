# Slack Message Exporter to CSV

Este projeto tem como objetivo converter mensagens do Slack, armazenadas em arquivos JSON, em arquivos CSV. Ele lê arquivos JSON de uma pasta específica, faz o parsing dessas mensagens e as escreve em arquivos CSV. Cada mensagem no Slack é processada, extraindo o nome de usuário e o texto da mensagem, para serem armazenados em formato CSV.

## Funcionalidades

- **Leitura de arquivos JSON**: O programa lê todos os arquivos JSON presentes na pasta `json-files/`.
- **Processamento de dados**: Para cada mensagem, o programa extrai o nome do usuário (se disponível) e o texto da mensagem.
- **Gravação em arquivos CSV**: O programa escreve os resultados em arquivos CSV na pasta `csv-files/`. Cada linha do CSV contém o nome do usuário e o texto da mensagem.

## Estrutura do Projeto

- `json-files/`: Diretório onde os arquivos JSON com as mensagens do Slack são armazenados.
- `csv-files/`: Diretório onde os arquivos CSV serão gerados a partir dos dados extraídos dos JSONs.

## Pré-Requisitos

Certifique-se de ter o Go (Golang) instalado em sua máquina. Você pode verificar se o Go está instalado executando o comando:

```bash
go version
