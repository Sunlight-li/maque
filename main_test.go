package main_test

import (
	"log"
	"regexp"
	"strings"
	"testing"
	"xuntian/common"
	"xuntian/conf"
	"xuntian/mod/check_up"
	"xuntian/mod/conn_ssh"
)

func TestGetdingding(t *testing.T) {
	reply := common.Send_info("你去过的最美的地方是哪", conf.Config_get().Alarm_url, conf.Config_get().Keyword, "营业系统", "日志规则")
	if reply == -1 {
		log.Println("发送失败")
	}
	// log.Println(conf.Config_get().Alarm_url, conf.Config_get().System_id, conf.Config_get().Keyword)
	t.Logf("%d\n", reply)
}

func TestLog(t *testing.T) {
	for _, v := range conf.Config_get().System_id {
		log.Println(v)
		log.Println(conf.Config_data(v))
		// log.Println(conf.Config_data(v).Logs.LogFilePath, conf.Config_data(v).Logs.LogRules)
		// go log_s.Log_a(conf.Config_data(v).Logs.LogFilePath, conf.Config_data(v).Logs.LogRules)
	}
}

func TestSsh(m *testing.T) {
	conn_ssh.Conn_ssh("fanyu.online", "root", "Lc753951..", "/bin/bash whoami;ls", 22)
	// conn_ssh.Local_cmd("dir")
}

func TestCheck(t *testing.T) {
	check_up.Check("system_1")
}

func TestFile_proof(t *testing.T) {
	b := `mod\`
	a := `mod\file_tamper_proof\file_proof.go`
	re := regexp.MustCompile(strings.Replace("^"+b, `\`, `\\`, -1))
	log.Println(re)
	log.Println(re.MatchString(a))
	if re.MatchString(a) {
		log.Println("true", b, a)
	}
}
