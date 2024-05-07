package redirect

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10/orm"
	"github.com/sapphiregaze/yuri-server/internal/database"
)

type Redirect struct {
	ID           int    `pg:",pk"`
	RedirectPath string `pg:",unique,notnull"`
	OriginalURL  string `pg:",notnull"`
}

func CreateSchema() error {
	db := database.ConnectDB()
	models := []interface{}{
		(*Redirect)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
		defer db.Close()
	}

	return nil
}

func RedirectHandler(ctx *gin.Context) {
	redirectPath := ctx.Param("redirect")

	redirect := Redirect{}

	db := database.ConnectDB()
	err := db.Model(&redirect).
		Where("redirect_path = ?", redirectPath).
		Select()
	defer db.Close()

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Redirect path not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, redirect.OriginalURL)
}

func CreateRedirectHandler(ctx *gin.Context) {
	var redirect Redirect
	if err := ctx.BindJSON(&redirect); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := database.ConnectDB()

	_, err := db.Model(&redirect).
		OnConflict("DO NOTHING").
		Insert()
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert redirect"})
		return
	}

	redirectUrl := fmt.Sprintf("%s%s", os.Getenv("HOST"), redirect.RedirectPath)
	ctx.JSON(http.StatusCreated, gin.H{"message": redirectUrl})
}
