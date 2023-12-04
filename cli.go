package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - create a blockchain and send genesis block reward to ADDRESS")
	
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}


func (cli *CLI) Run() {
	cli.validateArgs()
	
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	
	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(	1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}
}

func (cli *CLI) createBlockchain(address string) {
	//code

}
