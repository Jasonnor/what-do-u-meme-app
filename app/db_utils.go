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

func mockConnectDB(driver string, login string) (*sql.DB, error) {
	var db *sql.DB
	return db, nil
}

func getMemesByIds(db *sql.DB, memeIds []int) ([]memeDetail, error) {
	var memes []memeDetail

	sqlFormats := make([]string, len(memeIds))
	sqlFmtStr := strings.Join(sqlFormats, ",")
	sqlQuery := fmt.Sprintf("SELECT title, img_path, about FROM meme WHERE id IN(%s)", sqlFmtStr)

	rows, queryErr := db.Query(sqlQuery, memeIds)
	if queryErr != nil {
		log.Fatal(queryErr)
		return memes, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		var meme memeDetail
		err := rows.Scan(&meme.Title, &meme.ImageURL, &meme.About)
		if err != nil {
			log.Fatal(err)
			return memes, err
		}
		memes = append(memes, meme)
	}

	rowErr := rows.Err()
	if rowErr != nil {
		log.Fatal(rowErr)
		return memes, rowErr
	}

	return memes, nil
}
