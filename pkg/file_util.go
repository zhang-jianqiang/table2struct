package pkg

import (
	"errors"
	"os"
)

func OutputFile(path, content string) error {
	// 先判断文件是否存在
	// _, err := os.Stat(path)
	// if err != nil {
	// 	return err
	// }
	_, err := os.Lstat(path)
	if err == nil {
		return errors.New("文件已存在")
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}
