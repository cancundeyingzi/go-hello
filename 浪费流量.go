package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var totalDownloadedBytes int64 = 0
var lock sync.Mutex

// 下载文件并更新字节数
func run() {
	urls := []string{
		// "https://downloads.bdrive.com/netdrive/builds/0cfeb64848024a7ba4f28ce66549cbe2/NetDriveInstaller-3.17.730.dmg",
		// "https://op.yundasys.com/opserver/pages/downapp/image/img/bg-img.png",
		// "https://dldir1v6.qq.com/weixin/Windows/WeChatSetup.exe",
		"https://download.alicdn.com/wireless/qianniu/latest/qianniu_36217721281452.apk",
	}
	client := &http.Client{}
	headers := map[string]string{
		"User-Agent": "NSPlayer/12.00.22621.2428 WMFSDK/12.00.22621.2428",
	}

	for {
		for _, url := range urls {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				continue
			}

			for key, value := range headers {
				req.Header.Set(key, value)
			}

			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			buffer := make([]byte, 1024)
			for {
				n, err := resp.Body.Read(buffer)
				if n > 0 {
					lock.Lock()
					totalDownloadedBytes += int64(n)
					lock.Unlock()
				}
				if err != nil {
					if err == io.EOF {
						break
					}
					continue
				}
			}
		}
	}
}

// 打印下载流量信息
func printStats() {
	var downloadedGB float64 = 0
	for {
		time.Sleep(1 * time.Second)
		lock.Lock()
		downloadedGBOld := downloadedGB
		downloadedGB = float64(totalDownloadedBytes) / (1024 * 1024 * 1024) // 将字节转换为GB
		fmt.Printf("Downloaded: %.3f GB, %.2f MB/s\n", downloadedGB, (downloadedGB-downloadedGBOld)*1024)
		lock.Unlock()
	}
}

func main() {
	// 启动多个下载线程
	for i := 0; i < 20; i++ {
		go run()
	}

	// 启动一个线程每秒打印下载流量
	go printStats()

	// 防止主线程退出
	select {}
}
