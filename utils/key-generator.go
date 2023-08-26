package utils

import "fmt"

type RedisKey string

func RedisKeyGenerator(predefinedKey RedisKey, ids []string) string {

	var key string = fmt.Sprintf("%v", predefinedKey)
	for _, id := range ids {
		key = fmt.Sprintf("%v#%v", key, id)
	}

	return key
}
