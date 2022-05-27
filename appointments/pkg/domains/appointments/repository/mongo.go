package repository

import (
	"context"
	"fmt"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *MongoRepository) CreateAppointment(ctx context.Context, app model.Appointment) (*model.Appointment, error) {
	coll := m.client.Database(m.database).Collection(m.collection)
	result, err := coll.InsertOne(ctx, &app)
	if err != nil {
		return nil, errors.Wrap(appErr.ErrDatabase, err.Error())
	}

	app.ID = fmt.Sprintf("%v", result.InsertedID)
	return &app, nil
}

func (m *MongoRepository) UpdateAppointment(ctx context.Context, app model.Appointment) (*model.Appointment, error) {
	var updateApp model.Appointment
	id, err := primitive.ObjectIDFromHex(app.ID)
	if err != nil {
		return nil, errors.Wrap(appErr.ErrDatabase, err.Error())
	}

	app.ID = ""
	coll := m.client.Database(m.database).Collection(m.collection)
	if err := coll.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": &app}).Decode(&updateApp); err != nil {
		return nil, errors.Wrap(appErr.ErrDatabase, err.Error())
	}

	return &updateApp, nil
}

func (m *MongoRepository) DeleteAppointment(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(appErr.ErrDatabase, err.Error())
	}
	coll := m.client.Database(m.database).Collection(m.collection)
	result, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return err
	}

	return nil
}

func (m *MongoRepository) FindAllAppointments(ctx context.Context) ([]model.Appointment, error) {
	app := make([]model.Appointment, 0)
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
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var app model.Appointment
	coll := m.client.Database(m.database).Collection(m.collection)
	if err := coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&app); err != nil {
		return nil, err
	}
	return &app, nil
}

func (m *MongoRepository) MakeAppointment(ctx context.Context, id int) ([]model.Appointment, error) {
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
	app := make([]model.Appointment, 0)
	filter := bson.M{"user_id": id}
	coll := m.client.Database(m.database).Collection(m.collection)
	cursor, err := coll.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &app); err != nil {
		return nil, err
	}

	return app, nil
}

func (m *MongoRepository) FindAppointmentBySalonID(ctx context.Context, id int) ([]model.Appointment, error) {
	app := make([]model.Appointment, 0)
	filter := bson.M{"salon_id": id}
	coll := m.client.Database(m.database).Collection(m.collection)
	cursor, err := coll.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &app); err != nil {
		return nil, err
	}

	return app, nil
}

func (m *MongoRepository) AvaiableAppointment(ctx context.Context) ([]model.Appointment, error) {
	app := make([]model.Appointment, 0)
	coll := m.client.Database(m.database).Collection(m.collection)
	filter := bson.M{"user_id": 0}
	cur, err := coll.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &app); err != nil {
		return nil, err
	}
	return app, nil
}
