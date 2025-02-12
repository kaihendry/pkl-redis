package main

import (
	"context"
	"fmt"

	"github.com/kaihendry/pkl-redis/gen/control"
)

func main() {
	cfg, err := control.LoadFromPath(context.Background(), "rl.pkl")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running: %v\n", cfg.Rl.Toggle)
}
