package blockchain


type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash: []byte{},
		Data: []byte(data),
		PrevHash: prevHash,
		Nonce: 0,
	}

	pow := NewProof(block)
	nonce, hash := pow.Run()
	
	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks) - 1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}