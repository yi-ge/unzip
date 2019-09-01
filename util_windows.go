package unzip

import (
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

type OSVERSIONINFO struct {
	dwOSVersionInfoSize int32
	dwMajorVersion      int32
	dwMinorVersion      int32
	dwBuildNumber       int32
	dwPlatformId        int32
	szCSDVersion        [128]byte
}

func GetOsVersion() float32 {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	var os OSVERSIONINFO
	os.dwOSVersionInfoSize = int32(unsafe.Sizeof(os))
	GetVersionExA := kernel32.NewProc("GetVersionExA")
	rt, _, _ := GetVersionExA.Call(uintptr(unsafe.Pointer(&os)))
	if int(rt) == 1 {
		res, err := strconv.ParseFloat(strconv.Itoa(int(os.dwMajorVersion))+"."+strconv.Itoa(int(os.dwMinorVersion)), 32)
		if err != nil {
			return 0
		}
		return float32(res)
	}
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
