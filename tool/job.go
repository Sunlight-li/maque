package tool

import (
	"time"
	"xuntian/conf"
	"xuntian/mod/check_up"
)

// 健康检查
func Check_Job(v string) {
	timer := time.NewTimer(time.Duration(conf.Config_data(v).Check_up.CheckSpeed) * time.Second)
	for {
		timer.Reset(time.Duration(conf.Config_data(v).Check_up.CheckSpeed) * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			if check_up.Check(v) == -1 {
				return
			}

		}
	}

}
