package pkg

import (
	"bufio"
	"github.com/imroc/req/v3"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
)

// 取根目录地址
func GetRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("conf.Setup, fail to get current path: %v", err)
		os.Exit(0)

	}
	return dir
}

// 取配置文件路径
func GetConfileFile() string {
	// 配置文件路径 当前文件夹 + config.yaml
	configdir := path.Join(GetRootPath(), "config.yaml")
	if _, err := os.Stat(configdir); os.IsNotExist(err) {
		log.Fatalf("conf.Setup, fail to get config file: %v", err)
		os.Exit(0)
	}
	return configdir
}

// 生成随机的端口
func GetRandomPort() int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成20000到40000之间的随机整数

	return rand.Intn(20001) + 20000
}

// 获取验证码
func GetImg(fileContent string) string {
	// 输出编码后的字符串
	client := req.C()
	response, err := client.R().SetBody(fileContent).Post(Config.CaptchaServerurl)
	if err != nil {
		log.Fatalf("client.R().SetBody error:%v", err)
	}
	if !response.IsSuccessState() {
		log.Fatalf("response.IsSuccessState error:%v", err)
	}
	return response.String()
}

// 读取文本文件,传入文件名,返回文件内容
func ReadLines(filename string) ([]string, error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建一个Scanner对象
	scanner := bufio.NewScanner(file)

	// 定义一个字符串切片，用于保存文件中的每一行内容
	var lines []string

	// 循环读取文件中的每一行内容，并将其添加到字符串切片中
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// 返回字符串切片和nil
	return lines, nil
}

func ContainsString(arr []string, str string) bool {
	for _, val := range arr {
		if stringInString(str, val) {
			return true
		}
	}
	return false
}

func stringInString(s string, str string) bool {
	if len(s) > len(str) {
		return false
	}
	for i := 0; i <= len(str)-len(s); i++ {
		if str[i:i+len(s)] == s {
			return true
		}
	}
	return false
}
