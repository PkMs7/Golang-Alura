package main

import (
	"bytes"
	"encoding/json"
	"go-rest-api-gin/controllers"
	"go-rest-api-gin/database"
	"go-rest-api-gin/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	ID int
)

func SetupDasRotasDeTeste() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {

	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Tigrinho", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz:":"O Tigrinho paga!"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))

	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d, e o esperado era %d", resposta.Code, http.StatusOK)
	// }

}

func TestListandoTodosOsAlunosHandler(t *testing.T) {

	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestListandoAlunosPorCpfHandler(t *testing.T) {

	CPFTeste := "00000000004"

	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCpf)
	pathDaBusca := "/alunos/cpf/" + CPFTeste
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestListandoAlunosPorIdHandler(t *testing.T) {

	IDTeste := 4

	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	pathDaBusca := "/alunos/" + strconv.Itoa(IDTeste)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestDeletandoAlunoHandler(t *testing.T) {

	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/excluindo-alunos/:id", controllers.ExcluirAluno)
	pathDaBusca := "/excluindo-alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestEditandoAlunoHandler(t *testing.T) {

	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/editando-alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47345678901", RG: "573456789"}
	valorJson, _ := json.Marshal(aluno)
	pathEdicao := "/editando-alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathEdicao, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, http.StatusOK, resposta.Code)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "47345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "573456789", alunoMockAtualizado.RG)

}
