package main

import (
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*生成私钥*/

func generateRsakeypair(bits int) error {
	fmt.Println("创建私钥")
	//创建私钥
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("GenerateKey err:", err)
		return err
	}

	//对私钥编码，生成der格式字符串
	derText, err := x509.MarshalPKCS8PrivateKey(priKey)
	if err != nil {
		fmt.Println("MarshalPKCS8PrivateKey err:", err)
		return err
	}

	//对der字符串拼装到pem格式数据块中
	//填充block（证书）
	block := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   derText,
	}

	//创建文件保存私钥
	f, err := os.Create("rsaPriKey.pem")
	if err != nil {
		fmt.Println("Create rsaPriKey err:", err)
		return err
	}

	defer f.Close()

	//将pem进行base64编码，得到私钥
	err = pem.Encode(f, &block)
	if err != nil {
		fmt.Println("Encode err:", err)
		return err
	}

	fmt.Println("===========创建私钥完成========")

	fmt.Println("创建公钥")
	//通过私钥，生成公钥
	pubKey := priKey.PublicKey

	//对公钥进行编码，生成der格式字符串
	//使用指针
	derText, err = x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		fmt.Println("MarshalPKIXPublicKey err :", err)
		return err
	}

	//将der字符串拼装到pem格式数据块中
	block = pem.Block{
		Type:    "RSA PUBILC KEY",
		Headers: nil,
		Bytes:   derText,
	}

	//将pem格式数据块进行bases64编码  得到最终公钥
	f1, err := os.Create("rsaPublicKey.pem")
	if err != nil {
		fmt.Println("Create rsaPublicKey err ：", err)
		return err
	}
	defer f1.Close()

	pem.Encode(f1, &block)

	fmt.Println("===========创建公钥完成========")

	return nil
}

func main() {
	bits := 1024
	err := generateRsakeypair(bits)
	if err != nil {
		fmt.Println("私钥生成错误 err:", err)
		return
	}
}
