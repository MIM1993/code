package main

import (
	"fmt"
)

//
////定义区块结构
////实现基础字段
////补充字段
//type Block struct {
//	//前区块
//	PrevHash []byte
//	//哈希
//	Hash []byte
//	//数据
//	Data []byte
//}
//
////创建区块方法		区块       前hash值
//func NewBlock(data string, PrevHash []byte) *Block {
//	b := Block{
//		PrevHash: PrevHash,
//		Hash:     nil,
//		Data:     []byte(data),
//	}
//
//	//计算hash值
//	b.setHash()
//
//	return &b
//}
//
////提供计算区块hash值的方法
//func (b *Block) setHash() {
//	//data 是block各个字节流进行拼接
//	//二维切片
//	temp := [][]byte{
//		b.PrevHash,
//		b.Hash,
//		b.Data,
//	}
//	//Join函数
//	data := bytes.Join(temp, []byte{})
//
//	//比特币  sha256
//	hash := sha256.Sum256(data)
//
//	//赋值
//	b.Hash = hash[:]
//}
//
////定义区块连结构(使用数组模拟)
//type BlockChain struct {
//	Blocks []*Block //区块链
//}
//
//const genesisinfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
//
////提供一个创建区块链方法
//func NewBlockChain() *BlockChain {
//	genesisBlock := NewBlock(genesisinfo, nil)
//
//	bc := BlockChain{
//		Blocks: []*Block{genesisBlock},
//	}
//
//	return &bc
//
//}
//
////提供一个添加区块到链中的方法
//func (blockchain *BlockChain) AddBlockToChain(block *Block) {
//
//	blockchain.Blocks = append(blockchain.Blocks, block)
//}

//打印区块
func main() {
	bc := NewBlockChain()

	bc.AddBlockToChain("26号比特币暴涨20%")
	bc.AddBlockToChain("27号比特币暴涨30%")

	//遍历区块链
	for i, block := range bc.Blocks {
		fmt.Printf("\n==========当前区块高度%d============\n", i)
		fmt.Printf("PrevHash : %x\n", block.PrevHash)
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("MerkleRoot : %x\n", block.MerkleRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Bits : %d\n", block.Bits)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", string(block.Data))
	}
}

//字段
  //区块 block
  //目标值 tagget

//方法
  //run计算
  //功能：找到nonce，满足哈希值比目标值小
