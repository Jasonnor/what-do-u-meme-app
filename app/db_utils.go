package app

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func connectDB(driver string, login string) (*sql.DB, error) {
	var db *sql.DB
	db, openErr := sql.Open(driver, login)
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
	sqlFmtStr := strings.Join(sqlFormats, ",")
	sqlQuery := fmt.Sprintf(
		`
		SELECT
			meme.id,
			meme.title,
			meme.img_path,
			meme.about,
			tags.name
		FROM
			meme
		LEFT JOIN
			meme_tags
		ON
			meme.id = meme_tags.meme_id
		INNER JOIN
			tags
		ON
			meme_tags.tags_id = tags.id
		WHERE
			id IN(%s)
		`,
		sqlFmtStr)

	rows, queryErr := db.Query(sqlQuery, memeIds)
	if queryErr != nil {
		log.Fatal(queryErr)
		return memes, queryErr
	}
	defer rows.Close()

	var IDToMemeMap map[int]memeDetail
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
