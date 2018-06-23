package versionapp

import (
	"net/http"

	"github.com/labstack/echo"
)

//Value constante com a versão atual
const Value = "0.0.1"

//Data Estrutura de resposta da versao
type Data struct {
	Name    string `json:"app"`
	Version string `json:"version"`
}

//CurrentVersion Metodo de ação para rota "/"
func CurrentVersion(c echo.Context) error {
	v := &Data{
		Name:    "goponto",
		Version: Value,
	}
	return c.JSON(http.StatusOK, v)
}
