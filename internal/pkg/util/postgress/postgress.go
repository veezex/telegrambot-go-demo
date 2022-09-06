package postgress

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgress struct {
	Host     string
	Port     uint64
	User     string
	Password string
	DBname   string
}

func (p *Postgress) CreatePool() (*pgxpool.Pool, func()) {
	ctx, cancel := context.WithCancel(context.Background())

	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.DBname)

	// connect to database
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		logrus.Fatal("can't connect to database", err)
	}

	if err := pool.Ping(ctx); err != nil {
		logrus.Fatal("ping database error", err)
	}

	//// настраиваем
	//config := pool.Config()
	//config.MaxConnIdleTime = time.Minute
	//config.MaxConnLifetime = time.Hour
	//config.MinConns = 2
	//config.MaxConns = 4

	return pool, func() {
		cancel()
		pool.Close()
	}
}
