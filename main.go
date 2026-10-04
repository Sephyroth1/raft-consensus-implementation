package main

import (
	"fmt"
	"sync"

	config "consensus.com/Config"
	rpc "consensus.com/RPC"
	node "consensus.com/RaftNode"

	shared "consensus.com/shared"
)

func main() {
	config := config.ConfigEnv{
		MinTimeout: 150,
		MaxTimeout: 300,
		Peers:      []shared.NodeEntry{},
		Majority:   3,
	}

	ports := [...]int{
		8000,
		8001,
		8002,
		8003,
		8004,
		8005,
		8006,
		8007,
	}

	for port := range ports {
		n := node.NewNode(shared.NodeEntry{
			0,
			"127.0.0.1",
			port,
			shared.Follower,
			0,
			0,
			[]shared.LogEntry{},
			0,
			0,
			[]int{},
			[]int{},
		}, config)

		config.Peers = append(config.Peers, *n)

		r := rpc.Raft{
			RaftNode: n,
		}

		rpc.Register(&r)

		for port := range ports {
			_, err := rpc.Listen(port)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	var wg sync.WaitGroup
	for _, nodeEntry := range config.Peers {
		n := node.ConvertNode(nodeEntry)
		wg.Add(1)
		go func(n *node.Node) {
			defer wg.Done()
			n.Run()
		}(&n)
	}
}
