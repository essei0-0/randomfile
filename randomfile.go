package randomfile

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	txtExt = ".txt"
)

type FileType string

const (
	TXT FileType = "txt"
)

type Config struct {
	MaxSize int64
	MinSize int64
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateFile generates a file of the given fileType with random content and size
func GenerateFile(fileType FileType, config Config) (string, error) {
	var extension string
	switch fileType {
	case TXT:
		extension = txtExt
	default:
		return "", fmt.Errorf("unsupported file type: %v", fileType)
	}

	content := randomContent(config.MinSize, config.MaxSize)

	fileName := fmt.Sprintf("%s_%d%s", fileType, time.Now().Unix(), extension)
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

// ランダムな文字列を生成する
func randomContent(minSize, maxSize int64) string {
	size := rand.Int63n(maxSize-minSize+1) + minSize
	b := make([]rune, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// packageの初期化時に乱数のseedを設定する
func init() {
	rand.Seed(time.Now().UnixNano())
}
