package main

import (
	"context"
	"fmt"
	"os"

	"spotify_to_tidal/shell/platform"
)

func main() {
	shell := platform.New()
	fmt.Printf("spotify_to_tidal shell (%s)\n", shell.OS())
	if err := shell.Run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
