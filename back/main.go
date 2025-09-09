package main()

import (
	"context"
	"log"
	"net"
	"os"

	"crosspromo-backend/models"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type server struct {
	repo *models.MongoDBRepository
}

func main() {
	// MongoDB connection
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	repo := models.NewMongoDBRepository(client)

	// gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	RegisterCrosspromoServiceServer(s, &server{repo: repo})

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}