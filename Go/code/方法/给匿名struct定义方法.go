package main

import "sync"

// 下面是一个小trick。这个例子展示了简单的cache，其使用两个包级别的变量来实现，一个
// mutex互斥量(§9.2)和它所操作的cache：
var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// 和上面功能一致的struct写法
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func StructLookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
func main() {

}
