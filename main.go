package main

import (
	"MyWebService/book"
	"MyWebService/book/data"
	"MyWebService/config"
	"MyWebService/lib/db"
	"MyWebService/lib/server"
	"flag"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	cfg := config.Config{Version: "1.0.0"}

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|staging|prod)")
	flag.StringVar(&cfg.DbDsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	fx.New(
		server.Module,
		book.Module,
		db.Module,
		fx.Supply(&cfg),
		fx.Provide(newLogger),
		fx.Invoke(func(db *gorm.DB) error {
			// Migrate the schema
			return db.AutoMigrate(&data.Book{}) // TODO: Collect data models ?
		}),
	).Run()
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "[LOG] ", log.LstdFlags)
}
