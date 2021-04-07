package memredis



import (
	"sync"
	"time"
)

type MemRediskey struct {
	Name string
	//将键 key 的生存时间设置为 timestamp 所指定的毫秒数时间戳。
	PEXPIREAT int
	//恢复状态值
	//Recoverfunc func()
}
var (
	memredis sync.Map
	expireddic map[string]int64
	Cleanhz time.Duration=time.Millisecond*100
)
func init() {
	for {
		time.Sleep(Cleanhz)
		clean()
	}
}

func clean() {
		t := time.Now().Unix()
		for k, v := range expireddic {
			if t > v {
				memredis.Delete(k)
			}
		}
}

func GET(key string) (value interface{}, ok bool) {
	if expireddic[key]<time.Now().Unix(){
		memredis.Delete(key)
		return nil,false
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

func SET(key, value interface{}) {
	memredis.Store(key, value)
}


