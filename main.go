package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/kaihendry/pkl-redis/gen/control"
)

func main() {
	flush := flag.Bool("flush", false, "flush redis")
	flag.Parse()
	if *flush {
		fmt.Println("docker exec redis redis-cli -n 6 FLUSHDB")
	}
	redisCmds, err := genRedisBinds(flag.Args()[0])
	if err != nil {
		panic(err)
	}

	for _, cmd := range redisCmds {
		fmt.Println(cmd)
	}
}

func genRedisBinds(pkgFileName string) (rediscmds []string, err error) {
	rlCfg, err := control.LoadFromPath(context.Background(), pkgFileName)
	if err != nil {
		return nil, err
	}
	if rlCfg.Rl != nil && rlCfg.Rl.Toggle.String() != "" {
		rediscmds = append(rediscmds, fmt.Sprintf("docker exec redis redis-cli -n 6 SET RL:toggle %s", rlCfg.Rl.Toggle))
	}
	if rlCfg.Cb != nil && rlCfg.Cb.Toggle.String() != "" {
		rediscmds = append(rediscmds, fmt.Sprintf("docker exec redis redis-cli -n 6 SET CB:toggle %s", rlCfg.Cb.Toggle))
	}
	return rediscmds, nil
}
