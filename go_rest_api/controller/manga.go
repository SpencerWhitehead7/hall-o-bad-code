package controller

import (
	"fmt"
	"net/http"

	"github.com/SpencerWhitehead7/hall-o-bad-code/go_rest_api/repository"

	"github.com/gin-gonic/gin"
)

// Manga performs the business logic to responds to requests.
type Manga struct {
	repo *repository.Manga
}

// GetAll returns all manga in the manga table
func (cont *Manga) GetAll(c *gin.Context) {
	mangas, err := cont.repo.FindAll()
	if err != nil {
		fmt.Printf("Database call failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{"mangas": mangas})
}

// MangaFactory creates new Mangas.
func MangaFactory(repo *repository.Manga) *Manga {
	return &Manga{repo: repo}
}
