package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gin_healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"google.golang.org/grpc"
	"log"
	"microservice-uporabniki/controllers"
	docs "microservice-uporabniki/docs"
	"microservice-uporabniki/initializers"
	"microservice-uporabniki/middlewares"
	"microservice-uporabniki/models"
	pb "microservice-uporabniki/proto/users"
	"net"
	"sync"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToMysql()
	initializers.InitializeConsul()
}

var (
	grpcPort = flag.Int("grpc-port", 50051, "The gRPC server port")
)

type server struct {
	pb.UnimplementedRouteGuideServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	user := models.User{Name: in.GetName(), Email: in.GetEmail()}

	initializers.DB.Create(&user)
	return &pb.UserReply{Name: in.Name, Email: in.Email}, nil
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middlewares.MaintenanceMode())

	p := ginprometheus.NewPrometheus("gin")
	p.MetricsPath = "/users-service/metrics"

	serviceRoutes := r.Group("/users-service")
	{
		p.Use(r)

		docs.SwaggerInfo.BasePath = "/users-service"

		serviceRoutes.POST("/users", controllers.UsersCreate)
		serviceRoutes.GET("/users", controllers.UsersIndex)
		serviceRoutes.GET("/users/:id", controllers.UsersShow)
		serviceRoutes.PUT("/users/:id", controllers.UsersUpdate)
		serviceRoutes.DELETE("/users/:id", controllers.UsersDelete)

		serviceRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	config := gin_healthcheck.DefaultConfig()
	config.HealthPath = "/users-service/healthz"
	gin_healthcheck.New(r, config, []checks.Check{checks.SqlCheck{Sql: initializers.GetDb()}})

	// gRPC
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRouteGuideServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		s.Serve(lis)
		wg.Done()
	}()
	go func() {
		go r.Run()
		wg.Done()
	}()

	wg.Wait()
}
