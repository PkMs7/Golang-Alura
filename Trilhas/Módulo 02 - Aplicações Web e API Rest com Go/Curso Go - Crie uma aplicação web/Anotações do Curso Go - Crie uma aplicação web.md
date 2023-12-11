# Tópicos do Curso: Go - crie uma aplicação web

## 1.0 Servidor e struct de produtos

- Local do projeto ./Repositorios/appWeb

## 2.0 Conectando com banco de dados

- Cógigo sql em ./Repositorios/appWeb/db
- Driver do PostgreSQL: github.com/lib/pq
- Demais drivers Go: https://github.com/golang/go/wiki/SQLDrivers

## 3.0 Refatoração e página de novos produtos

- Arquitetura de MVC:

    - **controllers:** Local onde ficam as configurações de controle da aplicação
    - **db:** Local onde ficam as configurações de conexão com o banco
    - **models:** Local onde fica a estrutura do objeto a ser persistido/buscado no banco
    - **routes:** Local onde ficam as configurações de rotas da aplicação
    - **templates:** Local onde ficam as páginas web da aplicação

