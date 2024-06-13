package conf

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

// 读取配置文件基础配置
type Config_Basics struct {
	System_num int64
	System_id  []string
	Alarm_url  string
	Keyword    string
}

// 读取配置文件config
type Config struct {
	System_name string `json:"system_name"` // 系统名称
	// 日志规则
	Logs struct {
		LogFilePath      string   `json:"log_file_path"` // 日志文件路径
		LogRules         string   `json:"log_rules"`     // 日志文件内关键字规则（支持正则表达式）
		LogInfo          string   `json:"log_info"`      // 告警信息自定义，默认发送日志内的信息
		OnOff            bool     `json:"on-off"`        // 是否开启规则触发操作(关闭 false  开启 true)
		TriggerOperation struct { // 触发操作
			SshIp       string `json:"ssh_ip"`       // ssh远程登录ip
			SshPort     int8   `json:"ssh_port"`     // ssh远程登录port
			SshUser     string `json:"ssh_user"`     // ssh远程登录用户名
			SshPassword string `json:"ssh_password"` // ssh远程登录密码
			SshCmd      string `json:"ssh_cmd"`      // ssh远程需要执行的命令
			LocalCmd    string `json:"local_cmd"`    // 本地需要执行的命令
		} `json:"trigger_operation"` // 触发操作（仅限于linux系统）
	} `json:"log"` // 日志规则
	//业务健康检查
	Check_up struct {
		CheckSpeed   int    `json:"check_speed"` // 请求速度（秒）
		Check_time   int    `json:"check_time"`  // 超时时间
		CheckUrl     string `json:"check_url"`   // 业务url地址，支持http和https
		CheckType    string `json:"check_type"`  // 请求类型（GET和POST）
		Check_header []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"check_header"` // 设置请求头
		Check_body         string   `json:"check_body"`          // 请求体
		CheckCode          int      `json:"check_code"`          // 请求返回的code值
		CheckCustomization string   `json:"check_customization"` // 自定义请求返回的body
		CheckInfo          string   `json:"check_info"`          // 告警信息
		OnOff              bool     `json:"on-off"`              // 是否开启规则触发操作(关闭 false  开启 true)
		TriggerOperation   struct { // 触发操作（仅限于linux系统
			SshIp       string `json:"ssh_ip"`       // ssh远程登录ip
			SshPort     int8   `json:"ssh_port"`     // ssh远程登录port
			SshUser     string `json:"ssh_user"`     // ssh远程登录用户名
			SshPassword string `json:"ssh_password"` // ssh远程登录密码
			SshCmd      string `json:"ssh_cmd"`      // ssh远程需要执行的命令
			LocalCmd    string `json:"local_cmd"`    // 本地需要执行的命令
		} `json:"trigger_operation"`
	} `json:"check_up"`
	// 文件防篡改（md5）
	FileTamperProof struct {
		// FileDataPath     string   `json:"file_data_path"`  // 数据存储路径
		// FilePath         string   `json:"file_path"`       // 文件路径
		FolderPath        string   `json:"folder_path"`       // 文件目录路径
		File_white_list   []string `json:"file_white_list"`   // 白名单文件
		Folder_white_list []string `json:"folder_white_list"` // 白名单目录路径
		FileInfo          string   `json:"file_info"`         // 告警信息
		OnOff             bool     `json:"on-off"`            // 是否开启规则触发操作(关闭 false  开启 true)
		TriggerOperation  struct { // 触发操作（仅限于linux系统）
			SshIp       string `json:"ssh_ip"`       // ssh远程登录ip
			SshPort     int8   `json:"ssh_port"`     // ssh远程登录port
			SshUser     string `json:"ssh_user"`     // ssh远程登录用户名
			SshPassword string `json:"ssh_password"` // ssh远程登录密码
			SshCmd      string `json:"ssh_cmd"`      // ssh远程需要执行的命令
			LocalCmd    string `json:"local_cmd"`    // 本地需要执行的命令
		} `json:"trigger_operation"`
	} `json:"file_tamper_proof"`
	// linux命令执行监控
	Command struct {
		CommandRules     string   `json:"command_rules"` // 命令监控规则
		CommandInfo      string   `json:"command_info"`  // 告警信息
		OnOff            bool     `json:"on-off"`        // 是否开启规则触发操作(关闭 false  开启 true)
		TriggerOperation struct { // 触发操作（仅限于linux系统）
			SshIp       string `json:"ssh_ip"`       // ssh远程登录ip
			SshPort     int8   `json:"ssh_port"`     // ssh远程登录port
			SshUser     string `json:"ssh_user"`     // ssh远程登录用户名
			SshPassword string `json:"ssh_password"` // ssh远程登录密码
			SshCmd      string `json:"ssh_cmd"`      // ssh远程需要执行的命令
			LocalCmd    string `json:"local_cmd"`    // 本地需要执行的命令
		} `json:"trigger_operation"`
	} `json:"command"`
}

// 根据系统标识获取对应配置文件
func Config_data(a string) *Config {
	// 把配置文件读取到结构体上
	var config Config
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 设置配置文件的路径
	viper.AddConfigPath(".")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	// //将配置文件绑定到config上
	// viper.Unmarshal(&config)
	// log.Println("config: ", config)
	// log.Println(viper.Get("system_1"))
	// 将map转换为JSON字符串
	jsonData, _ := json.Marshal(viper.Get(a))
	// log.Println(string(jsonData))
	json.Unmarshal(jsonData, &config)
	return &config
}

// 获取配置文件基础数据
func Config_get() *Config_Basics {
	var config_Basics Config_Basics
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 设置配置文件的路径
	viper.AddConfigPath(".")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	// 将配置文件绑定到结构体上
	viper.Unmarshal(&config_Basics)
	return &config_Basics
}
