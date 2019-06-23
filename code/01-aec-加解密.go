package main

/*
	AEC加密
*/

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"bytes"
)

//加密
func aesCtrEncrypt(plainText, key []byte) ([]byte, error) {
	//aes 包  golang内置的
	//1、创建一个接口  参数：秘钥； 返回值一个分组接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//打印aes分组长度
	fmt.Println("block.BlockSize:", block.BlockSize())

	//2、创建分组模式ctr
	//crypto/cipher 包
	//iv要与算法长度一至 16 bit
	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	stream := cipher.NewCTR(block, iv)

	//3、加密  参数 ： 密文空间  明文
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)


	return dst, nil

}

//解密
func aesCtrDecrypt(encryptText, key []byte) ([]byte, error) {

	//TODO

	return []byte("hello world"), nil
}

func main() {
	//明文
	src := "你好"
	//秘钥  aes 16 字节
	key := "1234567887654321"
	//加密
	encryptData, err := aesCtrEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("加密错误,err:", err)
		return
	}

	//fmt.Printf("加密后的数据： %x\n", encryptData)
	fmt.Printf("encryptData: %x\n", encryptData)
	//解密
	//plainText, err := aesCtrDecrypt(encryptData, []byte(key))
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//
	//fmt.Printf("解密后数据：%s", plainText)

}
