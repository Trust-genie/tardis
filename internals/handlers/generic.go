package handlers

import (
	"net/http"
	"tardis/internals/logger"
	"tardis/internals/storage"

	"github.com/gin-gonic/gin"
)

type userinput struct {
	key   string      `json:"binding:Key"`
	value interface{} `json:"binding: value"`
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Welcome to curl",
	})
	return

}

func Create(ctx *gin.Context) {
	var input userinput

	ctx.ShouldBindBodyWithJSON(&input)

	if input.key == "" || input.value == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyOrValueParameter",
		})
		return
	}

	err := storage.Store.Put(input.key, input.value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error: Key Already Exists",
		})
		return
	}

	ctx.Status(http.StatusCreated)
	return

}

func Update(ctx *gin.Context) {
	var input userinput

	ctx.ShouldBindBodyWithJSON(&input)

	if input.key == "" || input.value == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyOrValueParameter",
		})
		return
	}

	err := storage.Store.Insert(input.key, input.value)
	if err != nil {
		logger.Log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error: InternalServerError",
		})

		return
	}

	ctx.Status(http.StatusOK)
	return

}

func Delete(ctx *gin.Context) {
	var input userinput

	ctx.ShouldBindBodyWithJSON(&input)
	if input.key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyParameter",
		})
		return
	}

	storage.Store.Delete(input.key)
	ctx.Status(http.StatusOK)

}
