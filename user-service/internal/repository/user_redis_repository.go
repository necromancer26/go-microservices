package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/necromancer26/go-microservices/user-service/config"
	"github.com/necromancer26/go-microservices/user-service/internal/models"
)

var ctx = context.Background()

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository() *RedisUserRepository {
	return &RedisUserRepository{client: config.NewRedisClient()}
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
	userID, err := r.client.Incr(ctx, "next_user_id").Result()
	if err != nil {
		return err
	}
	user.ID = int(userID)
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.client.Set(context.Background(), fmt.Sprintf("user:%d", user.ID), data, 0).Err()
}

func (r *RedisUserRepository) Update(user *models.User) error {
	ctx := context.Background()
	key := fmt.Sprintf("user:%d", user.ID)

	// Check if the user exists
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return fmt.Errorf("user with ID %d does not exist", user.ID)
	}

	// Marshal the updated user to JSON
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Save the updated user to Redis
	return r.client.Set(ctx, key, data, 0).Err()
}

func (r *RedisUserRepository) Delete(id int) error {
	return r.client.Del(context.Background(), fmt.Sprintf("user:%d", id)).Err()
}
func (r *RedisUserRepository) FindAll() ([]*models.User, error) {
	ctx := context.Background()
	keys, err := r.client.Keys(ctx, "user:*").Result()
	if err != nil {
		return nil, err
	}

	var users []*models.User
	for _, key := range keys {
		data, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var user models.User
		err = json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
func (r *RedisUserRepository) FindByName(name string) (*models.User, error) {
	ctx := context.Background()
	keys, err := r.client.Keys(ctx, "user:*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		data, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var user models.User
		err = json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}

		if user.Name == name {
			return &user, nil
		}
	}

	return nil, nil
}
