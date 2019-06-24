package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"crypto/x509"
	"crypto/tls"
)

func main() {
	/*需要告诉client，我们支持的证书*/
	//读取证书
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	//创建pool
	capool := x509.NewCertPool()

	//将我们认可的根证书，添加pool
	ok := capool.AppendCertsFromPEM(caCert)
	if !ok {
		fmt.Println("添加CA池失败！")
		return
	}

	//创建tls结构   ssl3.0=tls1.0
	config := tls.Config{
		RootCAs: capool,
	}

	//创建http客户端句柄
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &config,
		},
	}

	fmt.Println("客户端开始访问...")

	//发起GET请求
	response, err := client.Get("https://www.mim.com:8082")
	if err != nil {
		fmt.Println("Get err :", err)
		return
	}

	//获取body数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll err :", err)
		return
	}

	defer response.Body.Close()

	fmt.Printf("body : \n%s\n", body)
}
