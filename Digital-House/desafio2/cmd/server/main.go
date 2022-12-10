package main

import (
	"database/sql"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"digital-house/cmd/server/handler"
	"digital-house/internal/dentist"
	inboundOrders "digital-house/internal/inbound_orders"

	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := connection()
	if err != nil {
		log.Fatal(err)
	}

	repositoryInboundOrders := inboundOrders.NewRepository(conn)
	serviceInboundOrders := inboundOrders.NewService(repositoryInboundOrders)

	repositoryDentists := dentist.NewRepository(conn)
	serviceDentists := dentist.NewService(repositoryDentists)

	r := gin.Default()

	handler.NewInboundOrders(r, serviceInboundOrders)
	handler.NewDentist(r, serviceDentists)

	r.Run()
}

func connection() (*sql.DB, error) {
	user := "USER_DB"
	pass := "PASS_DB"
	port := "PORT_DB"
	host := "HOST_DB"
	database := "DATABASE"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, database)
	log.Println("connection successful")
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	return conn, nil
}
