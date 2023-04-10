package ner_spliter

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/wrpromail/annotate-helper/pkg/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Run() {
	// 打开文件
	inputFile, err := os.Open("raw.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer inputFile.Close()

	// 读取文件内容
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		// 去除空行和注释行
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			// 去除序号和空格
			re := regexp.MustCompile(`^\d+.\s+`)
			content := re.ReplaceAllString(line, "")
			// 写入新文件
			log.Info(content)

			filename := filepath.Join("/Users/wangrui/go/src/github.com/wrpromail/annotate-helper/files/out1", utils.RandomString(10)+".txt")

			// 创建文件
			file, err := os.Create(filename)
			if err != nil {
				panic(err)
			}

			// 将内容写入文件
			writer := bufio.NewWriter(file)
			_, err = writer.WriteString(content)
			if err != nil {
				panic(err)
			}

			// 关闭文件
			writer.Flush()
			file.Close()
		}
	}

	// 检查 scanner 是否失败，返回错误信息
	if err := scanner.Err(); err != nil {
		return
	}
}
