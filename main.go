package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strings"
	"time"
	"tuyue/tuyue_tools/cptool/cp"
	"tuyue/tuyue_tools/cptool/utils"
)

// 模糊匹配，查找PID echo $(pgrep node)
// cptool -src=./GateServer/ -s=1000 -e=1000 -o=8850

var (
	h         bool
	source    string
	oldPort   int
	startPort int
	endPort   int
)

func init() {
	flag.BoolVar(&h, "h", false, "帮助")
	flag.StringVar(&source, "src", "", "源文件夹")
	flag.IntVar(&startPort, "s", 0, "网关服务起始端口")
	flag.IntVar(&endPort, "e", 0, "网关服务结束端口")
	flag.IntVar(&oldPort, "o", 0, "需要替换的网关端口")
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	if source == "" {
		log.Println("-src 源文件夹路径不能为空")
		return
	}

	if startPort == 0 {
		log.Println("-s 网关起始端口缺失")
		return
	}

	if endPort == 0 {
		log.Println("-e 网关起结束口缺失")
		return
	}

	if startPort > endPort {
		log.Println("网关起始端口不能小于结束端口")
		return
	}

	if oldPort == 0 {
		log.Println("-o 需要替换的网关端口参数为空")
		return
	}

	source = strings.ReplaceAll(source, "\\", "/")

	if strings.HasPrefix(source, "./") {
		source = strings.TrimPrefix(source, "./")
	}

	if strings.HasSuffix(source, "/") {
		source = strings.TrimSuffix(source, "/")
	}

	exists, _ := utils.PathExists(source)
	if !exists {
		log.Println("err", "源文件夹不存在")
		return
	}

	sourceExecute := fmt.Sprintf("%s%s", source, "/GateServer")
	fileInfo, err := os.Stat(sourceExecute)
	if err != nil {
		log.Println("os.Stat failed:", err)
		return
	}

	fileMode := fileInfo.Mode()
	log.Println("file_mode:", fileMode)

	for port := startPort; port <= endPort; port++ {
		s := time.Now()
		dest := fmt.Sprintf("%s_%d", source, port)
		// err := cp1.CopyDir(source, dest)
		err := cp.CopyDir(source, dest)
		if err == nil {
			configFile := fmt.Sprintf("%s%s", dest, "/config/default.lua")
			err = ReplaceConfigGatewayPort(configFile, oldPort, port)
			if err != nil {
				log.Println("[error]", "修改网关端口", err)
			}

			executeFile := fmt.Sprintf("%s%s", dest, "/GateServer")
			err = os.Chmod(executeFile, fileMode)
			if err != nil {
				log.Println("[error]", "修改可执行权限", err)
			}
		} else {
			log.Println("[error]", err)
		}
		log.Println(dest, time.Now().Sub(s))
	}
}

func ReplaceConfigGatewayPort(filename string, oldPort int, newPort int) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("[error]", err)
		return err
	}
	configContent := string(data)
	configContent = strings.Replace(configContent, fmt.Sprintf("%d", oldPort), fmt.Sprintf("%d", newPort), 1)
	err = ioutil.WriteFile(filename, []byte(configContent), 0666)
	if err != nil {
		log.Println("[error]", err)
		return err
	}
	return nil
}
