package widget

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Time struct {
}

func NewTime() *Time {
	return &Time{}
}

func (t *Time) ServerTimeContrast() error {
	return nil

	//var genre string
	//
	//genre = "taobao"
	//st, err := t.taobao()
	//if err != nil {
	//	genre = "suning"
	//	st, err = t.suning()
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//nts := time.Now().UnixMilli()
	//dv := int64(1000 * 5)
	//ndv := nts - st.UnixMilli()
	//if ndv > dv {
	//	return fmt.Errorf("server time ahead[%s]: %d s", genre, ndv/1000)
	//} else if ndv < -dv {
	//	return fmt.Errorf("server time behind[%s]: %d s", genre, ndv/1000)
	//}
	//
	//return nil
}

func (t *Time) suning() (time.Time, error) {
	var st time.Time

	type apiTime struct {
		SysTime2 string `json:"sysTime2"`
		SysTime1 string `json:"sysTime1"`
	}

	// {"sysTime2":"2024-09-12 10:53:08","sysTime1":"20240912105308"}
	_resp, err := http.Get("http://quan.suning.com/getSysTime.do")
	if err != nil {
		return st, err
	}

	defer func() {
		_ = _resp.Body.Close()
	}()

	res, err := io.ReadAll(_resp.Body)
	if err != nil {
		return st, fmt.Errorf("%s: %v", string(res), err)
	}

	var _apiTime apiTime
	err = json.Unmarshal(res, &_apiTime)
	if err != nil {
		return st, fmt.Errorf("%s: %v", string(res), err)
	}

	loc, _ := time.LoadLocation("Local") //获取时区
	st, err = time.ParseInLocation(time.DateTime, _apiTime.SysTime2, loc)
	if err != nil {
		return st, fmt.Errorf("%s: %v", string(res), err)
	}

	return st, nil
}

func (t *Time) taobao() (time.Time, error) {
	var st time.Time

	type apiTime struct {
		Api  string   `json:"api"`
		V    string   `json:"v"`
		Ret  []string `json:"ret"`
		Data struct {
			T string `json:"t"`
		} `json:"data"`
	}

	// {"api":"mtop.common.getTimestamp","v":"*","ret":["SUCCESS::接口调用成功"],"data":{"t":"1677234514487"}}
	_req, err := http.NewRequest("GET", "https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp", nil)
	if err != nil {
		return st, err
	}

	_req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	_req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	_req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	_req.Header.Set("Cache-Control", "no-cache")
	_req.Header.Set("Connection", "keep-alive")
	_req.Header.Set("Host", "api.m.taobao.com")
	_req.Header.Set("Pragma", "no-cache")
	_req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	_resp, err := http.DefaultClient.Do(_req)
	if err != nil {
		return st, err
	}

	defer func() {
		_ = _resp.Body.Close()
	}()

	res, err := io.ReadAll(_resp.Body)
	if err != nil {
		return st, fmt.Errorf("%s: %v", string(res), err)
	}

	var _apiTime apiTime
	err = json.Unmarshal(res, &_apiTime)
	if err != nil {
		return st, fmt.Errorf("%s: %v", string(res), err)
	}

	tss := _apiTime.Data.T
	ts, err := strconv.ParseInt(tss, 10, 64)
	if err != nil {
		return st, fmt.Errorf("%s: %v", string(res), err)
	}

	return time.UnixMilli(ts), nil
}

func InitialTime(parentCmd *cobra.Command) {
	parentCmd.AddCommand(&cobra.Command{
		Use: "__check-time",
		RunE: func(cmd *cobra.Command, args []string) error {
			return NewTime().ServerTimeContrast()
		},
	})
}
