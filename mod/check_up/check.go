package check_up

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"xuntian/common"
	"xuntian/conf"
	"xuntian/mod/conn_ssh"
)

func Check(v string) int {
	// url := "http://example.com/health" // 替换为要检查的URL
	// 设置超时时间
	client := http.Client{
		Timeout: time.Second * time.Duration(conf.Config_data(v).Check_up.Check_time),
	}
	// 创建请求体
	payload := bytes.NewBufferString(conf.Config_data(v).Check_up.Check_body) // 根据API要求设置请求体
	req, err := http.NewRequest(conf.Config_data(v).Check_up.CheckType, conf.Config_data(v).Check_up.CheckUrl, payload)
	if err != nil {
		log.Println("创建"+conf.Config_data(v).Check_up.CheckType+"请求失败:", err)
		return -1
	}
	// 设置请求头
	// h = map[string]string{}
	for _, v := range conf.Config_data(v).Check_up.Check_header {
		req.Header.Set(v.Name, v.Value)
	}
	// req.Header.Add("Content-Type", "application/json")

	// 发起POST请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println(conf.Config_data(v).Check_up.CheckType+"请求失败:", err)
		return -1
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应体失败:", err)
	}

	log.Println("响应状态码:", resp.StatusCode, "响应体:", string(body))
	if conf.Config_data(v).Check_up.CheckCode != 0 {
		if resp.StatusCode == conf.Config_data(v).Check_up.CheckCode {
			log.Println("code返回值对应，服务状态正常")
		} else {
			log.Println("code返回值不对应，服务状态异常")
			if conf.Config_data(v).Check_up.CheckInfo != "" {
				reply := common.Send_info(conf.Config_data(v).Check_up.CheckInfo, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
				if reply == -1 {
					log.Println("发送失败")
				}
				// 触发操作
				if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
					// 执行触发操作
					if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
						a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
						a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
				}
			} else {
				reply := common.Send_info("code返回值不对应，服务状态异常", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
				if reply == -1 {
					log.Println("发送失败")
				}
				// 触发操作
				if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
					// 执行触发操作
					if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
						a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
						a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
				}
			}
		}
	}
	if conf.Config_data(v).Check_up.CheckCustomization != "" {
		log.Println("执行了")
		if string(body) == conf.Config_data(v).Check_up.CheckCustomization {
			log.Println("body返回值对应，服务状态正常")
		} else {
			log.Println("body返回值不对应，服务状态异常")
			if conf.Config_data(v).Check_up.CheckCustomization != "" {
				reply := common.Send_info(conf.Config_data(v).Check_up.CheckInfo, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
				if reply == -1 {
					log.Println("发送失败")
				}
				// 触发操作
				if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
					// 执行触发操作
					if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
						a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
						a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
				}
			} else {
				reply := common.Send_info("body返回值不对应，服务状态异常", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
				if reply == -1 {
					log.Println("发送失败")
				}
				// 触发操作
				if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
					// 执行触发操作
					if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
						a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
						a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
						reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "健康检查")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
				}
			}
		}
	}
	return 0
}
