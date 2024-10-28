package handlers

import (
	"net/http"
	"tardis/internals/logger"
	"tardis/internals/storage"

	"github.com/gin-gonic/gin"
)

type userinput struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Welcome to Tardis, High Performance In-Memory Database",
	})
	return

}

func Create(ctx *gin.Context) {
	var input userinput

	ctx.ShouldBindBodyWithJSON(&input)

	if input.Key == "" || input.Value == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyOrValueParameter",
		})
		return
	}

	err := storage.Store.Put(input.Key, input.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error: Key Already Exists",
		})
		return
	}

	ctx.Status(http.StatusCreated)
	return

}
func Retrieve(ctx *gin.Context) {
	var user userinput
	ctx.ShouldBindBodyWithJSON(&user)

	if user.Key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyOrValueParameter",
		})
		return
	}

	value, err := storage.Store.Get(user.Key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error: InvalidKey",
		})

		return
	}

	//return value
	ctx.JSON(http.StatusOK, gin.H{
		user.Key: value,
	})

}

func Update(ctx *gin.Context) {
	var input userinput

	ctx.ShouldBindBodyWithJSON(&input)

	if input.Key == "" || input.Value == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyOrValueParameter",
		})
		return
	}

	err := storage.Store.Insert(input.Key, input.Value)
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
	if input.Key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "InvalidKeyParameter",
		})
		return
	}

	storage.Store.Delete(input.Key)
	ctx.Status(http.StatusOK)

}
