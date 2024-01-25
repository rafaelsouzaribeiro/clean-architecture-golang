package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/configs"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/infra/event/handle"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/infra/graphql/graph"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/infra/grpc/pb"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/infra/grpc/services"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/infra/web/webserver"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handle.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUserCase := NewListOrderUseCase(db)

	webserver := webserver.NewWebServer(config.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	webserver.AddHandler("/orders", webOrderHandler.List)
	fmt.Println("Starting web server on port", config.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := services.NewOrderService(*createOrderUseCase, *listOrderUserCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUserCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", config.GRAPHQLServerPort)
	http.ListenAndServe(":"+config.GRAPHQLServerPort, nil)

}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
