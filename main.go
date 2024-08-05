package main

import (
	"fmt"

	cache "github.com/SergioPopovs176/mem-cache/Cache"
)

func main() {
	fmt.Println("Hello mem cache")

	cache := cache.New()
	fmt.Println(cache)

	cache.Set("number", 12345)
	cache.Set("string", "bla bla bla")
	cache.Set("float", 123.45)
	cache.Set("array", [3]string{"one", "two", "three"})
	fmt.Println(cache)

	v := cache.Get("string")
	fmt.Println(v)

	v = cache.Get("number")
	fmt.Println(v)

	fmt.Println(cache)

	cache.Delete("number")
	fmt.Println(cache)

	v = cache.Get("number")
	fmt.Println(v)
}
