package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USERNAME")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected.")

	// by artist
	artist := "John Coltrane"
	albums, err := albumByArtist(db, artist)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album by %v : %v \n", artist, albums)

	// by id
	album, err := albumById(db, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album by %v : %v \n", artist, album)

	// add a new item
	albId, err := addAlbum(db, Album{
		Title:  "lorem ipsum dolor sit amit",
		Artist: "john doe",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Add Album, insert with ID %d", albId)

}

func albumByArtist(db *sql.DB, artist string) ([]Album, error) {

	var albums []Album

	rows, err := db.Query("SELECT * from album where artist = ? ", artist)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist : %q: %v ", artist, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist : %q: %v ", artist, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArtist : %q: %v ", artist, err)
	}

	return albums, nil

}

func albumById(db *sql.DB, id int64) (Album, error) {

	var alb Album

	rows := db.QueryRow("SELECT * from album where id = ?", id)

	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumById : data not found")
		}
		return alb, fmt.Errorf("albumById : %v: %v", id, err)
	}

	return alb, nil

}

func addAlbum(db *sql.DB, alb Album) (int64, error) {

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,  ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil

}
