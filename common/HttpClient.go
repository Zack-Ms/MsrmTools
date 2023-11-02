package common

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
)

/*
 * @desc    http请求工具类
 * @author 	Zack
 * @date 	2020/9/22 11:58
 ****************************************
 */

/*
*
发送post请求
- reqUrl：请求域名 + 接口地址（URI 域名后缀）
- paramsStr：参数 按照&拼接字符串
返回数据
- 数据
*/
func DoPost(reqUrl, paramsStr string) ([]byte, error) {
	req, err := http.NewRequest("POST", reqUrl, bytes.NewReader([]byte(paramsStr)))
	if err != nil {
		return nil, errors.New("POST 构建请求异常")
	}
	// 这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("POST 发送数据异常")
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("POST 接收数据异常")
	}
	return respBytes, nil
}

/*
*
发送Get请求
- reqUrl：请求域名 + 接口地址（URI 域名后缀）+ 参数串
返回数据
- 数据
*/
func DoGet(reqUrl string) ([]byte, error) {
	return DoGetByCustom(reqUrl, "application/x-www-form-urlencoded")
}

func DoGetByCustom(reqUrl, contentType string) ([]byte, error) {
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, errors.New("GET 构建请求异常")
	}
	contentType = contentType + "; charset=utf-8"
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GET 发送数据异常")
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("GET 接收数据异常")
	}
	return respBytes, nil
}

func DoGetParams(remoteUrl string, queryValues url.Values) (body []byte, err error) {
	client := &http.Client{}
	body = nil
	uri, err := url.Parse(remoteUrl)
	if err != nil {
		return
	}
	if queryValues != nil {
		values := uri.Query()
		if values != nil {
			for k, v := range values {
				queryValues[k] = v
			}
		}
		uri.RawQuery = queryValues.Encode()
	}
	reqest, err := http.NewRequest("GET", uri.String(), nil)
	reqest.Header.Add("Content-Type", "application/json")

	response, err := client.Do(reqest)
	defer response.Body.Close()
	if err != nil {
		return
	}

	if response.StatusCode == 200 {
		respBytes, _ := io.ReadAll(response.Body)
		return respBytes, nil
	}
	return
}
