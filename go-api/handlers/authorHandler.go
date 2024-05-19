package handlers

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"encoding/json"
	"go-api/internal/db"
)

func ListAuthors(w http.ResponseWriter, r *http.Request) {
	connStr := "postgresql://user:password@db:5432/testdb?sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DBに接続できません。", err)
	}

	queries := db.New(conn)
	getResult, err := queries.ListAuthors(context.Background())
	if err != nil {
		log.Fatal("失敗しました。", err)
	}

	fmt.Println(getResult)

	w.WriteHeader(http.StatusOK) // ステータスコードを設定
	json.NewEncoder(w).Encode(getResult)
	return
}
