package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/mstefa/go-api/internal/platform/server"
	"github.com/mstefa/go-api/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "ml_app_user2"
	dbPass = "ml_app_user2"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "goapi"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
