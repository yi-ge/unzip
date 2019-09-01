// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package unzip

import "os"

func GetOsVersion() float32 {
	return 0
}

// FileIsExist -判断文件是否存在  存在返回 true 不存在返回false
func FileIsExist(filename string) bool {
	var exist = true

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}

	return exist
}
