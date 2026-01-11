package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// GetShell 返回当前操作系统的 shell 和参数
func GetShell() (shell string, args []string) {
	if runtime.GOOS == "windows" {
		return "cmd", []string{}
	}

	// 优先使用环境变量中的 SHELL
	if envShell := os.Getenv("SHELL"); envShell != "" {
		return envShell, []string{}
	}

	// 尝试按优先级查找可用的 shell
	shells := []string{"/bin/bash", "/bin/zsh", "/bin/sh"}
	for _, sh := range shells {
		if _, err := os.Stat(sh); err == nil {
			return sh, []string{}
		}
	}

	// 最后回退到 sh（应该总是存在）
	return "sh", []string{}
}

// GetShellCommand 返回执行命令的 shell 和参数
func GetShellCommand(command string) (shell string, args []string) {
	shell, _ = GetShell()
	if runtime.GOOS == "windows" {
		return shell, []string{"/c", command}
	}
	return shell, []string{"-c", command}
}

// NewShellCmd 创建一个交互式 shell 命令
func NewShellCmd() *exec.Cmd {
	shell, _ := GetShell()
	if runtime.GOOS == "windows" {
		return exec.Command(shell)
	}
	// Unix 系统使用 -i 启用交互模式
	return exec.Command(shell, "-i")
}

// NewShellCommandCmd 创建一个执行指定命令的 shell 命令
func NewShellCommandCmd(command string) *exec.Cmd {
	shell, args := GetShellCommand(command)
	return exec.Command(shell, args...)
}
