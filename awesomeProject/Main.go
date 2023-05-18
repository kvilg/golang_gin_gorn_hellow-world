package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	Connect()

	var err error

	//err := DB.Migrator().DropTable(User{})
	//if err != nil {
	//	return
	//}

	//err = DB.Migrator().CreateTable(User{})
	//if err != nil {
	//	return
	//}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/sell", getSel)
	err = router.Run("localhost:8080")
	if err != nil {
		return
	}

}

func getSel(c *gin.Context) {

	var users []User

	DB.Model(&User{}).Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

func getAlbums(c *gin.Context) {
	users := []User{{Name: "qq"}}

	result := DB.Create(&users)

	err := result.Error // returns error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}

}
