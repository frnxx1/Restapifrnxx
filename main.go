package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type error struct {
	Error string `json:"error"`
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

func main() {
	router := GetRouter()
	router.Run("localhost:8080")

}

func postAlbums(c *gin.Context) {
	var newAlbums album

	if err := c.BindJSON(&newAlbums); err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"Bad request"})
		return
	}
	albums = append(albums, newAlbums)
	c.IndentedJSON(http.StatusCreated, newAlbums)
}

func getAlbumId(c *gin.Context) {
	id := c.Param("id")

	for _, i := range albums {
		if i.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, error{"not found"})
}

func updateAlbumId(c *gin.Context) {
	id := c.Param("id")
	for a, i := range albums {
		if i.ID == id {
			c.BindJSON(&albums[a])
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, error{"not found"})

}

func deleteAlbumId(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, error{"not found"})

}

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumId)
	router.DELETE("/albums/:id", deleteAlbumId)
	router.PUT("/albums/:id", updateAlbumId)
	router.POST("/albums", postAlbums)
	return router
}
