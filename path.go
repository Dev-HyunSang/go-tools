// (C) 2022 HyunSang Park <me@hyunsang.dev>
package gotools

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/huandu/xstrings"
)

// 디렉토리 내부에 있는 오브젝트를 읽습니다.
func DirectoryItemReader(path string) (string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, items := range files {
		return items.Name(), nil
	}

	return "", nil
}

// 파일의 확장자를 찾습니다.
// ex. helloworld.go => ".go"
func FindExt(fileName string) string {
	var copyExtensions string
	flag := 1
	for index := len(fileName) - 1; index >= 0; index-- {
		if fileName[index] == '.' {
			flag = 0
			break
		}
		copyExtensions += string(fileName[index])
	}
	if flag == 1 {
		return ""
	}
	copyExtensions = xstrings.Reverse(copyExtensions)

	return copyExtensions
}

// 파일을 삭제하는 함수입니다.
func DeleteFile(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}

// 현재의 경로를 알려줍니다.
// ex ~/dev/golang/hello-world
func Route() (path string) {
	path, err := os.Getwd()
	if err != nil {
		log.Print("Failed to Local PathWay Checking")
		os.Exit(1)
	}
	return path
}
