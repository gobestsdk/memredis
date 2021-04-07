package memredis

import (
	"sync"
	"time"
)

var (
	memredis   sync.Map
	expireddic               = map[string]int64{}
	Cleanhz    time.Duration = time.Millisecond * 250
)

func init() {

	go func() {
		for {
			time.Sleep(Cleanhz)
			clean()
		}
	}()
}

func clean() {
	t := time.Now().Unix()
	for k, v := range expireddic {
		if t > v {
			memredis.Delete(k)
			delete(expireddic, k)
		}
	}
}

func GET(key string) (value interface{}, ok bool) {
	if expireddic[key] < time.Now().Unix() {
		memredis.Delete(key)
		return nil, false
	}
	return memredis.Load(key)
}

//GETString 当key不存在时，返回nil
func GETString(key string) (value *string) {
	v, ok := GET(key)
	if ok {
		var k = v.(string)
		return &k
	} else {
		return nil
	}
}

//GETint 当key不存在时，返回0
func GETint(key string) (value *int) {
	v, ok := GET(key)
	if ok {
		var k = v.(int)
		return &k
	} else {
		return nil
	}
}

func SET(key string, value interface{}, expire time.Duration) {
	memredis.Store(key, value)
	if expire > 0 {
		expireddic[key] = time.Now().Add(expire).Unix()
	}
}
