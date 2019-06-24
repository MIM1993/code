package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

//签名过程

//读取私钥
//数字签名
//				公钥             被签名数据
func rsaSignData(filename string, src []byte) ([]byte, error) {
	prikey, err := readRsaPriKey(filename)
	if err != nil {
		return nil, err
	}

	//创建hash值

	hash := sha256.Sum256(src)

	//签名
	sigData, err := rsa.SignPKCS1v15(rand.Reader, prikey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}

	return sigData, nil
}

//认证\
//				   私钥             校验数据
func rsaVweifyData(filename string, src []byte, sigData []byte) bool {
	//读取公钥
	pubKey, err := readRsaPubKey(filename)
	if err != nil {
		return false
	}

	hash := sha256.Sum256(src)

	//签名认证
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], sigData)
	if err != nil {
		return false
	}

	return true
}

func main() {
	src := []byte("好好学习")
	sig, err := rsaSignData("rsaPriKey.pem", src)
	if err != nil {
		fmt.Println("签名错误")
		return
	}
	fmt.Printf("%x\n", sig)
	//src = []byte("好好学习1")
	result := rsaVweifyData("rsaPublicKey.pem", src, sig)
	fmt.Println("验证结果", result)

}
