package main

// @desc go os 包调用，运行系统命令等示例
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//osTest()
	//osCmdTest()
	startProcessTest()
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

// os/exec - 执行系统指令
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

// os.StartProcess - 启动外部二进制应用，是比较底层的 api，一般使用 os/exec
func startProcessTest() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// list file
	pid, err := os.StartProcess("/bin/ls", []string{"ls ~/", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)

}
