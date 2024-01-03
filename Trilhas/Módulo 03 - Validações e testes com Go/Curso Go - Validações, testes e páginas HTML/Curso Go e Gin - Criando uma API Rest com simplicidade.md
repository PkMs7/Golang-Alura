# Tópicos do Curso: Go - desenvolvendo uma API Rest


## 1.0 Instalando e criando a primeira rota com Gin

- Validator: Pacote de validação Golang (gopkg.in/validator.v2)

- Validações foram feitas inicialmente nos modelos

## 2.0 Testes

- Testes: Podem ser feitos no Postman com Javascript na própria ferramenta.

- Testes com Go: Por convenção todas as funções de teste devem começar com a palavra chave "Test". 

    - Uma ferramenta nativa para testes do Go é o "testing"
    - O comando para realizar testes é "go test"
    - Um pacote que pode ser baixado para complementar os testes nativos do go é o Testify (github.com/stretchr/testify)

## 3.0 Testando os endpoints

- Mock: são objetos que simulam o comportamento de objetos reais de forma controlada. São normalmente criados para testar o comportamento de outros objetos.

- Os testes podem ser exibidos em formato de release, configurando no Setup com o gin.SetMode(gin.ReleaseMode)

- Para se executar um único teste específico o comando no CLI é `go test -run {nome da função do teste}`

## 4.0 Aprofundando em testes

- 

## 5.0 Deletando, editando e buscando alunos

- O Go também possui recursos para renderizar HTML.