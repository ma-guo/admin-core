package fileupload

import (
	"io"
	"net/http"
	"os"
)

func getContentType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 读取前 512 字节用于 MIME 检测
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	// 自动检测 Content-Type
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
