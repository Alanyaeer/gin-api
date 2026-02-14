package idgenerator

import (
	"log"
	"strconv"

	"github.com/sony/sonyflake/v2"
)

var snowflakeGenerator *sonyflake.Sonyflake

func init() {
	var err error
	snowflakeGenerator, err = sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic(err)
	}
}

// NextID 生成下一个唯一ID（单例模式）
func NativeNextID() int64 {
	id, err := snowflakeGenerator.NextID()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func NextID() string {
	id, err := snowflakeGenerator.NextID()
	if err != nil {
		log.Fatal(err)
	}
	return strconv.FormatInt(id, 10)
}
