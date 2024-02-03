package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID		int64
	Title	string
	Artist	string
	Price	float32
}

func albumsByArtist(name string) ([]Album, error){
	var albums []Album

	query, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("AlbumByArtist %q: %v", name, err)
	}
	defer query.Close() // This will release resources from memory


	for query.Next() {
		var album Album

		if err := query.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("AlbumByArtist %q: %v", name, err)
		}

		albums = append(albums, album)
	}

	if err := query.Err(); err != nil {
		return nil, fmt.Errorf("AlbumByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumById(id int64) (Album, error) {
	row := db.QueryRow("select * from album where id = ?", id)

	var alb Album
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("AlbumById %d: album not found", id)
		}
		return alb, fmt.Errorf("AlbumById %d: %v", id, err)
	}

	return alb, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec(`insert into album 
		(id, title, artist, price) 
		values (?, ?, ?, ?)`, 
		album.ID, album.Title, album.Artist, album.Price,
	)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	return id, nil
}

func main() {
	// Connection values
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",
		AllowNativePasswords: true,
	}

	// Get SQL database handler
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	// lets ping the db
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database is connected")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums are: %v\n", albums)

	// Create a row in the table
	id, err := addAlbum(Album{
		ID: 7,
		Title: "Go Programming",
		Artist: "Istiyak Hossain",
		Price: 100,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Fetch newly created album data
	album, err := albumById(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v\n", album)
}




