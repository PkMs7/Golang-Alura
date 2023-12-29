# Tópicos do Curso: Go - desenvolvendo uma API Rest

## Passos para criação da api rest

1. Criar o repositório
2. Criar o projeto go.mod
3. Criar o arquivo main.go
4. Instalar o Gin no projeto (github.com/gin-gonic/gin)
5. Inicializar o Gin no projeto (server)
6. Modularizar o código: Inicialmente Rotas e Controladores
7. Criar os models iniciais da aplicação com sua estrutura (struct)
8. Conectar a API a um Banco de Dados
9. Preparar o Docker (Doker para windows: https://cursos.alura.com.br/course/docker-e-docker-compose/task/29235)
10. Criar o docker-compose base da aplicação (usar o comando docker-compose up)
11. Configurar o Servidor na imagem PG Admin (ver o local do host: docker-compose exec postgres sh depois hostmame -i)
12. Configurar o ORM do Go (https://gorm.io/) e conectar com o banco
13. Configurar o driver do banco utilizado (gorm.io/driver/postgres)
14. Criar as tabelas no banco com migrations do gorm
15. Codificar o CRUD com o ORM

## 1.0 Instalando e criando a primeira rota com Gin

- Gin: Gin é uma estrutura web HTTP escrita em Go (Golang). Possui uma API semelhante ao Martini, mas com desempenho até 40 vezes mais rápido que o Martini.

## 2.0 Modularizando o código, modelos e banco de dados

- Params Gin: Recebe os dados da requisição no corpo da requisição, em um objeto em JSON.

## 3.0 Struct, Banco de Dados e ORM

- Migrations: Uma forma de gerenciar as mudanças na estrutura do banco de dados de uma aplicação. Eles permitem que você crie, altere ou exclua tabelas, colunas, índices e outros elementos do banco de dados de forma consistente e controlada.

## 4.0 Implementando rotas HTTP

- 

## 5.0 Deletando, editando e buscando alunos

- 
