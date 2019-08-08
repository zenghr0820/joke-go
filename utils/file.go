package utils

import (
	"fmt"
	"log"
	"os"
)

// 检查文件是否存在并且创建文件
func CheckFileAndCreate(path string) (err error) {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else {
		fmt.Println("path not exists ", path)
		err := os.MkdirAll(path, 0711)

		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
			return err
		}
		return nil
	}
}
