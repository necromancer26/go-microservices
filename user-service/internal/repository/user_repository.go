package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/necromancer26/go-microservices/user-service/config"
	"github.com/necromancer26/go-microservices/user-service/internal/models"
)

type UserRepository interface {
	FindAll() (map[string]string, error)
	FindByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}

var ctx = context.Background()

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository() *RedisUserRepository {
	return &RedisUserRepository{client: config.NewClient()}
}
func (r *RedisUserRepository) FindByID(id int) (*models.User, error) {
	key := fmt.Sprintf("user:%d", id)
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // not found
	} else if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = json.Unmarshal([]byte(val), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *RedisUserRepository) FindByEmail(email string) (*models.User, error) {
	val, err := r.client.Get(context.Background(), fmt.Sprintf("user:email:%s", email)).Result()
	if err == redis.Nil {
		return nil, nil // not found
	} else if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = json.Unmarshal([]byte(val), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *RedisUserRepository) Save(user *models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.client.Set(context.Background(), fmt.Sprintf("user:%d", user.ID), data, 0).Err()
}

func (r *RedisUserRepository) Update(user *models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.client.Set(context.Background(), fmt.Sprintf("user:%d", user.ID), data, 0).Err()
}

func (r *RedisUserRepository) Delete(id int) error {
	return r.client.Del(context.Background(), fmt.Sprintf("user:%d", id)).Err()
}
func (r *RedisUserRepository) FindAll() (map[string]string, error) {
	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	for _, key := range keys {
		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		values[key] = val
	}

	return values, nil
}
