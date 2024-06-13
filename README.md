# 麻雀
# 一款安全防护工具，适用于linux系统，有日志分析、业务健康检查、文件防篡改三大功能
### 程序在不断优化中，有好的建议请发送邮件或者加入qq和微信沟通群  
QQ群聊：194442843，点击链接加入群聊【麻雀-安全监控系统】：http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=dAA7qG3Ofckg15QZ1Ww9dLCQxVGbiVB6&authKey=hrByJ5ewpU5%2F7dL%2BeOz9DWin0cdWj83CG2QiHm0%2FbRJhyaXUtQPU3gV68eHbOjq0&noverify=0&group_code=194442843    
微信群聊：![Alt text](%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20240613103955.jpg)



## 程序配置文件

```yaml
## 欢迎大家使用体验有啥建议请以邮件形式发送至【1352113079@qq.com】
# 文件名: config.yaml
# 创建日期: 2024-06-07
# 作者: 天才的疯子

# 系统个数
system_num: 2

# 系统名称标识（必填项）
system_id: system_1
# 告警系统URL（仅仅支持钉钉webhook）
alarm_url: https://oapi.dingtalk.com/robot/send?access_token=XXXXX

# 钉钉关键字（不配置不能发送信息）
keyword: 【告警中心】

#系统名称标识
system_1: 
  # 系统名称
  system_name: 某某某系统
  # 日志规则
  log:
    # 日志文件路径
    log_file_path: aaa.log
    # 日志文件内关键字规则（支持正则表达式） (?i)error 忽略大小写在每条日志中搜索error关键字
    log_rules: (?i)error
    # 告警信息自定义，默认发送日志内的信息
    log_info: 
    # 是否开启规则触发操作(关闭 false  开启 true)
    on-off: false
    # 触发操作（仅限于linux系统）
    trigger_operation:
      # ssh远程登录ip
      ssh_ip: fanyu.online
      # ssh远程登录port
      ssh_port: 22
      # ssh远程登录用户名
      ssh_user: root
      # ssh远程登录密码
      ssh_password: 
      # ssh远程需要执行的命令 解释器(例如：/bin/bash或者/usr/bin/python3)
      ssh_cmd: whoami; ls -l; 
      # 本地需要执行的命令  解释器(例如：/bin/bash或者/usr/bin/python3)
      local_cmd:
  # 业务健康检查
  check_up:
    # 请求速度（秒）
    check_speed: 3
    # 超时时间（秒）
    check_time: 5
    # 业务url地址，支持http和https
    check_url: http://127.0.0.1:50001/api/v1/wtch
    # 请求类型（GET和POST）
    check_type: POST
    # 设置请求头
    check_header: 
      - name: "Content-Type"
        value: "application/json"
      # - name: "Accept"
      #   value: "application/json"
    # 请求体（可以为空）
    check_body: "{\"token\":\"1234\",\"imgurl\":\"\",\"text\":\"haha\",\"username\":\"愚蠢的土拨鼠\",\"usernick\":\"\",\"type\":\"1\"}"
    # 请求返回的code值 
    check_code: 200
    # 自定义请求返回的body
    check_customization: 
    # 告警信息
    check_info:
    # 是否开启规则触发操作(关闭 false  开启 true)
    on-off: false
    # 触发操作（仅限于linux系统）
    trigger_operation:
      # ssh远程登录ip
      ssh_ip:
      # ssh远程登录port
      ssh_port:
      # ssh远程登录用户名
      ssh_user:
      # ssh远程登录密码
      ssh_password:
      # ssh远程需要执行的命令 解释器(例如：/bin/bash或者/usr/bin/python3)
      ssh_cmd:
      # 本地需要执行的命令  解释器(例如：/bin/bash或者/usr/bin/python3)
      local_cmd:
  # 文件防篡改（md5）
  file_tamper_proof:
    # 文件目录路径
    folder_path: "."
    # 白名单文件和路径
    file_white_list: [mod\file_tamper_proof\file_proof.go,config.yaml]
    # 告警信息
    file_info: 
    # 是否开启规则触发操作(关闭 false  开启 true)
    on-off: false
    # 触发操作（仅限于linux系统）
    trigger_operation:
      # ssh远程登录ip
      ssh_ip:
      # ssh远程登录port
      ssh_port:
      # ssh远程登录用户名
      ssh_user:
      # ssh远程登录密码
      ssh_password:
      # ssh远程需要执行的命令 解释器(例如：/bin/bash或者/usr/bin/python3)
      ssh_cmd:
      # 本地需要执行的命令  解释器(例如：/bin/bash或者/usr/bin/python3)
      local_cmd:
  # linux命令执行监控
  command:
    # 命令监控规则
    command_rules:
    # 告警信息
    command_info:
    # 是否开启规则触发操作(关闭 false  开启 true)
    on-off: false
    # 触发操作（仅限于linux系统）
    trigger_operation:
      # ssh远程登录ip
      ssh_ip:
      # ssh远程登录port
      ssh_port:
      # ssh远程登录用户名
      ssh_user:
      # ssh远程登录密码
      ssh_password:
      # ssh远程需要执行的命令 解释器(例如：/bin/bash或者/usr/bin/python3)
      ssh_cmd:
      # 本地需要执行的命令  解释器(例如：/bin/bash或者/usr/bin/python3)
      local_cmd:
```