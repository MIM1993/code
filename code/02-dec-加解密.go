package main

import (
	"crypto/des"
	"crypto/cipher"
	"bytes"
	"fmt"
)

//输入明文，得到密文
func desCBCEncrypt(plaintext, key []byte) ([]byte, error) {
	//第一步：创建des密码接口, 输入秘钥，返回接口
	// func NewCipher(key []byte) (cipher.Block, error)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := block.BlockSize()

	//第二步：创建cbc分组
	// 返回一个密码分组链接模式的、底层用b解密的BlockMode接口
	// func NewCBCEncrypter(b Block, iv []byte) BlockMode
	//创建一个8字节的初始化向量
	iv := bytes.Repeat([]byte("1"), size)
	mode := cipher.NewCBCEncrypter(block, iv)

	//第三步：填充
	plaintext, err = paddingNumber(plaintext, size)
	if err != nil {
		fmt.Println("paddingNumber err:", err)
		return nil, err
	}

	//第四步：加密
	//密文与明文共享空间，没有额外分配
	mode.CryptBlocks(plaintext /*密文*/ , plaintext /*明文*/)

	return plaintext, nil
}

//输入密文，得到明文
func desCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	//第一步：创建des密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := block.BlockSize()

	//第二步：创建cbc分组
	iv := bytes.Repeat([]byte("1"), size)
	mode := cipher.NewCBCDecrypter(block, iv)

	//第三步：解密
	mode.CryptBlocks(encryptData, encryptData)

	//第四步: 去除填充
	encryptData, err = unpaddingNumber(encryptData)
	if err != nil {
		fmt.Println("unpaddingNumber err:", err)
		return nil, err
	}

	return encryptData, nil
}

//填充数据
func paddingNumber(src []byte, blockSize int) ([]byte, error) {
	//获取分组后剩余长度
	size := len(src)
	leftNumber := size % blockSize

	//填充个数
	neednumber := blockSize - leftNumber
	fmt.Println("填充个数", neednumber)

	num := byte(neednumber)
	//num := string(neednumber)
	//num1 := strconv.Itoa(neednumber)
	//fmt.Printf("string(needNumber) : %x\n",num)
	//fmt.Printf("itoa : %x\n",num1)
	//fmt.Println("num", num)

	//创建一个切片
	newslice := bytes.Repeat([]byte{num}, neednumber)

	//将新切片追加
	src = append(src, newslice...)

	fmt.Println("paddingNumber")
	return src, nil
}

//解密后去除填充
func unpaddingNumber(src []byte) ([]byte, error) {
	fmt.Println("unpaddingNumber")

	size := len(src)

	//获取最后的byte数字
	lastchar := src[size-1]

	//转为int
	num := int(lastchar)

	//截取切片
	src = src[:size-num]

	return src, nil
}

func main() {
	src := "床前明月光，疑是地上霜"
	key := "12345678"

	//加密处理
	encryptData, err := desCBCEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("加密后的数据： %x\n", encryptData)

	//调用解密函数
	plainText, err := desCBCDecrypt(encryptData, []byte(key))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("解密后的数据: %s\n", plainText)

}
