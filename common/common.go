package common

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"
)

//ASCII排序 并链接成字符串
func AsciiSort(params []byte) (map[string]string, string, error) {
	data := make(map[string]string, 8)
	if err := json.Unmarshal(params, &data); err != nil {
		return nil, "", err
	}
	var key []string
	var content string
	for k, _ := range data {
		key = append(key, k)
	}
	sort.Strings(key)
	for _, v := range key {
		if data[v] == "" {
			continue
		}
		content += v + "=" + data[v] + "&"
	}
	content = strings.TrimRight(content, "&")
	return data, content, nil
}

//错误信息
func ErrMsg(msg string) error {

	return errors.New(msg)

}
