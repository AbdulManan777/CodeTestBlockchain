package assignment02

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce int
	BlockData        []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
	
}

type Blockchain struct {
	ChainHead *Block
}
var flag=true
func GenerateNonce(blockData []Transaction) int {
	
	//var onlyOnce sync.Once


	    if(flag==true){

			rand.Seed(time.Now().UnixNano())
		    flag=false
		}

		
		
	
	

	


	v := rand.Intn(100000)

	return v


}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
		blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {


	block:=new(Block)
	block.BlockData=blockData
    
	block.Nonce=GenerateNonce(block.BlockData)

	

	if chainHead == nil {
		block.PrevPointer = nil
		block.PrevHash = "000000000000000 (Nil)" //Genesis Block so storing 0000000 as prev hash
	}else 
	 {
       block.PrevPointer = chainHead
	   block.PrevHash = chainHead.CurrentHash
	}

	block.CurrentHash=CalculateHash(blockData,block.Nonce)
	
	return block

	
	



}

func ListBlocks(chainHead *Block) {


	i:=1
	
	for chainHead != nil {
		
		fmt.Println("\nBlock # ", i )
		fmt.Println("_____________Data of this Block_______________:   ")
		fmt.Println()
		DisplayTransactions(chainHead.BlockData)
		fmt.Println()
		fmt.Println("Previous Pointer: ", &chainHead.PrevPointer)
		fmt.Println("Current Block Nonce: ", chainHead.Nonce)
		fmt.Println("Previous Hash:    ", chainHead.PrevHash)
		fmt.Println("Current Hash: ", chainHead.CurrentHash)
		chainHead = chainHead.PrevPointer
		i++
	  }
	  fmt.Println()
	  
	 // DisplayTransactions(chainHead.BlockData)
	
}

func DisplayTransactions(blockData []Transaction) {

	for i:=0; i<len(blockData); i++{
        fmt.Println("Transaction-------------------",i,"-----------------------")
		
		fmt.Println("Sender: ", blockData[i].Sender)
		fmt.Println("Reciever: ", blockData[i].Receiver)
		fmt.Println("Amount: ", blockData[i].Amount)

	}
	
}

func NewTransaction(sender string, receiver string, amount int) Transaction {

	t1:=new(Transaction)

	t1.Sender=sender
	t1.Receiver=receiver
	t1.Amount=amount

	return *t1
	
	
}


