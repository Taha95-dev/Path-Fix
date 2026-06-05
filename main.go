package main

import (
    "fmt"
    "os"

	"pathfix/pkg/search"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: prefix <command>")
		return
	}

	cmd := os.Args[1]

	if search.InPATH(cmd){
		fmt.Printf("✅'%s' is already in your PATH\n", cmd)
		return
	}
	found := search.FindInPaths(cmd)
	if found != "" {
		fmt.Printf("Found '%s' at %s\n", cmd, found)
		return
	} else {
		fmt.Printf("Could not find '%s' on your system\n", cmd)
	}
}
