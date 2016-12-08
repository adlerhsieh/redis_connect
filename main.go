package main

import (
	"fmt"
)

func main() {
	r := RedisConnector{}
	r.Connect()
	defer r.Close()

	res := r.Get("foo")
	fmt.Println(res)

	// res = r.Rpush("xfoo", "bar")
	// fmt.Println(res)
}
