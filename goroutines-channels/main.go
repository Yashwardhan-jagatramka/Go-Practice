package main

import (
	_ "github.com/Yashwardhan-jagatramka/matchers"
	"github.com/Yashwardhan-jagatramka/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
