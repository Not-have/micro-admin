package utils

/**
 * io 是一个标准库包，提供了处理输入和输出操作的基础功能。
 */
import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Encryption(str string) string {

	md5 := md5.New()

	io.WriteString(md5, str)

	// return string(md5.Sum([]byte("")))

	// 或者下面的返回
	return fmt.Sprintf("%x", md5.Sum(nil))
}
