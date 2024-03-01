package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/sbhagate-infoblox/user.app/handler"
	"github.com/sbhagate-infoblox/user.app/pb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user_app"
)

func main() {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	dbSQL, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	defer dbSQL.Close()

	err = dbSQL.Ping()
	if err != nil {
		panic(err)
	}

	// sqlStatement := `INSERT INTO users(account_id, compartment_id, user_name) VALUES ('111', 'CMPRT1', 'Saksham')`
	// _, err = dbSQL.Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSQL,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	listner, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Println(err)
	}
	//s := grpc.NewServer()

	s, err := handler.NewGRPCServer(db)
	if err != nil {
		log.Println(err)
	}

	ser, err := handler.NewUsersServer()
	if err != nil {
		log.Println(err)
	}

	pb.RegisterUsersServer(s, ser)
	err = s.Serve(listner)
	if err != nil {
		log.Println(err)
	}
}
