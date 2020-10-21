package blockchain

import (
	"bytes"
	"data/tools"
	"time"
)

//定义区块结果体
type Block struct {
	Height 		int64 //区块高度，第几个区块
	TimeStamp 	int64//区块产生的时间戳
	PrevHash 	[]byte//上一个区块的hash值
	Data 		[]byte//数据字段
	Hash 		[]byte//当前区块hash值
	Version 	string//版本号
	Nonce 		int64 //区块对应的nonce值
}

//创建一个新区块
func NewBlock(height int64,prevHash []byte,data []byte) (Block) {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		Version:   "0X01",
	}

	//1.将block结构体数据转换成【】byte类型
	heightBytes,_ := tools.Int64ToByte(block.Height)
	timeStampBytes,_ := tools.Int64ToByte(block.TimeStamp)
	versionBytes := tools.StringToBytes(block.Version)

	var blockBytes []byte
	//bytes.Join 拼接
	bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	},[]byte{})

	// 调用hash计算，对区块进行sha256计算
	block.Hash = tools.SHA256HashBlock(blockBytes)
	//挖矿竞争，获得记账权
	return block
}

//创建创世区块
func CreateGenesisBlock()Block  {
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
    return genesisBlock
}