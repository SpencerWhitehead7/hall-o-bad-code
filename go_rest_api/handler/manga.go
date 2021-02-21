package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SpencerWhitehead7/hall-o-bad-code/go_rest_api/controller"
	"github.com/SpencerWhitehead7/hall-o-bad-code/go_rest_api/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Manga adds routes for retrieving information about all manga, one specific manga, or a specific page of a specific chapter of a specific manga.
func Manga(api *gin.RouterGroup, db *pgxpool.Pool) {

	cont := controller.MangaFactory(repository.MangaFactory(db))
	manga := api.Group("/manga")

	manga.GET("", func(c *gin.Context) {
		cont.GetAll(c)
	})

	manga.GET("/:mangaID", func(c *gin.Context) {
		mangaID, err := strconv.Atoi(c.Param("mangaID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"specific manga": fmt.Sprintf("w/ its chapters %v", mangaID)})
	})

	manga.GET("/:mangaID/:chapterNum/:pageNum", func(c *gin.Context) {
		mangaID, err := strconv.Atoi(c.Param("mangaID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		chapterNum, err := strconv.Atoi(c.Param("chapterNum"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		pageNum, err := strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"specific page": fmt.Sprintf("of chapter (static asset) %v %v %v", mangaID, chapterNum, pageNum)})
	})
}
