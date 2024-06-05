package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vk/internal/config"

	_ "github.com/lib/pq"
)

const (
	USERS              = "users"
	BOOKS              = "books"
	COLLECTIONS        = "collections"
	COLLECTION_BOOKS   = "collections_books"
	COMMENTS           = "comments"
	COLLECTION_RATINGS = "collection_ratings"
	PARSEDATE          = "02-01-2006"
)

func runMigrations(db *sql.DB, migrationsPath string) error {
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") {
			filePath := filepath.Join(migrationsPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("Error applying migrations: %v", err)
				return err
			}

			_, err = db.Exec(string(content))
			if err != nil {
				log.Fatalf("Error applying migrations: %v", err)
				return err
			}

			log.Printf("Applied migration: %s", file.Name())
		}
	}

	return nil
}

func InitializeDB() (*sql.DB, error) {
	info := config.GetConf().DB

	// urlPostgres := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%d",
	// 	info.User, info.Name, info.Password, info.Host, info.Port)

	// urlPostgres := "postgres://slashersdcat:I01OL8DaGFnfIfnFR2J0Zk1Bkhf6ynKg@dpg-cp5nbtn79t8c73f1m51g-a/bookshelf_lo2k"
	urlPostgres := "postgres://slashersdcat:JdSOaFfRlgLpZnp7PkZ2sKzWOgW0EDzZ@dpg-cpfr8lv79t8c73e8ndb0-a/bookshelf_57rn"

	db, err := sql.Open("postgres", urlPostgres)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		//log.Println(urlPostgres)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	if err := runMigrations(db, info.Path); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
		return nil, err
	}

	log.Print("Init database")
	return db, nil
}
