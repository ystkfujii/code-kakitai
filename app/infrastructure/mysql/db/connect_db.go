package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"github/code-kakitai/code-kakitai/util"
)

const maxRetries = 5
const delay = 5 * time.Second

var (
	once  sync.Once
	query *dbgen.Queries
	dbcon *sql.DB
)

// contextからQueriesを取得する。contextにQueriesが存在しない場合は、パッケージ変数からQueriesを取得する
func GetQuery(ctx context.Context) *dbgen.Queries {
	txq := util.GetQueries(ctx)
	if txq != nil {
		return txq
	}
	return query
}

func SetQuery(q *dbgen.Queries) {
	query = q
}

func GetDB() *sql.DB {
	return dbcon
}
func SetDB(d *sql.DB) {
	dbcon = d
}

func NewMainDB() {
	once.Do(func() {
		dbcon, err := connect()
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})
}

// dbに接続する：最大5回リトライする
func connect() (*sql.DB, error) {
	cfg := config.GetConfig().DB
	for i := 0; i < maxRetries; i++ {
		connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		fmt.Println(connect)
		db, err := sql.Open("mysql", connect)
		if err != nil {
			return nil, fmt.Errorf("could not open db: %w", err)
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("could not connect to db: %v", err)
		log.Printf("retrying in %v seconds...", delay/time.Second)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("could not connect to db after %d attempts", maxRetries)
}
