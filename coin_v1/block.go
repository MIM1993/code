package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//定义区块结构
//实现基础字段  前区块 hash 数据
//补充字段   version  时间戳 难度值
type Block struct {
	//version
	Version uint64

	//前区块
	PrevHash []byte

	//交易的根哈希值
	MerkleRoot []byte

	//时间戳
	TimeStamp uint64

	//难度值
	Bits uint64

	//随机数
	Nonce uint64

	//哈希
	Hash []byte

	//数据
	Data []byte
}

//创建区块方法		区块       前hash值
func NewBlock(data string, PrevHash []byte) *Block {
	b := Block{
		Version:    0,
		PrevHash:   PrevHash,
		MerkleRoot: nil,
		TimeStamp:  uint64(time.Now().Unix()),
		Bits:       0, //随意写的
		Nonce:      0, //随意写的
		Hash:       nil,
		Data:       []byte(data),
	}

	//计算hash值
	b.setHash()

	return &b
}

//提供计算区块hash值的方法
func (b *Block) setHash() {
	//data 是block各个字节流进行拼接
	//二维切片
	temp := [][]byte{
		uintToByte(b.Version),
		b.PrevHash,
		b.MerkleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(b.Nonce),
		b.Hash,
		b.Data,
	}
	//Join函数
	data := bytes.Join(temp, []byte{})

	//比特币  sha256
	hash := sha256.Sum256(data)

	//赋值
	b.Hash = hash[:]
}
