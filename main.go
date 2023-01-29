package main

import (
	"fmt"
	"os/exec"
	"service-monitoring/sedemail"
	"strings"
)

func cmd() string {
	// 修改cmd编码为utf8
	cmdUtf8 := exec.Command("chcp", "65001")
	_, err := cmdUtf8.Output()
	if err != nil {
		fmt.Println(err)
	}

	// 调用cmd查询进程，进程名称例如：TIM.exe
	cmd := exec.Command("TASKLIST", "/FI", "IMAGENAME eq TIM1.exe")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	str := strings.TrimSpace(string(out)) //去除字符串两边空格

	if string(str) == "INFO: No tasks are running which match the specified criteria." {
		fmt.Println("未找到进程")
		st := "stop"
		return st
	} else {
		fmt.Println("进程运行中")
		st := "run"
		return st
	}
	return ""
}

func main() {
	if cmd() == "run" {
		sedemail.Sedmail()
	}
}
