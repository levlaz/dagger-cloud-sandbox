package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	cv := client.CacheVolume("hubvisor_test_0")
	output, err := client.
		Container().
		From("node:latest").
		WithMountedCache("/opt/node_modules", cv).
		WithWorkdir("/opt").
		WithExec([]string{"npm", "install", "lodash"}).
		Stdout(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
