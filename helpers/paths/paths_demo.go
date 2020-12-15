package paths

import (
	"log"
	"path/filepath"
)

// GetPath 获取指定位置的绝对路径
func GetPath(currPath string) string {

	log.Println("获取当前程序执行路径")

	dir, err := filepath.Abs(filepath.Dir(currPath))
	if err != nil {
		log.Println(err)
		return ""
	}
	return dir
}
