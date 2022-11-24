package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	pokemonpc "github.com/maksattur/grpc-pokemon/gen"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

var collection *mongo.Collection

type server struct {
	pokemonpc.PokemonServiceServer
}

type pokemonItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Pid         string             `bson:"pid"`
	Name        string             `bson:"name"`
	Power       []string           `bson:"power"`
	Description string             `bson:"description"`
}

func getPokemonData(data *pokemonItem) *pokemonpc.Pokemon {
	return &pokemonpc.Pokemon{
		Id:          data.ID.Hex(),
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}
}

func (*server) CreatePokemon(ctx context.Context, req *pokemonpc.CreatePokemonRequest) (*pokemonpc.CreatePokemonResponse, error) {
	pokemon := req.GetPokemon()

	data := &pokemonItem{
		Pid:         pokemon.GetPid(),
		Name:        pokemon.GetName(),
		Power:       pokemon.GetPower(),
		Description: pokemon.GetDescription(),
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can not convert to oid %v", err),
		)
	}

	return &pokemonpc.CreatePokemonResponse{
		Pokemon: &pokemonpc.Pokemon{
			Id:          oid.Hex(),
			Pid:         pokemon.GetPid(),
			Name:        pokemon.GetName(),
			Power:       pokemon.GetPower(),
			Description: pokemon.GetDescription(),
		},
	}, nil
}

func (*server) ReadPokemon(ctx context.Context, req *pokemonpc.ReadPokemonRequest) (*pokemonpc.ReadPokemonResponse, error) {
	pokemonID := req.GetId()
	oid, err := primitive.ObjectIDFromHex(pokemonID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("can not parse ID %v", err),
		)
	}
	data := &pokemonItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find pokemon with specified ID: %v", err),
		)
	}

	return &pokemonpc.ReadPokemonResponse{
		Pokemon: getPokemonData(data),
	}, nil
}

func (*server) UpdatePokemon(ctx context.Context, req *pokemonpc.UpdatePokemonRequest) (*pokemonpc.UpdatePokemonResponse, error) {
	pokemon := req.GetPokemon()
	oid, err := primitive.ObjectIDFromHex(pokemon.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("can not parse ID %v", err),
		)
	}

	data := &pokemonItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find pokemon with specified ID: %v", err),
		)
	}

	data.Pid = pokemon.GetPid()
	data.Name = pokemon.GetName()
	data.Power = pokemon.GetPower()
	data.Description = pokemon.GetDescription()

	_, err = collection.ReplaceOne(context.Background(), filter, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can not update object in MongoDB: %v", err),
		)
	}

	return &pokemonpc.UpdatePokemonResponse{
		Pokemon: getPokemonData(data),
	}, nil
}

func (*server) DeletePokemon(ctx context.Context, req *pokemonpc.DeletePokemonRequest) (*pokemonpc.DeletePokemonResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("can not parse ID %v", err),
		)
	}

	filter := bson.M{"_id": oid}
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can not delete object in MongoDB: %v", err),
		)
	}
	return &pokemonpc.DeletePokemonResponse{
		Id: oid.Hex(),
	}, nil
}

func (*server) ListPokemon(_ *pokemonpc.ListPokemonRequest, stream pokemonpc.PokemonService_ListPokemonServer) error {
	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("unknown internal error: %v", err))
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &pokemonItem{}

		if err := cur.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error while decoding data from MongoDB: %v", err),
			)
		}

		stream.Send(&pokemonpc.ListPokemonResponse{Pokemon: getPokemonData(data)})
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error: %v", err),
		)
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUrl := os.Getenv("MONGODB_URL")

	fmt.Println("Connecting to MongoDB")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pokemon Service Started")
	collection = client.Database("pokemondb").Collection("pokemon")

	lis, err := net.Listen("tcp", "0.0.0.0:4040")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pokemonpc.RegisterPokemonServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
