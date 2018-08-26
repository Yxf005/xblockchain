package core

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp int64        // 创建时间
	Data []byte            // 区块内容
	PreBlockHash []byte    // 前一个区块的hash 值
	Hash []byte           // 此区块的 hash 值
}

/**
创建新的区块
 */
func NewBlock(data string, preBlockHash []byte)  *Block {

	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
	block.SetHash()
	return block
}

/**
设置 hash 值
Block 的接口方法
 */
func (b *Block) SetHash()  {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

/**
创建 创世 区块（第一个区块）
 */
func NewGenesisBlock() *Block {
	return NewBlock("this is genesis block", []byte{})

}