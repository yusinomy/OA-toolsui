package common

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Printbody(r *http.Response) string {
	s := ""
	content, _ := ioutil.ReadAll(r.Body)
	s = fmt.Sprintf(string(content))
	return s
}

func Get(Target string, payload string) *http.Response {
	var resp *http.Transport
	if Proxy.Text != "" {
		uri, _ := url.Parse(Proxy.Text)
		resp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(uri),
		}
	} else {
		resp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{
		Timeout:   time.Duration(3) * time.Second,
		Transport: resp,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse //不允许跳转
		}}
	rsp, err := http.NewRequest("Get", Target+payload, nil)
	if err != nil {
		log.Println("请求失败")
	}
	rsp.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	r, err := client.Do(rsp)
	return r
}

func Post(Target string, payload string, exp string) (*http.Response, *http.Request) {
	var resp *http.Transport
	if Proxy.Text != "" {
		uri, _ := url.Parse(Proxy.Text)
		resp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(uri),
		}
	} else {
		resp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{
		Timeout:   time.Duration(3) * time.Second,
		Transport: resp,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse //不允许跳转
		}}
	rsp, err := http.NewRequest("POST", Target+payload, strings.NewReader(exp))
	if err != nil {
		log.Println(err)
	}
	rsp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	r, err := client.Do(rsp)
	return r, rsp
}
