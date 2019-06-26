package main

import "math/big"

//实现挖矿功能

//字段
//区块 block
//目标值 tagget

//方法
//run计算
//功能：找到nonce，满足哈希值比目标值小

type ProofofWork struct {
	//区块
	block *Block
	//目标值 与计算出的哈希值作对比
	target *big.Int
}

//创建一个工作量证明
//block  用户提供
//target 系统提供
func NewProofofWork(block *Block) *ProofofWork {
	pow := ProofofWork{
		block: block,
	}

	//难度值  写死 后面补充
	targetStr := "0000100000000000000000000000000000000000000"
	tmpBigInt := new(big.Int)
	tmpBigInt.SetString(targetStr, 16)

	pow.target = tmpBigInt

	return &pow
}
