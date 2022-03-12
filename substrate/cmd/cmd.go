package main

import (
	"fmt"

	_ "github.com/aws/aws-sdk-go/aws/awserr"
)

func main() {
	Run("main")
}

func Run(v string) {
	fmt.Println("running import-go-modules in ", v)
}
