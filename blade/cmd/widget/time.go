package widget

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
	"github.com/spf13/cobra"
)

type Time struct {
}

func NewTime() *Time {
	return &Time{}
}

func (t *Time) ServerTimeContrast(threshold int64) error {
	clockOffset, err := t.ntp()
	if err != nil {
		return fmt.Errorf("ntp: %w", err)
	}

	seconds := int64(clockOffset.Seconds()) // // 计算秒数

	remainingSeconds := clockOffset - time.Duration(seconds)*time.Second // 计算毫秒数（扣除秒数后剩余的部分转换为毫秒）
	milliseconds := remainingSeconds.Milliseconds()
	remainingMillis := remainingSeconds - time.Duration(milliseconds)*time.Millisecond // 计算微秒数（扣除秒数和毫秒数后剩余的部分转换为微秒）
	microseconds := remainingMillis.Microseconds()

	if threshold <= 0 {
		threshold = int64(100 * time.Millisecond)
	}

	if milliseconds > threshold {
		return fmt.Errorf("server time ahead: %d s %d ms %d us", seconds, milliseconds, microseconds)
	} else if milliseconds < -threshold {
		return fmt.Errorf("server time behind: %d s %d ms %d us", seconds, milliseconds, microseconds)
	}

	return nil
}

func (t *Time) ntp() (time.Duration, error) {
	// 向 NTP 服务器请求标准时间
	// pool.ntp.org 是一个全球自动负载均衡的 NTP 服务器地址
	// ntp.aliyun.com (阿里云 NTP)
	// ntp.tencent.com (腾讯云 NTP)
	// cn.ntp.org.cn (中国 NTP 服务器)
	// time.google.com (谷歌公共 NTP)
	// ntp.ntsc.ac.cn (中国科学院国家授时中心)
	options := ntp.QueryOptions{Timeout: 3 * time.Second}
	response, err := ntp.QueryWithOptions("ntp.aliyun.com", options)
	if err != nil {
		return 0, fmt.Errorf("failed to obtain standard time: %w", err)
	}

	return response.ClockOffset, nil
}

func InitialTime(parentCmd *cobra.Command) {
	parentCmd.AddCommand(&cobra.Command{
		Use: "__check-time",
		RunE: func(cmd *cobra.Command, args []string) error {
			return NewTime().ServerTimeContrast(0)
		},
	})
}
