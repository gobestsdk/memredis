package memredis

func Keys(pattern string) (keys []string) {
	memredis.Range(func(key, value interface{}) (re bool) {
		k := key.(string)
		if pattern == "" || isMatch(k, pattern) {
			keys = append(keys, k)
		}
		return true
	})
	return keys
}
