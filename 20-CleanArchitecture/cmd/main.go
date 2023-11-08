package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/configs"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/event/handler"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/graph"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/grpc/pb"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/grpc/service"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/web/webserver"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// esse main sobe 3 servidores, fazendo a mesma coisa ou o mesmo Use Case.
	// WebServer, gRPC server, GraphQL server.
	// "TIP -> remember to run go with main.go and wire.go"

	// Configs
	configs, err := configs.LoadConfig(".") // aonde esta o arquivo de conf.
	if err != nil {
		panic(err)
	}

	// Banco de dados
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// RabbitMQ
	rabbitMQChannel := getRabbitMQChannel()

	// inicia o event dispatcher
	eventDispatcher := events.NewEventDispatcher()
	// registra um evento de OrderCreated, com o handler e recebe o RabbitMQ.
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// Use Case
	// Criando o UseCase - depende do DB, Event.
	// Aqui usamos o wire para gerar esse DI. // ver wire_gen.go
	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	// subindo o webserver
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	// registra o chi logger e os handlers
	// usamos o go para não travar o app aqui.
	go webserver.Start()

	// Server gRPC
	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	// abrir na porta tcp
	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	// subimos o server usando o go para não travar o app aqui.
	go grpcServer.Serve(lis)

	// no resolver passamos aqui o UseCase, depende dele.
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	// rotas
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	// aqui não colocamos o go, para o app não cair.
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	// TODO: need put this at var env.
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
