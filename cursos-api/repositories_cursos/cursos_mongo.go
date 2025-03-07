package repositories_cursos

import (
	"context"
	cursosDAO "cursos/dao_cursos"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Mongo struct {
	client     *mongo.Client
	database   string
	collection string
}

const (
	connectionURI = "mongodb://%s:%s"
)

func NewMongo(config MongoConfig) Mongo {
	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	ctx := context.Background()
	uri := fmt.Sprintf(connectionURI, config.Host, config.Port)
	cfg := options.Client().ApplyURI(uri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, cfg)
	if err != nil {
		log.Panicf("error connecting to mongo DB: %v", err)
	}

	return Mongo{
		client:     client,
		database:   config.Database,
		collection: config.Collection,
	}
}

func (repository Mongo) GetCourseByID(ctx context.Context, id string) (cursosDAO.Curso, error) {
	// Get from MongoDB
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error converting id to mongo ID: %w", err)
	}
	result := repository.client.Database(repository.database).Collection(repository.collection).FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error finding document: %w", result.Err())
	}

	// Convert document to DAO
	var cursoDAO cursosDAO.Curso
	if err := result.Decode(&cursoDAO); err != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error decoding result: %w", err)
	}
	return cursoDAO, nil
}

func (repository Mongo) Create(ctx context.Context, curso cursosDAO.Curso) (string, error) {
	result, err := repository.client.Database(repository.database).Collection(repository.collection).InsertOne(ctx, curso)
	if err != nil {
		return "", fmt.Errorf("error creating document: %w", err)
	}
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("error converting mongo ID to object ID")
	}
	return objectID.Hex(), nil
}

func (repository Mongo) Update(ctx context.Context, curso cursosDAO.Curso) error {
	// Convert curso ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(curso.Course_id)
	println(curso.Course_id)
	if err != nil {
		return fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	// Create an update document
	update := bson.M{}

	// Only set the fields that are not empty or their default value
	if curso.Nombre != "" {
		update["nombre"] = curso.Nombre
	}
	if curso.Categoria != "" {
		update["categoria"] = curso.Categoria
	}
	if curso.Descripcion != "" {
		update["descripcion"] = curso.Descripcion
	}
	if curso.Valoracion != 0 { // Assuming 0 is the default for Valoracion
		update["valoracion"] = curso.Valoracion
	}
	if curso.Duracion != 0 {
		update["duracion"] = curso.Duracion
	}
	if curso.Requisitos != "" {
		update["requisitos"] = curso.Requisitos
	}
	if curso.Url_image != "" {
		update["url_image"] = curso.Url_image
	}
	if !curso.Fecha_inicio.IsZero() {
		update["fecha_inicio"] = curso.Fecha_inicio
	}
	if curso.Capacidad != 0 {
		update["capacidad"] = curso.Capacidad
	}
	// Update the document in MongoDB
	if len(update) == 0 {
		return fmt.Errorf("no fields to update for curso ID %s", curso.Course_id)
	}

	filter := bson.M{"_id": objectID}
	result, err := repository.client.Database(repository.database).Collection(repository.collection).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("error updating document: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", curso.Course_id)
	}

	return nil
}

func (repository Mongo) GetCourses(ctx context.Context) (cursosDAO.Cursos, error) {
	cursor, err := repository.client.Database(repository.database).Collection(repository.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error getting documents: %w", err)
	}
	defer cursor.Close(ctx)

	var cursos cursosDAO.Cursos
	if err := cursor.All(ctx, &cursos); err != nil {
		return nil, fmt.Errorf("error decoding documents: %w", err)
	}
	return cursos, nil
}
