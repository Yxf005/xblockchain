package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"utils"
)

var maxNonce = math.MaxInt64

const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow

}

func (pow *ProofOfWork) createData(nonce int) []byte {

	data := bytes.Join([][]byte{
		pow.block.PreBlockHash,
		pow.block.Data,
		utils.Int2Hex(pow.block.Timestamp),
		utils.Int2Hex(int64(targetBits)),
		utils.Int2Hex(int64(nonce)),
	},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Execute() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing  %s \n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.createData(nonce)

		hash = sha256.Sum256(data)

		fmt.Printf("\r %x", hash)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break

		} else {
			nonce++
		}

	}
	fmt.Println()
	return nonce, hash[:]

}

func (pow *ProofOfWork) Validata() bool {
	var hashInt big.Int

	data := pow.createData(pow.block.Nonce)

	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1

}
