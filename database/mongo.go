package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "sppm/config"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectMongo() {
    uri := config.Env("MONGO_URI")

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("❌ Error creating MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("❌ Error connecting to MongoDB: %v", err)
    }

    
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("❌ MongoDB not responding: %v", err)
    }

    MongoClient = client
    MongoDB = client.Database(config.Env("MONGO_DB"))

    log.Println("✅ Connected to MongoDB!")
}
