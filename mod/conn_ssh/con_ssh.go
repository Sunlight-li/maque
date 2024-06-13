package conn_ssh

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"golang.org/x/crypto/ssh"
)

func Conn_ssh(host, user, pass, cmd string, port int8) string {
	sshHost := host
	sshUser := user
	sshPassword := pass
	sshType := "password"
	sshPort := port
	log.Println(host, user, pass, port)
	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Println("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("创建ssh session 失败", err)
	}
	defer session.Close()
	//执行远程命令
	combo, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Println("远程执行cmd 失败", err)
	}
	log.Println("命令输出:", string(combo))
	return string(combo)
}

func Local_cmd(cmd string) string {
	cmds := exec.Command(cmd) // 执行本地的命令
	output, err := cmds.CombinedOutput()
	if err != nil {
		log.Println("命令执行失败:", err)
	}
	log.Println("命令输出：", string(output))
	return string(output)
}
