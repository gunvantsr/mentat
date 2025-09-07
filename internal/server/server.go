package server

import (
	"fmt"
	"net"

	"github.com/gunvantsr/mentat/config"
)

type Server struct {
	Config *config.Config
}

func New(cfg *config.Config) (*Server, error) {

	fmt.Println("Server initialized with config:", cfg)

	startNode(cfg.Nodes)

	return (&Server{
		Config: cfg,
	}), nil
}

func startNode(nodes []config.Node) {
	for _, node := range nodes {
		fmt.Printf("Starting node: %+v\n", node)
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", node.Host, node.Port))
		if err != nil {
			fmt.Printf("Error starting node %s:%d - %v\n", node.Host, node.Port, err)
			continue
		}
		go handleConnection(listener)
	}
}

func handleConnection(listener net.Listener) {
	defer listener.Close()
}
