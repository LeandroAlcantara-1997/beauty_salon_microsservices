package repository

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoRepostory(client *mongo.Client, database, collection string) *MongoRepository {
	return &MongoRepository{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (m *MongoRepository) Create(ctx context.Context, app model.Appointment) (*model.Appointment, error) {
	var (
		result *mongo.InsertOneResult
		err    error
	)
	coll := m.client.Database(m.database).Collection(m.collection)
	if result, err = coll.InsertOne(ctx, &app); err != nil {
		return nil, err
	}

	app.ID = fmt.Sprintf("%v", result.InsertedID)
	return &app, nil
}

func (m *MongoRepository) Update(ctx context.Context, app model.Appointment) (*model.Appointment, error) {
	var (
		result       *mongo.SingleResult
		updateApp    model.Appointment
		updateResult *mongo.UpdateResult
		err          error
	)
	coll := m.client.Database(m.database).Collection(m.collection)
	if err = coll.FindOne(ctx, bson.M{"_id": app.ID}).Decode(&updateApp); err != nil {
		return nil, result.Err()
	}

	if updateResult, err = coll.UpdateByID(ctx, updateApp.ID, &app); err != nil {
		return nil, err
	}
	app.ID = fmt.Sprintf("%v", updateResult.UpsertedID)

	return &app, nil
}

func (m *MongoRepository) FindAll(ctx context.Context) ([]model.Appointment, error) {
	var app []model.Appointment
	coll := m.client.Database(m.database).Collection(m.collection)
	cur, err := coll.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &app); err != nil {
		return nil, err
	}

	return app, nil
}

func (m *MongoRepository) FindAppointmentByID(ctx context.Context, id string) (*model.Appointment, error) {
	var app model.Appointment
	coll := m.client.Database(m.database).Collection(m.collection)
	if err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&app); err != nil {
		return nil, err
	}
	return &app, nil
}

func (m *MongoRepository) MakeAppointment(ctx context.Context) ([]model.Appointment, error) {
	var app []model.Appointment
	coll := m.client.Database(m.database).Collection(m.collection)
	cursor, err := coll.Find(ctx, bson.M{"user_id": nil})
	if err != nil {
		return nil, err
	}

	if err := cursor.Decode(&app); err != nil {
		return nil, err
	}
	return app, nil
}

func (m *MongoRepository) FindAppointmentByUserID(ctx context.Context, id int) ([]model.Appointment, error) {
	var (
		cursor *mongo.Cursor
		app    []model.Appointment
		err    error
	)
	coll := m.client.Database(m.database).Collection(m.collection)
	if cursor, err = coll.Find(ctx, bson.M{"user_id": id}, options.Find()); err != nil {
		return nil, err
	}

	if err = cursor.Decode(&app); err != nil {
		return nil, err
	}

	return app, nil
}

func (m *MongoRepository) FindAppointmentBySalonID(ctx context.Context, id int) ([]model.Appointment, error) {
	var (
		cursor *mongo.Cursor
		app    []model.Appointment
		err    error
	)
	coll := m.client.Database(m.database).Collection(m.collection)
	if cursor, err = coll.Find(ctx, bson.M{"salon_id": id}, options.Find()); err != nil {
		return nil, err
	}

	if err = cursor.Decode(&app); err != nil {
		return nil, err
	}

	return app, nil
}
