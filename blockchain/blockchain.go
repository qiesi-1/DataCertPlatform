package blockchain

import (
	"errors"
	"github.com/bolt"
	"math/big"
)

var CHAIN  *BlockChain
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
	var bc *BlockChain
	//1.先打开文件
	db,err :=bolt.Open(BLOCKCHAIN,0600,nil)
	//2.查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BLOCK_NAME))//假设有桶
		if bucket == nil{//没有桶，新建
			bucket,err = tx.CreateBucket([]byte(BLOCK_NAME))
			if err !=nil {
				panic(err.Error())
			}
		}
		lastHash :=bucket.Get([]byte(LAST_HASH))
		if len(lastHash) ==0 {//桶中没有lasthash记录，需创建创世区块，保存
			//创世区块
			genesis := CreateGenesisBlock()
			genesisBytes := genesis.Serialize()
			//创世区块保存到boltdb中
			bucket.Put(genesis.Hash,genesisBytes)
			//更新指向最新区块的lasthash值
			bucket.Put([]byte(LAST_HASH),genesis.Hash)
			bc = &BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
		}else {
			lastHash1 := bucket.Get([]byte(LAST_HASH))
			bc = &BlockChain{
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
	CHAIN = bc
	return *bc
}


//该方法用于遍历区块链，并将所有区块查出，返回
func (bc BlockChain) QueryAllBlocks()([]*Block,error){

	blocks := make([]*Block,0)//切片容器，用于盛放查询到的区块
	db := bc.BoltDb
	var err error
	//从chain文件查询所有区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_NAME))
		if bucket == nil{
			err = errors.New("chaxunshibaibianli!")
			return err
		}
		//bucket存在
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)//默认值零的大整数
		for{

			//根据区块hash值获取对应区块
			eachBlockBytes :=bucket.Get(eachHash)
			//反序列化操作
			eachBlock,_:= DeSerialize(eachBlockBytes)
			//将遍历到的每一个区块放到容器里
			blocks = append(blocks,eachBlock)
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig)==0 {//找到创世区块
				break //跳出循环
			}
			//不满足条件，未找到创世区块
			eachHash = eachBlock.PrevHash
		}
		return  nil
	})
	return  blocks,err
}


//该方法用于完成根据用户输入的区块高度查询对应信息
func (bc BlockChain) QueryBlockByHeigt(height int64)(*Block,error){
	if height <0 {
		return nil,nil
	}
	db := bc.BoltDb

	var errs error
	var eachBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_NAME))
		if bucket ==nil {
			errs = errors.New("读取区块信息失败")
			return errs
		}
		//定义一个外部变量
		eachHash := bc.LastHash
		for {
			//获取到最后一个区块的hash
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
			eachBlock,errs := DeSerialize(eachBlockBytes)
			if errs != nil {
				//fmt.Println("遍历区块数据：",eachBlockBytes)
				return errs
			}
			if eachBlock.Height < height {
				break
			}
			if eachBlock.Height== height {
				break
			}
			//如果高度不满足条件
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return 	eachBlock,errs

}




//保存数据到区块链中；先生成一个新区块，然后将新区块添加到区块链中
func(bc BlockChain ) SaveData(data []byte)(Block,error) {
	//1.从文件读取到最新区块
	db := bc.BoltDb
	var lastBlock *Block
	//error的自定义
	var  err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_NAME))
		if bucket == nil {
			err = errors.New("读取区块链数据失败")
			return err
			//panic("读取区块链数据失败")
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		//lastBlockBytes := bucket.Get(lastHash) 后面已经赋值，代码变更
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock,_ =DeSerialize(lastBlockBytes)
		return nil
	})
	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1,lastBlock.Hash,data)

	//将新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BLOCK_NAME))
		blockBytes :=newBlock.Serialize()
		//把新创建的区块存入到boltdb中
		bucket.Put(newBlock.Hash,blockBytes)
		//更新lasthash对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		bc.LastHash = newBlock.Hash//将区块链实例的lasthash值更新为最新区块的hash
		return nil
	})
	//返回值语句，newblock，err可能包含错误信息，不可为nil
	return newBlock,err
}

//该方法用于根据用户输入的认证号查询到对应的区块信息
func (bc BlockChain)QueryBloockByCertId(cert_id string) (*Block,error) {
	db := bc.BoltDb
	var err error
	var block *Block
	db.View(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BLOCK_NAME))
		if bucket == nil {//判断桶是否存在
			err = errors.New("查询链上数据发生错误，sorry")
			return err
		}
		eachHash :=bc.LastHash
		eachBig := new(big.Int)
		zeroBig := new(big.Int)
		for  {
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock,err :=DeSerialize(eachBlockBytes)
			if err != nil{
				break
			}
			//将遍历到的区块数据跟用户提供的认证号进行比较
			if string(eachBlock.Data) == cert_id{//if成立找到区块
				block = eachBlock
				break
			}
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig==e{

			}

			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return block,err
}
