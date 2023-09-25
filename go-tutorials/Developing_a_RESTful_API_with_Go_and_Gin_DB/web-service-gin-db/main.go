package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func getDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 "dev",
		Passwd:               "123456",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3305",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db, err
}

func initDbConnection() {
	_, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDbConnection()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	/*
		curl --location 'localhost:8080/albums' \
		--header 'Content-Type: application/json' \
		--data '{
			"title": "rowsAffected",
			"artist": "John Coltrane",
			"price": 56.99
		}'
	*/
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// 1. getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
			return
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// 2. postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "add new album error"})
		return
	}

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "add new album error"})
		return
	}

	id, err := result.LastInsertId()
	newAlbum.ID = strconv.Itoa(int(id))

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// 3. getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var album Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "get album by id error, id=" + id})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "get album by id error, id=" + id})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
