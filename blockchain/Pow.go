package blockchain

import (
	"bytes"
	"data/tools"
	"math/big"
)



const DIFFICULTY = 10
//工作量证明算法结构体
type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block  Block //要找的nonce值对应区块

}


//实例化一个pow、算法实例
func NewPoW(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t,255-DIFFICULTY)
	pow := ProofOfWork{
		Target: t,
		Block:  Block{},
	}
	return pow
}

//run方法用于寻找合适的nonce值
func (p ProofOfWork) Run() ([]byte,int64) {
	var nonce int64
	nonce = 0

	var blockHash []byte
	for {//使用无限循环
		block := p.Block
		heightBytes, _ := tools.Int64ToByte(block.Height)
		timeBytes, _ := tools.Int64ToByte(block.TimeStamp)
		versionBytes := tools.StringToBytes(block.Version)

		nonceBytes, _ := tools.Int64ToByte(nonce)
		var blockBytes []byte
		//已有区块信息和尝试nonce值得拼接信息
		blockBytes = bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.PrevHash,
			block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})
		//区块和尝试的nonce值拼接后得到的hash值
		blockHash = tools.SHA256HashBlock(blockBytes)

		//目标值
		target := p.Target
		var hashBig *big.Int   //声明和定义
		hashBig = new(big.Int) //为变量声明地址
		hashBig = hashBig.SetBytes(blockHash)

		if hashBig.Cmp(target) == -1 {
			//满足条件，停止寻找
			break
		}
		nonce++
		//fmt.Println("nonce值",nonce)
	}
	//将找到的符合规则的nonce返回
	return blockHash,nonce
}


