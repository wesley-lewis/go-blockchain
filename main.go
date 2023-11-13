package main

import "github.com/wesley-lewis/go-blockchain/network"

// server: container
// Transport => tcp, udp,
// Block
// Txn
// Keypairs

func main() {
	trLocal := network.NewLocalTransport("LOCAL")

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)

	s.Start()
}
