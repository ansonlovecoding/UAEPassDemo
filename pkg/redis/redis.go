package redis

import (
	"UAEPassDemo/pkg/config"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

var MyRedis Redis

type Redis struct {
	client redis.UniversalClient
}

func init() {
	err := MyRedis.InitRedisDB(config.LocalConfig.Redis.Address, config.LocalConfig.Redis.Password, 5)
	if err != nil {
		panic(err)
	}
}

func (r *Redis) InitRedisDB(address, password string, timeout int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
		PoolSize: 100,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	r.client = client
	return nil
}

func (r *Redis) SetAccessCode(state, accessCode string, expiration int) error {
	key := fmt.Sprintf("%s_%s", UAEPassAccessCode, state)
	ok, err := r.client.SetNX(context.Background(), key, accessCode, time.Duration(expiration)*time.Second).Result()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("setting access code to redis was failed")
	}
	return nil
}

func (r *Redis) GetAccessCode(state string) (string, error) {
	key := fmt.Sprintf("%s_%s", UAEPassAccessCode, state)
	res, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (r *Redis) GetAccessToken(state string) (string, error) {
	key := fmt.Sprintf("%s_%s", UAEPassToken, state)
	res, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (r *Redis) SetAccessToken(state, token string, expiration int) error {
	key := fmt.Sprintf("%s_%s", UAEPassToken, state)
	_, err := r.client.Set(context.Background(), key, token, time.Duration(expiration)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}
