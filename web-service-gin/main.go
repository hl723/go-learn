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

// albums slice for initializing album data
var albums = []album {
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums will return a JSON of all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds the received JSON album to our albums slice
func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind received JSON to new album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	// add the new album to albums
	albums = append(albums, newAlbum)

	// Return 201 status code and return new Album
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// return the album specified by ID
// ID is parameter from client
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, record := range albums {
		if record.ID == id {
			c.IndentedJSON(http.StatusOK, record) 
			return
		}
	}	

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
