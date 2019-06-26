package main

//定义区块连结构(使用数组模拟)
type BlockChain struct {
	Blocks []*Block //区块链
}

//常量
const genesisinfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

//提供一个创建区块链方法
func NewBlockChain() *BlockChain {
	genesisBlock := NewBlock(genesisinfo, nil)

	bc := BlockChain{
		Blocks: []*Block{genesisBlock},
	}

	return &bc
}

//提供一个添加区块到链中的方法
//参数  只需要数据  不需要Pervhash值
func (bc *BlockChain) AddBlockToChain(data string) {

	//得到最后一个区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	//前hash值
	prevHash := lastBlock.Hash

	//创建block
	newBlock := NewBlock(data, prevHash)

	//添加到链中
	bc.Blocks = append(bc.Blocks, newBlock)
}
