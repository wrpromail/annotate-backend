package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Run 将临时编写的文件转化为意图识别 json 列表
func Run() {
	// 打开文件
	file, err := os.Open("raw.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 逐行读取文件
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 && line[0] != '#' {
			lines = append(lines, line)
		}
	}

	// 拼接 json 对象
	var result []map[string]string
	for i := 0; i < len(lines); i += 2 {
		if i+1 >= len(lines) {
			break
		}
		obj := make(map[string]string)
		obj["text"] = lines[i]
		obj["intent"] = lines[i+1]
		result = append(result, obj)
	}

	// 输出结果
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
