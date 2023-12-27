# Tópicos do Curso: Go - desenvolvendo uma API Rest

## Passos para criação da api rest

1. Criar o repositório
2. Criar o projeto go.mod
3. Criar o arquivo main.go
4. Configurar inicialmente o servidor para acessar pelo navegador
5. Modularizar o código: Inicialmente Rotas e Controladores
6. Criar os models iniciais da aplicação com sua estrutura (struct)
7. Montar a exibição do JSON através do controlador e criar a rota que será exposto
8. Criar um roteador (nesse projeto utilizado o Gorilla Mux: https://github.com/gorilla/mux)
9. Personalizar as rotas de acordo com os dados que se quer. Ex.: Trazer dados por id específico
10. Preparar o Docker (Doker para windows: https://cursos.alura.com.br/course/docker-e-docker-compose/task/29235)
11. Criar a estrutura de tabelas do banco
12. Criar o docker-compose base da aplicação (usar o comando docker-compose up)
13. Configurar o Servidor na imagem PG Admin (ver o local do host: docker-compose exec postgres sh depois hostmame -i)
14. Configurar o ORM do Go (https://gorm.io/)
15. Codificar o CRUD com o ORM

## 1.0 Json, Request, Response e Go

- Json: O JavaScriptObjectNotation é um formato de dados leve e de fácil leitura utilizado para troca de informações entre sistemas computacionais

- Request: A Request ou requisição traduzindo diretamente para português, é o pedido que um cliente realiza a nosso servidor.

- Response: A Response (resposta) nada mais é do que a resposta que o servidor envia ao cliente.

## 2.0 Roteador, recursos por ID e Docker

- Método GET: O método GET solicita a representação de um recurso específico. Requisições utilizando o método GET devem retornar apenas dados.

- Docker: Docker é um conjunto de produtos de plataforma como serviço que usam virtualização de nível de sistema operacional para entregar software em pacotes chamados contêineres. 

- Contêiners: Os contêineres são isolados uns dos outros e agrupam seus próprios softwares, bibliotecas e arquivos de configuração.

## 3.0 Conexão com banco e exibindo os dados

- ORM: Object-Relational Mapping (ORM), em português, mapeamento objeto-relacional, é uma técnica para aproximar o paradigma de desenvolvimento de aplicações orientadas a objetos ao paradigma do banco de dados relacional.

## 4.0 Criando, deletando e editando com Gorm

- CRUD: Create, Read, Update, Delete

## 5.0 Middleware e integração com front-end

- Middleware: Uma funcionalidade intermediária, evitando duplicidade de código.

- CORS: Cross-Origin Resource Sharing ou CORS é um mecanismo que permite que recursos restritos em uma página da web sejam recuperados por outro domínio fora do domínio ao qual pertence o recurso que será recuperado.