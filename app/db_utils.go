package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	connectString := `
		user=postgres 
		password=meme 
		host=35.192.115.150 
		dbname=meme-db 
		sslmode=disable 
	`
	db, openErr := sql.Open("postgres", connectString)
	if openErr != nil {
		log.Fatal(openErr)
		return db, openErr
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return db, err
	}

	return db, nil
}

func getMemesByIds(db *sql.DB, memeIds []int) ([]memeDetail, error) {
	var memes []memeDetail

	sqlFormats := make([]string, len(memeIds))
	for i := range sqlFormats {
		sqlFormats[i] = strconv.Itoa(memeIds[i])
	}
	sqlFmtStr := strings.Join(sqlFormats, ",")
	sqlQuery := fmt.Sprintf(
		`
		SELECT
			meme.id,
			meme.title,
			meme.image_path,
			meme.about,
			tag.name
		FROM
			meme
		LEFT JOIN
			meme_tag
		ON
			meme.id = meme_tag.meme_id
		INNER JOIN
			tag
		ON
			meme_tag.tag_id = tag.id
		WHERE
			meme.id IN(%s)
		`,
		sqlFmtStr)

	rows, queryErr := db.Query(sqlQuery)
	if queryErr != nil {
		log.Print(queryErr)
		return memes, queryErr
	}
	defer rows.Close()

	IDToMemeMap := make(map[int]memeDetail)

	for rows.Next() {
		var (
			id    int
			title string
			path  string
			about string
			tag   string
		)
		err := rows.Scan(
			&id,
			&title,
			&path,
			&about,
			&tag)
		if err != nil {
			log.Fatal(err)
			return memes, err
		}

		meme, exists := IDToMemeMap[id]
		if exists {
			meme.Tags = append(meme.Tags, tag)
		} else {
			tags := []string{tag}
			meme := memeDetail{
				ID:       id,
				Title:    title,
				ImageURL: path,
				About:    about,
				Tags:     tags,
			}
			IDToMemeMap[id] = meme
		}
	}

	rowErr := rows.Err()
	if rowErr != nil {
		log.Fatal(rowErr)
		return memes, rowErr
	}

	for _, meme := range IDToMemeMap {
		memes = append(memes, meme)
	}

	return memes, nil
}
