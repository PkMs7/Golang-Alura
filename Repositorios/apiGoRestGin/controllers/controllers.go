package controllers

import (
	"go-rest-api-gin/database"
	"go-rest-api-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Saudacao(c *gin.Context) {

	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "O " + nome + " paga!",
	})

}

func ExibeTodosOsAlunos(c *gin.Context) {

	//c.JSON(200, models.Alunos)

	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)

}

func ExibeAlunoPorId(c *gin.Context) {

	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"not_found": "Aluno não encontrado!"})
		return

	}
	c.JSON(http.StatusOK, aluno)

}

func ExibeAlunoPorCpf(c *gin.Context) {

	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func CriarAluno(c *gin.Context) {

	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)

}

func EditarAluno(c *gin.Context) {

	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}

func ExcluirAluno(c *gin.Context) {

	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno Deletado com sucesso"})

}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "notFound.html", nil)
}
