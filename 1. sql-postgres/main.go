package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Album struct {
	ID		int64
	Title	string
	Artist	string
	Price	float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("select * from album where artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v\n", name, err)
	}
	defer rows.Close() // Defer closing rows so that any resources it holds will be released when the function exits

	for rows.Next() {
		var alb Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumsByArthist %q: %v\n", name, err)
		}

		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v\n", name, err)
	}

	return albums, nil
}

func albumById(id int64) (Album, error) {
	var album Album

	row := db.QueryRow("select * from album where id = $1", id)

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("AlbumById %d: Id does not exists", id)
		}

		return album, fmt.Errorf("AlbumById %d: %v\n", id, err)
	}

	return album, nil
}

func addAlbum(album Album) error {
	_, err := db.Exec(`insert into album (
		id, title, artist, price)
		values($1, $2, $3, $4);`,
		album.ID, album.Title, album.Artist, album.Price,
	)
	if err != nil {
		return fmt.Errorf("AddAlbum: %v\n", err)
	}

	return nil
}

func main() {
	conURL :=	"user=" + os.Getenv("DBUSER") + 
				" password="+ os.Getenv("DBPASS") + 
				" dbname=recordings sslmode=disable"

	var err error
	db, err = sql.Open("postgres", conURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database is connected")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Artist albums: %v\n", albums)

	// Create albums
	if err := addAlbum(Album{ID: 10, Title: "Programming in go", Artist: "Md. Istiyak Hossain", Price: 22.32,}); err != nil {
		log.Fatal(err)
	}

	album, err := albumById(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("One album data: %v\n", album)
}
