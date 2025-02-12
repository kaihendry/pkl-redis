package main

import (
	"context"
	"fmt"

	"github.com/kaihendry/pkl-redis/gen/control"
)

func main() {
	cfg, err := control.LoadFromPath(context.Background(), "config.pkl")
	if err != nil {
		panic(err)
	}

	fmt.Printf("docker exec redis redis-cli -n 6 SET RL:toggle %s\n", cfg.Rl.Toggle)
	fmt.Printf("docker exec redis redis-cli -n 6 SET CB:toggle %s\n", cfg.Cb.Toggle)
}
