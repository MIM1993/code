package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

//加密
func rsaEncryptdata(fiiename string, src []byte) ([]byte, error) {

	//调用函数，得到公钥
	pubKey, err := readRsaPubKey(fiiename)
	if err != nil {
		return nil, err
	}
	//加密
	info, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//解密
func rsaDecryptData(filename string, src []byte) ([]byte, error) {
	//读取私钥文件
	priKey, err := readRsaPriKey(filename)
	if err != nil {
		return nil, err
	}
	//解密
	info, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, src)
	if err != nil {
		return nil, err
	}

	return info, err
}

func main() {
	src := []byte("好好学习，天天向上")
	encode_src, err := rsaEncryptdata("rsaPublicKey.pem", src)
	if err != nil {
		fmt.Println("rsaEncryptdata err:", err)
		return
	}
	fmt.Printf("密文 ： %x\n", encode_src)

	plainText, err := rsaDecryptData("rsaPriKey.pem", encode_src)
	if err != nil {
		fmt.Println("解密过程报错:", err)
		return
	}

	fmt.Printf("明文：%s\n", plainText)

}
