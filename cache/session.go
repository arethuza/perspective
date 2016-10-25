package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/arethuza/perspective/misc"
	"gopkg.in/redis.v5"
	"time"
)

var client *redis.Client
var expiration time.Duration

func CreateRedisClient(config *misc.Config) error {
	redisAddr := fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort)
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	expiration, err = time.ParseDuration(config.RedisExpiration)
	return err
}

func CreateUserSession(config *misc.Config, user interface{}) (string, error) {
	token, b, err := misc.GenerateRandomString(config.TokenLength)
	if err != nil {
		return "", err
	}
	hasher := sha256.New()
	hasher.Write(b)
	tokenHash := hasher.Sum(nil)
	key := "perspective:session:" + hex.EncodeToString(tokenHash)
	value, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = client.Set(key, string(value), expiration).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}
