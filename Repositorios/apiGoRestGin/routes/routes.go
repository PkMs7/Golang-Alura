package routes

import (
	"go-rest-api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.ExibePaginaIndex)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCpf)
	r.POST("/criando-alunos", controllers.CriarAluno)
	r.PATCH("/editando-alunos/:id", controllers.EditarAluno)
	r.DELETE("/excluindo-alunos/:id", controllers.ExcluirAluno)
	r.Static("/assets", "./assets")
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run(":8000")

}
