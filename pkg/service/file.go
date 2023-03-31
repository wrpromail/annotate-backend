package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mymlops/annotate-helper/pkg/constant"
	"mymlops/annotate-helper/pkg/dao"
	"os"
	"strconv"
	"strings"
)

func ListFile(c *gin.Context) {
	var result []dao.DataFile
	err := DDB.engine.Find(&result)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}

func GetFile(c *gin.Context) {
	id := c.Param("id")
	var result dao.DataFile
	has, err := DDB.engine.ID(id).Get(&result)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !has {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}
	c.JSON(200, result)
}

func GetFileOntology(c *gin.Context) {
	// id := c.Param("id")
	c.JSON(200, constant.DefaultOntology)
}

func GetFileLineCount(c *gin.Context) {
	id := c.Param("id")
	var result dao.DataFile
	has, err := DDB.engine.ID(id).Get(&result)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !has {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}
	rst := DLTM.GetLines(id)
	if len(rst) != 0 {
		c.JSON(200, gin.H{
			"count": len(rst),
		})
		return
	} else {
		rst, err := DLTM.ReadFile(id)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"count": len(rst),
		})
	}
}

type LocalTextFileMgr struct {
	cacheMap map[string][]string
}

func (l *LocalTextFileMgr) GetFileLine(fileId string, lineNumber int) (string, error) {
	if l.cacheMap == nil {
		l.cacheMap = make(map[string][]string)
		return "", errors.New("file not found")
	}
	if lines, ok := l.cacheMap[fileId]; ok {
		if lineNumber >= len(lines) {
			return "", errors.New("line number out of range")
		}
		return lines[lineNumber], nil
	}
	return "", errors.New("file not found")
}

func (l *LocalTextFileMgr) GetLines(fileId string) []string {
	if l.cacheMap == nil {
		l.cacheMap = make(map[string][]string)
		return []string{}
	}
	if lines, ok := l.cacheMap[fileId]; ok {
		return lines
	}
	return nil
}

func (l *LocalTextFileMgr) IsFileLineNumberValid(fileId string, lineNumber int) (bool, string) {
	if l.cacheMap == nil {
		l.cacheMap = make(map[string][]string)
		return false, ""
	}
	if lines, ok := l.cacheMap[fileId]; ok {
		if lineNumber >= len(lines) {
			return false, ""
		}
		return true, lines[lineNumber]
	}
	return false, ""
}

func readFileToLine(filepath string) (result []string, err error) {

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines []string
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}
	if err = sc.Err(); err != nil {
		fmt.Println(err)
		return
	}

	return lines, nil
}

func (l *LocalTextFileMgr) ReadFile(fileId string) (rst []string, err error) {
	var result dao.DataFile
	has, err := DDB.engine.ID(fileId).Get(&result)
	if err != nil {
		return
	}
	if !has {
		return rst, errors.New("文件记录不存在")
	}
	// 文件的 accessInfo 不存在
	if result.AccessInfo == "" {
		return rst, errors.New("文件记录没有关联实际访问信息")
	}
	var access = dao.FileAccess{Id: result.AccessInfo}
	has, err = DDB.engine.ID(result.AccessInfo).Get(&access)
	if err != nil {
		return rst, err
	}
	if !has {
		return rst, errors.New("文件访问信息不存在")
	}
	if access.Type == 1 {
		rst, err = readFileToLine(access.AccessInfo)
		if len(rst) > 0 {
			l.cacheMap[fileId] = rst
		}
		return
	} else {
		return rst, errors.New("暂不支持的文件访问类型")
	}
}

// todo 考虑在访问具体文件 line 时是否也加载文件
func GetFileLine(c *gin.Context) {
	id := c.Param("id")
	lineNumber := c.Param("number")
	ln, _ := strconv.Atoi(lineNumber)
	rst, err := DLTM.GetFileLine(id, ln)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, rst)
}

var DLTM *LocalTextFileMgr

func init() {
	DLTM = &LocalTextFileMgr{
		cacheMap: make(map[string][]string),
	}
}

type AnnotateRequest struct {
	Option string `json:"option"`
}

func FileLineAnnotate(c *gin.Context) {
	id := c.Param("id")
	lineNumber := c.Param("number")
	ln, err := strconv.Atoi(lineNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	valid, str := DLTM.IsFileLineNumberValid(id, ln)
	if !valid {
		c.JSON(500, gin.H{
			"error": "line number out of range",
		})
		return
	}

	var req AnnotateRequest
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.Option != "" {
		rst := &dao.Annotation{
			Id:         uuid.New().String(),
			Source:     str,
			Content:    req.Option,
			DataFileId: id,
		}
		_, err = DDB.engine.Insert(rst)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"result": "ok",
		})

	} else {
		c.JSON(500, gin.H{
			"error": "option is empty",
		})
	}

}
