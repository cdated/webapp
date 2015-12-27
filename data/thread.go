package data

import (
	"database/sql"
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	db := db()
	defer db.Close()
	rows, err := db.Query("SELECT id, uuid, topic, user_id, " +
		"created_at FROM threads ORDER BY created_at DESC")

	if err != nil {
		return
	}

	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic,
			&th.UserId, &th.CreatedAt); err != nil {
			return
		}

		threads = append(threads, th)
	}

	rows.Close()
	return
}

func db() (database *sql.DB) {
	database, err := sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (thread *Thread) NumReplies() (count int) {
	db := db()
	defer db.Close()
	rows, err := db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)

	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}

	rows.Close()
	return
}
