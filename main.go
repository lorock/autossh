package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/lorock/autossh/core"
)

func main() {
	// 获取当前文件所在的目录绝对路径
	//	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	if err != nil {
	//		fmt.Println("error:", err)
	//		return
	//	}
	//	fmt.Println(path)
	// 如果你只对多个返回值里面的几个感兴趣
	// 可以使用下划线(_)来忽略其他的返回值
	// 取第一个返回值，忽略第二个值
	Hpath, _ := Home()
	app := core.App{ServersPath: Hpath + "/.servers.json"}
	app.Exec()
}

// Home 获取家目录
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

// homeUnix Unix
func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

// homeWindows windows系统
func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
