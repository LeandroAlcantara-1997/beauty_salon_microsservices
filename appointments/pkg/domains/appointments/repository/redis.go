package repository

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/go-redis/redis"
)

const (
	userID     = "user_"
	salonID    = "salon_"
	expiration = time.Hour
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(c *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: c,
	}
}

func (r *RedisRepository) CreateAppMemoryByID(app model.Appointment) error {
	appMemory, err := json.Marshal(app)
	if err != nil {
		return err
	}

	if re := r.client.Set(app.ID, appMemory, expiration); re.Err() != nil {
		return re.Err()
	}

	return nil
}

func (r *RedisRepository) CreateAppMemoryByUserID(app []model.Appointment) error {
	id := strconv.Itoa(app[0].UserID)

	byteApp, err := json.Marshal(&app)
	if err != nil {
		return err
	}

	if err := r.client.Set(fmt.Sprintf("%v%s", userID, id), byteApp, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisRepository) CreateAppMemoryBySalonID(app []model.Appointment) error {
	id := strconv.Itoa(app[0].SalonID)

	byteApp, err := json.Marshal(&app)
	if err != nil {
		return err
	}

	if err := r.client.Set(fmt.Sprintf("%v%s", salonID, id), byteApp, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisRepository) FindAppByIDMemory(id string) (*model.Appointment, error) {
	var app model.Appointment
	byteApp, err := r.client.Get(id).Bytes()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteApp, &app); err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *RedisRepository) FindAppByUserIDMemory(id int) ([]model.Appointment, error) {
	var app []model.Appointment
	byteApp, err := r.client.Get(fmt.Sprintf("%v%d", userID, id)).Bytes()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteApp, &app); err != nil {
		return nil, err
	}

	return app, nil
}

func (r *RedisRepository) FindAppBySalonIDMemory(id int) ([]model.Appointment, error) {
	var app []model.Appointment
	byteApp, err := r.client.Get(fmt.Sprintf("%v%d", salonID, id)).Bytes()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteApp, &app); err != nil {
		return nil, err
	}

	return app, nil
}
