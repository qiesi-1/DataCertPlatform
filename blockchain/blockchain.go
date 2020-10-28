package blockchain

import (
	"github.com/bolt"
)

const BLOCKCHAIN  = "blockchain.db"
const BLOCK_NAME  = "blocks"
const LAST_HASH = "lasthash"

//区块链结构体的定义，代表的是一条区块链
type BlockChain struct {
	LastHash []byte //表示区块链中最新区块的hash，用于查找最新区块内容
	BoltDb  *bolt.DB//区块链中操作区块数据文件的数据库操作对象
}

//创建一条区块链

func NewBlockChian() BlockChain {
	var bc BlockChain
	//1.先打开文件
	db,err :=bolt.Open(BLOCKCHAIN,0600,nil)
	//2.查看chain。db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BLOCK_NAME))//假设有桶
		if bucket == nil{//没有桶，新建
			bucket,err = tx.CreateBucket([]byte(BLOCK_NAME))
			if err !=nil {
				panic(err.Error())
			}
		}
		lastHash :=bucket.Get([]byte(LAST_HASH))
		if len(lastHash) ==0 {//桶中没有lasthash记录，徐常见创世区块，保存
			//创世区块
			genesis := CreateGenesisBlock()
			genesisBytes := genesis.Serialize()
			//创世区块保存到boltdb中
			bucket.Put(genesis.Hash,genesisBytes)
			//更新指向最新区块的lasthash值
			bucket.Put([]byte(LAST_HASH),genesis.Hash)
			bc = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
		}else {
			lastHash1 := bucket.Get([]byte(LAST_HASH))
			bc = BlockChain{
				LastHash: lastHash1,
				BoltDb:   db,
			}
		}

		return nil
	})
	//if err!=nil {
	//	panic(err.Error())
	//}
	//
	//
	////把创世区块保存到数据库文件中
	//db.Update(func(tx *bolt.Tx) error {
	//	bucket,err :=tx.CreateBucket([]byte(BLOCK_NAME))
	//	if err!=nil{
	//		panic(err.Error())
	//	}
	//	fmt.Println(bucket)
	//	/*
	//	//序列化
	//	genesisBytes := genesis.Serialize()
	//	//把创世区块储存到桶中
	//	bucket.Put(genesis.Hash,genesisBytes)
	//	//更新最新区块的hash值记录
	//	bucket.Put([]byte(LAST_HASH),genesis.Hash)
	//	*/
	//
	//	return nil
	//})
	return bc
}

//保存数据到区块链中；先生成一个新区块，然后将新区块添加到区块链中
func(bc BlockChain ) SaveData(data []byte) {
	//1.从文件读取到最新区块
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_NAME))
		if bucket != nil {
			panic("读取区块链数据失败")
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		//lastBlockBytes := bucket.Get(lastHash) 后面已经赋值，代码变更
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock,_ =DeSerialize(lastBlockBytes)
		return nil
	})
	//新建一个区块
	newblock := NewBlock(lastBlock.Height+1,lastBlock.Hash,data)
	//将新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BLOCK_NAME))
		//把新创建的区块存入到boltdb中
		bucket.Put(newblock.Hash,newblock.Serialize())
		//更新lasthash对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH),newblock.Hash)
		bc.LastHash =newblock.Hash//将区块链实例的lasthash值更新为最新区块的hash

		return nil
	})



}