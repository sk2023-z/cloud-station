package main

import (
	"fmt"

	"github.com/sk2023-z/cloud-station/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
