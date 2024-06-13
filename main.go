package main

import (
	"log"
	"xuntian/conf"
	"xuntian/mod/file_tamper_proof"
	"xuntian/mod/log_s"
	"xuntian/tool"
)

func main() {
	// 设置日志前缀
	log.SetPrefix("[麻雀]")
	// 设置日志标志
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate | log.Ltime)
	log.Println(conf.Config_get().System_num)
	//读取配置文件中每个系统数据
	for _, v := range conf.Config_get().System_id {
		log.Println(v)
		log.Println(conf.Config_data(v))
		log.Println(conf.Config_data(v).Check_up.Check_body)
		// log.Println(conf.Config_data(v).Logs.LogFilePath, conf.Config_data(v).Logs.TriggerOperation.SshPassword)
		if conf.Config_data(v).Logs.LogFilePath != "" {
			go log_s.Log_a(conf.Config_data(v).Logs.LogFilePath, conf.Config_data(v).Logs.LogRules, v)
		}

		if conf.Config_data(v).Check_up.CheckUrl != "" {
			go tool.Check_Job(v)
		}

		if conf.Config_data(v).FileTamperProof.FolderPath != "" {
			go file_tamper_proof.MonitorFiles(conf.Config_data(v).FileTamperProof.FolderPath, v)
		}
		for {
		}
	}

}
