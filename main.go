package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func Menu() {
	var bc *Blockchain

	for {
		fmt.Println("Choose option:")
		fmt.Println("  1. Create new blockchain")
		fmt.Println("  2. Add block to the blockchain")
		fmt.Println("  3. Print all the blocks of the blockchain")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		switch choice {
		case 1:
			fmt.Println("You selected Option 1")
			bc = NewBlockchain()

		case 2:
			fmt.Println("You selected Option 2")
			if bc == nil {
				fmt.Println("Please create a new blockchain first")
			} else {
				var tnxs []*Transaction
				// make tnx
				//tnx = make([]*Transaction, 0)
				for {
					tnx := &Transaction{} // tnx is a pointer to a Transaction struct

					fmt.Printf("Input sender: ")
					fmt.Scan(&tnx.Sender)
					fmt.Printf("Input receiver: ")
					fmt.Scan(&tnx.Receiver)
					// Create a byte slice of the specified length
					randomBytes := make([]byte, 10)

					// Read random bytes from the crypto/rand.Reader into the byte slice
					_, err := rand.Read(randomBytes)
					if err != nil {
						fmt.Println("Error generating random bytes:", err)
						return
					}
					tnx.Signature = randomBytes
					// add tnx to transactions
					tnxs = append(tnxs, tnx)
					fmt.Println("Do you want to add more transaction? (y/n)")
					var choice string
					fmt.Scan(&choice)
					if choice == "n" || choice == "N" {
						break
					}
				}
				bc.AddBlock(tnxs)
			}
		case 3:
			fmt.Println("You selected Option 3")
			if bc == nil {
				fmt.Println("Please create a new blockchain first")
			} else {
				bc.printChain()
			}

		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

func main() {
	Menu()
}

// func TestBlockChain(t *testing.T){
// 	bc := NewBlockchain()
// 	bc.AddBlock("Send 1 BTC to Ivan")
// 	bc.AddBlock("Send 2 more BTC to Ivan")

// 	for _, block := range bc.blocks {

// 	}
// }
