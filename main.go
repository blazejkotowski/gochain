package main

import (
	"fmt"
	"time"

	"github.com/blazejkotowski/gochain/models"
)

func main() {
	bc := models.NewBlockchain()
	time.Sleep(2 * time.Second)

	transactions := []*models.Transaction{
		models.NewTransaction("1st block, 1st transaction"),
		models.NewTransaction("1st block, 2nd transaction"),
	}
	block := models.NewBlock(bc.GetLastBlock().Hash(), transactions)
	bc.AddBlock(block)
	time.Sleep(2 * time.Second)

	transactions = []*models.Transaction{models.NewTransaction("2nd block, 1st transaction")}
	block = models.NewBlock(bc.GetLastBlock().Hash(), transactions)
	bc.AddBlock(block)

	fmt.Println(bc.String())
}
