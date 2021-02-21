// Package main is a server for handling API requests.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/SpencerWhitehead7/hall-o-bad-code/go_rest_api/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), "postgresql://spencer:@localhost:5432/no-manga") // any db really
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := r.Group("/api")

	api.GET("/search", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"search": "results"})
	})

	api.GET("/chapters", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"chapters": "by latest update"})
	})

	handler.Manga(api, dbpool)

	// you'd splt all .Groups out into their own handlers too
	mangaka := api.Group("/mangaka")
	mangaka.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mangaka": "by alphabetical"})
	})
	mangaka.GET("/:mangakaID", func(c *gin.Context) {
		mangakaID := c.Param("mangakaID")
		c.JSON(http.StatusOK, gin.H{"specific mangaka": "w/ their manga" + mangakaID})
	})

	magazine := api.Group("/magazine")
	magazine.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"magazines": "by alphabetical"})
	})
	magazine.GET("/:magazineID", func(c *gin.Context) {
		magazineID := c.Param("magazineID")
		c.JSON(http.StatusOK, gin.H{"specific magazine": "w/ its mangaka and manga" + magazineID})
	})

	r.Run() // listen and serve on 8080
}
