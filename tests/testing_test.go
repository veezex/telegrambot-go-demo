//go:build integration
// +build integration

package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	remote "gitlab.ozon.dev/veezex/homework/internal/pkg/api/grpc/v1/client"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
	postgressPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/storage/postgress"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/postgress"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/util/settings"
	pb "gitlab.ozon.dev/veezex/homework/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
	"sync"
	"testing"
)

var lock sync.RWMutex

func getConfig() settings.Config {
	return settings.New("../.env.test").GetConfig()
}

func setUpClient(t *testing.T) (storage.AppleStorage, func()) {
	t.Helper()
	config := getConfig()

	// grpc client init
	conn, err := grpc.Dial(fmt.Sprintf(":%d", config.GrpcServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// db control
	lock.Lock()
	pool, closeConnection := getDbPool(config)
	clearDb(pool)

	return remote.New(pb.NewAppleServiceClient(conn)), func() {
		conn.Close()

		// clear db
		clearDb(pool)
		closeConnection()
		lock.Unlock()
	}
}

func setUpDb(t *testing.T) (storage.AppleStorage, func()) {
	t.Helper()
	config := getConfig()
	pool, closeConnection := getDbPool(config)

	lock.Lock()
	clearDb(pool)
	return postgressPkg.New(pool), func() {
		clearDb(pool)
		closeConnection()
		lock.Unlock()
	}
}

func getDbPool(config settings.Config) (*pgxpool.Pool, func()) {
	// init postgress connection
	p := &postgress.Postgress{
		Host:     config.DbHost,
		Port:     config.DbPort,
		User:     config.DbUser,
		Password: config.DbPassword,
		DBname:   config.DbName,
	}
	return p.CreatePool()
}

func clearDb(pool *pgxpool.Pool) {
	rows, err := pool.Query(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		log.Fatal(err)
	}

	var tables []string
	for rows.Next() {
		value, err := rows.Values()
		if err != nil {
			log.Fatal(err)
		}

		tables = append(tables, value[0].(string))
	}
	rows.Close()

	if len(tables) == 0 {
		log.Fatal("You should run migrations first")
	}

	_, err = pool.Exec(context.Background(), fmt.Sprintf("Truncate table %s", strings.Join(tables, ",")))
	if err != nil {
		log.Fatal(err)
	}

}
