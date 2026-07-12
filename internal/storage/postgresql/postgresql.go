package postgresql

import (
	"log/slog"

	"github.com/beyond-alok/paperwork/internal/config"
	"github.com/jackc/pgx"
)

func Connect(dbConfig config.DBConfig) (*pgx.Conn,error) {
	cfg := pgx.ConnConfig {
		Port : dbConfig.Port,
		Database: dbConfig.Database,
		User : dbConfig.User,
		Password: dbConfig.Password,
	}

	conn,err := pgx.Connect(cfg)
	if err != nil {
		slog.Error("database: postgresql failed to connect", "error",err)
		return nil,err
	}
	return conn,nil
}
