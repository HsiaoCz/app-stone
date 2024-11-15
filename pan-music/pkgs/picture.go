package pkg

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// 从默认的图片中随机获取一张作为用户的头像
func GetPicture(name string) string {
	// 将默认的图片路径放在一个env里面
	aPathStr := os.Getenv(name)
	// fmt.Println(aPathStr)
	aPath := strings.Split(aPathStr, " ")
	// fmt.Println(aPath)
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(aPath))
	// fmt.Println(i)
	return aPath[i]
}

