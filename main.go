package main

import "github.com/shanedell/gorevproxy/cmd"

func main() {
	cli := cmd.New()

	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
