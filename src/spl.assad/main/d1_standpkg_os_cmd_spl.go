package main

// go os 包调用，运行系统命令等示例
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//osTest()
	//osCmdTest()
}

// os 包 - 获取系统信息，进行系统操作等
func osTest() {
	// 获取系统变量
	goPath := os.Getenv("GOPATH")
	fmt.Println(goPath)
	// 获取 pid
	pid := os.Getpid()
	fmt.Println(pid)
}

// os/exec - go 执行系统指令
func osCmdTest() {
	cmd := "ls -l"
	// 执行系统指令并获取输出，使用 bash shell 解释器
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println(err.Error())
	}
}
