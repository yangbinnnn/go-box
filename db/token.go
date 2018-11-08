package db

import (
	"fmt"
	"time"
)

const TOKEN_PREFIX = "token"

func tokenKey(token string) string {
	return fmt.Sprintf("%s_%s", TOKEN_PREFIX, token)
}

func AddUserToken(token, uid string, expiredSecond int64) error {
	err := RedisDB.Set(tokenKey(token), uid, time.Duration(expiredSecond)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUserToken(token string) (uid string, ttl int64) {
	uid, err := RedisDB.Get(tokenKey(token)).Result()
	if err != nil {
		return "", 0
	}
	fttl, err := RedisDB.TTL(tokenKey(token)).Result()
	if err != nil {
		return "", 0
	}
	return uid, int64(fttl)
}
