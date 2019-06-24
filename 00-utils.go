package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func readRsaPubKey(filename string) (*rsa.PublicKey, error) {
	//读取公钥文件
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//解码
	block, _ := pem.Decode(info)

	//得到der
	der := block.Bytes

	//得到公钥
	pubInterface, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}

	//断言
	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, nil
	}

	return pubKey, nil
}

func readRsaPriKey(filename string) (*rsa.PrivateKey, error) {
	//读取私钥文件
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	//解码
	block, _ := pem.Decode(info)

	//得到der
	der := block.Bytes

	//得到公钥
	priKeyInter,err :=x509.ParsePKCS8PrivateKey(der)
	if err!=nil{
		return nil,err
	}

	//断言
	priKey,ok :=priKeyInter.(*rsa.PrivateKey)
	if !ok{
		return nil,errors.New("priKey no ok!")
	}

	return priKey, nil
}
