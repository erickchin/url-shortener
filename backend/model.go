package main

import (
	"database/sql"
	"log"
	"time"
)

type url struct {
	ID    		int     `json:"id"`
	UrlCode  	string  `json:"url_code"`
	ShortUrl 	string 	`json:"short_url"`
	OriginalUrl string 	`json:"original_url"`
	CreatedOn 	string  `json:"created_on"`
}

type urlLog struct {
	ID    		int     `json:"id"`
	UrlCode  	string  `json:"url_code"`
	IpAddress 	string 	`json:"ip_address"`
	AccessedOn 	string  `json:"accessed_on"`
}

/*
	GetUrlLogs - Gets the logs of the specific urlCode within the database

		urlCode - Unique url code for shortening

	Returns - an array of urlLog
*/
func GetUrlLogs(db *sql.DB, urlCode string) ([]urlLog, error) {
	rows, err := db.Query("SELECT id, url_code, ip_address, accessed_on FROM logs WHERE url_code = $1", urlCode)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	urlLogs := []urlLog{}

	for rows.Next() {
		var ul urlLog
		err := rows.Scan(&ul.ID, &ul.UrlCode, &ul.IpAddress, &ul.AccessedOn)
		if err != nil {
			return nil, err
		}
		urlLogs = append(urlLogs, ul)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return urlLogs, nil
}

/*
	CreateUrl - Inserts a row into urls table with its unique url
*/
func (u *url) CreateUrl(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO urls(id, url_code, short_url, original_url, created_on) VALUES(DEFAULT, $1, $2, $3, $4) RETURNING id",
		u.UrlCode, u.ShortUrl, u.OriginalUrl, time.Now().Format(time.RFC3339)).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

/*
	LogInstance - Inserts a log when a client has done a request
*/
func (l *urlLog) LogInstance(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO logs(id, url_code, ip_address, accessed_on) VALUES(DEFAULT, $1, $2, $3) RETURNING id",
		l.UrlCode, l.IpAddress, time.Now().Format(time.RFC3339)).Scan(&l.ID)

	if err != nil {
		return err
	}

	return nil
}

/*
	GetOriginalUrl - Runs a query on urls table with the short to recieve the original url
*/
func (u *url) GetOriginalUrl(db *sql.DB) error {
	return db.QueryRow("SELECT original_url FROM urls WHERE url_code=$1",
		u.UrlCode).Scan(&u.OriginalUrl)
}

