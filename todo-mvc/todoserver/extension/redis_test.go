package extension

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/cast"
)

func TestSet(t *testing.T) {
	reply, err := RedisClient.Do("SET", "key1", "value1")
	if err != nil {
		panic(err)
	}
	fmt.Println("replyType", reflect.TypeOf(reply).String())
	fmt.Println("setReply", reply)
}

func TestGet(t *testing.T) {
	reply, err := RedisClient.Do("Get", "key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("replyType", reflect.TypeOf(reply).String())
	fmt.Println("getReply", cast.ToString(reply))
}
