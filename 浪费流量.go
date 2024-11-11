package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 全局变量，用于存储下载的总字节数
var totalDownloadedBytes int64 = 0

// 互斥锁，用于保护totalDownloadedBytes的并发访问
var lock sync.Mutex

// 下载文件并更新字节数
func run() {
	// 定义要下载的URL列表
	urls := []string{
		// "https://downloads.bdrive.com/netdrive/builds/0cfeb64848024a7ba4f28ce66549cbe2/NetDriveInstaller-3.17.730.dmg",
		// "https://op.yundasys.com/opserver/pages/downapp/image/img/bg-img.png",
		// "https://dldir1v6.qq.com/weixin/Windows/WeChatSetup.exe",
		"https://s.momocdn.com/s1/u/jbcjffabc/hertown/bg.png",
	}
	// 创建HTTP客户端
	client := &http.Client{Timeout: 2 * time.Second}
	// 定义请求头
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0",
	}

	// 无限循环，持续下载
	for {
		// 遍历每个URL
		for _, url := range urls {
			// 创建HTTP GET请求
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				continue // 如果请求创建失败，跳过当前URL
			}
			// 设置请求头
			for key, value := range headers {
				req.Header.Set(key, value)
			}
			// 执行请求
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("请求失败:", err)
				continue // 如果请求失败，跳过当前URL
			}
			// 确保响应体关闭
			func() {
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
						fmt.Println("读取出错:", err)
						break
					}
				}
			}()

		}
	}
}

// 打印下载流量信息
func printStats() {
	var downloadedGB float64 = 0
	for {
		time.Sleep(1 * time.Second) // 每秒打印一次
		lock.Lock()
		downloadedGBOld := downloadedGB
		// 将字节转换为GB
		downloadedGB = float64(totalDownloadedBytes) / (1024 * 1024 * 1024)
		// 打印下载的总量和速度
		fmt.Printf("Downloaded: %.3f GB, %.2f MB/s\n", downloadedGB, (downloadedGB-downloadedGBOld)*1024)
		lock.Unlock()
	}
}

func main() {
	// 手动输入线程数量
	var numThreads int
	fmt.Print("请输入线程数量: ")
	_, err := fmt.Scan(&numThreads)
	if err != nil {
		return // 如果输入错误，退出程序
	}

	// 启动多个下载线程
	for i := 0; i < numThreads; i++ {
		go run()
	}

	// 启动一个线程每秒打印下载流量
	go printStats()

	// 防止主线程退出
	select {}
}
