package file_tamper_proof

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"xuntian/common"
	"xuntian/conf"
	"xuntian/mod/conn_ssh"

	"github.com/fsnotify/fsnotify"
)

// 文件对应目录
var Pathmap = make(map[string]string)

// 文件和路径白名单
var file_white = make(map[string]string)

func File_White(v, a string) bool {
	// 添加白名单文件和目录到map文件中
	for _, b := range conf.Config_data(v).FileTamperProof.File_white_list {
		if b == a {
			log.Println("该文件已加白，不告警：", b, a)
			return true
		} else {
			re := regexp.MustCompile(strings.Replace("^"+b, `\`, `\\`, -1))
			if re.MatchString(a) {
				log.Println("该路径已加白，下面所有文件不告警：", b, a)
				return true
			}
		}
	}
	return false
}
func MonitorFiles(path, v string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = monitorDirRecursively(watcher, path)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// log.Println(Pathmap)
				// log.Println(Pathmap[event.Name])
				// 文件白名单判断
				// log.Println(File_White(v, event.Name))

				// 白名单文件
				// log.Println(File_White(v, event.Name))
				if event.Op&fsnotify.Write == fsnotify.Write && !File_White(v, event.Name) {
					log.Println("写入", event.Name)
					if conf.Config_data(v).FileTamperProof.FileInfo != "" {
						reply := common.Send_info("conf.Config_data(v).FileTamperProof.FileInfo\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					} else {
						reply := common.Send_info("监控到文件:"+event.Name+"触发写入操作\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					// 触发操作
					if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
						// 执行触发操作
						if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
							a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
						if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
							a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
					}
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename && !File_White(v, event.Name) {
					log.Println("重命名", event.Name)
					if conf.Config_data(v).FileTamperProof.FileInfo != "" {
						reply := common.Send_info("conf.Config_data(v).FileTamperProof.FileInfo\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					} else {
						reply := common.Send_info("监控到文件:"+event.Name+"触发重命名操作\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					// 触发操作
					if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
						// 执行触发操作
						if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
							a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
						if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
							a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
					}
				}
				if event.Op&fsnotify.Create == fsnotify.Create && !File_White(v, event.Name) {
					log.Println("创建:", event.Name)
					if conf.Config_data(v).FileTamperProof.FileInfo != "" {
						reply := common.Send_info("conf.Config_data(v).FileTamperProof.FileInfo\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					} else {
						reply := common.Send_info("监控到文件:"+event.Name+"触发创建操作\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					// 触发操作
					if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
						// 执行触发操作
						if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
							a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
						if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
							a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
					}
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove && !File_White(v, event.Name) {
					log.Println("删除:", event.Name)
					if conf.Config_data(v).FileTamperProof.FileInfo != "" {
						reply := common.Send_info("conf.Config_data(v).FileTamperProof.FileInfo\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					} else {
						reply := common.Send_info("监控到文件:"+event.Name+"触发删除操作\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					// 触发操作
					if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
						// 执行触发操作
						if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
							a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
						if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
							a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
					}
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod && !File_White(v, event.Name) {
					log.Println("修改权限:", event.Name)
					if conf.Config_data(v).FileTamperProof.FileInfo != "" {
						reply := common.Send_info("conf.Config_data(v).FileTamperProof.FileInfo\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					} else {
						reply := common.Send_info("监控到文件:"+event.Name+"触发权限变更操作\n", conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
						if reply == -1 {
							log.Println("发送失败")
						}
					}
					// 触发操作
					if conf.Config_data(v).Logs.OnOff && conf.Config_data(v).Logs.TriggerOperation.SshIp != "" && conf.Config_data(v).Logs.TriggerOperation.SshPassword != "" && conf.Config_data(v).Logs.TriggerOperation.SshPort != 0 && conf.Config_data(v).Logs.TriggerOperation.SshUser != "" {
						// 执行触发操作
						if conf.Config_data(v).Logs.TriggerOperation.SshCmd != "" {
							a := conn_ssh.Conn_ssh(conf.Config_data(v).Logs.TriggerOperation.SshIp, conf.Config_data(v).Logs.TriggerOperation.SshUser, conf.Config_data(v).Logs.TriggerOperation.SshPassword, conf.Config_data(v).Logs.TriggerOperation.SshCmd, conf.Config_data(v).Logs.TriggerOperation.SshPort)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
						if conf.Config_data(v).Logs.TriggerOperation.LocalCmd != "" {
							a := conn_ssh.Local_cmd(conf.Config_data(v).Logs.TriggerOperation.SshCmd)
							reply := common.Send_info("已执行触发操作, 执行结果返回:\n"+a, conf.Config_get().Alarm_url, conf.Config_get().Keyword, conf.Config_data(v).System_name, "文件防篡改")
							if reply == -1 {
								log.Println("发送失败")
							}
						}
					}
				}

				// 处理其他事件...
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	// 添加白名单文件和目录到map文件中
	for _, v := range conf.Config_data(v).FileTamperProof.File_white_list {
		file_white[v] = v
	}
	log.Println("文件目录加载完成，开始监控文件变化情况")
	log.Println(conf.Config_data(v).FileTamperProof.File_white_list)
	log.Println(conf.Config_data(v).FileTamperProof.Folder_white_list)
	fmt.Println("Watching... Press Ctrl+C to stop.")
	<-done // 等待停止信号
}

// 递归监控目录
// monitorDirRecursively 递归监控目录及其所有子目录
func monitorDirRecursively(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是不是目录
		if info.IsDir() {
			// 避免监控目录自身，因为已经在上层调用中添加了
			// if path != dir {
			err = watcher.Add(path)
			if err != nil {
				return err
			}
			// }
		}
		// 判断是不是文件
		if !info.IsDir() {
			log.Println(path)
			log.Println(strings.TrimSuffix(path, info.Name()))
			if path == info.Name() {
				log.Println("dir:", dir)
				Pathmap[path] = dir
			} else {
				Pathmap[path] = strings.TrimSuffix(path, info.Name())
			}

		}
		return nil
	})
}
