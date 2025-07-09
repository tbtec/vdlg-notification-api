package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	rdbms "github.com/tbtec/tremligeiro/internal/infra/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	migrationsLocalDir   = "./scripts/migrations"
	migrationsDockerDir  = "/app/migrations/"
	migrationsLocalPath  = "file://./scripts/migrations"
	migrationsDockerPath = "file:///app/migrations/"
	schemaName           = "vdlg_video"
)

type PostgreSQLConf struct {
	User   string
	Pass   string
	Url    string
	Port   int
	DbName string
}

func New(conf PostgreSQLConf) (rdbms.RDBMS, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", conf.Url, conf.User, conf.Pass, conf.DbName, conf.Port)

	slog.InfoContext(context.Background(), "Connection to PostgreSQL database...")

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{
			TablePrefix:   schemaName + ".",
			SingularTable: true,
		}})

	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return rdbms.RDBMS{}, err
	}
	slog.InfoContext(context.Background(), "✅ Connection to PostgreSQL database successfully")

	return rdbms.RDBMS{DB: db}, nil
}

func Migrate(conf PostgreSQLConf) error {
	var migrationsPath string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.User, conf.Pass, conf.Url, conf.Port, conf.DbName)

	slog.InfoContext(context.Background(), "Initializing migrations...")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return err
	}

	driver, err := migratePostgres.WithInstance(db, &migratePostgres.Config{
		DatabaseName: conf.DbName,
	})

	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return err

	}

	if _, err := os.Stat(migrationsLocalDir); os.IsNotExist(err) {
		migrationsPath = migrationsDockerPath
	} else {
		migrationsPath = migrationsLocalPath
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		conf.DbName, driver)
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return err
	}
	err2 := m.Up()
	if err2 != nil {
		if err2.Error() == "no change" {
			slog.InfoContext(context.Background(), "No migrations to apply...")
		} else {
			return err
		}
	}

	slog.InfoContext(context.Background(), "Finished migrations")

	return nil
}
