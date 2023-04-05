package text

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"mymlops/annotate-helper/pkg/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileSplit interface{}

type LocalFileSplit struct {
	path string
	out  string
}

func checkPath(path string) error {

	// 检查路径是否存在
	_, err := os.Stat(path)
	if err != nil {
		// 如果路径不存在，创建目录
		if os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0644); err != nil {
				return err
			}
		}
	}

	// 检查路径是否是目录
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return errors.New("path exists and not directory")
	}

	return nil
}

func (l *LocalFileSplit) Split() {
	if err := checkPath(l.out); err != nil {
		log.Error(err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	// 读取文本文件内容
	bytes, err := ioutil.ReadFile(l.path)
	if err != nil {
		panic(err)
	}
	content := string(bytes)

	// 按行分割字符串
	lines := strings.Split(content, "\n")

	// 循环处理每一行内容
	for _, line := range lines {
		// 如果是注释或空行，跳过处理
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// 生成随机文件名
		filename := filepath.Join(l.out, utils.RandomString(10)+".txt")

		// 创建文件
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		// 将内容写入文件
		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(line)
		if err != nil {
			panic(err)
		}

		// 关闭文件
		writer.Flush()
		file.Close()

		// 输出文件名
		fmt.Printf("生成文件：%s\n", filename)
	}
}
