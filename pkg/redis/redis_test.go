package redis

import (
	"UAEPassDemo/pkg/config"
	"testing"
)

func TestRedisDB_SetAccessCode(t *testing.T) {
	MyRedis.InitRedisDB(config.LocalConfig.Redis.Address, config.LocalConfig.Redis.Password, 5)
	err := MyRedis.SetAccessCode("hahha", "lala", 0)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRedis_GetAccessCode(t *testing.T) {
	MyRedis.InitRedisDB(config.LocalConfig.Redis.Address, config.LocalConfig.Redis.Password, 5)
	accessCode, err := MyRedis.GetAccessCode("hahha")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(accessCode)
}
