// (C) 2022 HyunSang Park <me@hyunsang.dev>
package gotools

import (
	"io/ioutil"
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

func RemoveFile(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
