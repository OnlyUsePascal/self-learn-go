package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// curl http://localhost:8080/albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//	curl http://localhost:8080/albums \
//	    --include \
//	    --header "Content-Type: application/json" \
//	    --request "POST" \
//	    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind to json
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// append to slices
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum) // return SUCCESS with new item
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
