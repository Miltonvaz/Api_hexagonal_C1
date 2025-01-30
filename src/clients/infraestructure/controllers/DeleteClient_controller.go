package controllers

import (
	"ejercicio1/src/clients/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteClientController struct {
	usecase application.DeleteClient
}

func NewDeleteClientController(usecase application.DeleteClient) *DeleteClientController{
	return &DeleteClientController{usecase: usecase}
}

func (dc_c *DeleteClientController) Execute(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))

	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := dc_c.usecase.Execute(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}