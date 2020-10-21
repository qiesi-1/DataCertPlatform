package blockchain

import "math/big"

//工作量证明算法结构体
type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block  Block //要找的弄侧值对应区块

}

//实例化一个pow、算法实例
func NewPoW(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t,255)
	pow := ProofOfWork{
		Target: t,
		Block:  Block{},
	}
	return pow
}
