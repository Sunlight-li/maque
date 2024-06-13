package log_s

import (
	"log"
	"os"
	"regexp"
	"xuntian/common"
	"xuntian/conf"
	"xuntian/mod/conn_ssh"

	"github.com/hpcloud/tail"
)

// 日志规则
func Log_a(log_path, rul, v string) {
	//设置seek 到文末
	seek := &tail.SeekInfo{}
	seek.Offset = 0
	seek.Whence = os.SEEK_END
	//设置配置
	config := tail.Config{}
	config.Follow = true
	config.Location = seek

	t, err := tail.TailFile(log_path, config)
	if err != nil {
		log.Println(err)
	}
	for line := range t.Lines {
		re := regexp.MustCompile(rul)
		matches := re.FindString(line.Text)
		if matches != "" {
			log.Println("告警日志内容：", line.Text)
			if conf.Config_data(v).Logs.LogInfo != "" {
				reply := common.Send_info(conf.Config_data(v).Logs.LogInfo, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "日志规则")
				if reply == -1 {
					log.Println("发送失败")
				}
				if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
					// 执行触发操作
					if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
						a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "日志规则")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
						a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "日志规则")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
				}
			} else {
				reply := common.Send_info(line.Text, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "日志规则")
				if reply == -1 {
					log.Println("发送失败")
				}
				if conf.Config_data(v).Logs.OnOff {
					// 执行触发操作
					if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
						a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "日志规则")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
						a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "日志规则")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
				}
			}

		}
	}
}
