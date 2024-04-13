package storage

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

type FileInfo struct {
	Size         int64
	IsTermfiFile bool
	FileName     string
}

func GetFileInfo(uri string) FileInfo {
	size := getRemoteFileSize(uri)

	u, _ := url.Parse(uri)

	fileName := path.Base(u.Path)

	objectInfo := tryParseTermfiObjectInfo(fileName)

	fileInfo := FileInfo{
		Size:     size,
		FileName: fileName,
	}
	if objectInfo != nil {
		fileInfo.IsTermfiFile = true
		fileInfo.FileName = objectInfo.fileName
	}

	return fileInfo
}

func getRemoteFileSize(url string) int64 {
	headResp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Print(err.Error())
	}

	return int64(size)
}

func getLocalFileSize(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Print(err.Error())
	}

	return fi.Size(), nil
}
