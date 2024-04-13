package storage

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type termfiObjectInfo struct {
	fileName string
}

const termfiObjectRegex = "tf-[a-zA-Z0-9]+-(?<filename>.*)"

func createTermfiObjectName(filePath string) string {
	base := path.Base(filePath)
	id := generateFileId()

	return fmt.Sprintf("tf-%s-%s", id, base)
}

func tryParseTermfiObjectInfo(name string) *termfiObjectInfo {
	reg := regexp.MustCompile(termfiObjectRegex)
	res := reg.FindStringSubmatch(name)
	if len(res) != 2 {
		return nil
	}

	objectInfo := termfiObjectInfo{
		fileName: res[1],
	}

	return &objectInfo
}

func generateFileId() string {
	fullId := uuid.New()

	idString := strings.ReplaceAll(fmt.Sprintf("%s", fullId), "-", "")

	return idString[0:8]
}
