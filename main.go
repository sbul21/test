package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goombaio/namegenerator"
)

func main() {
	f, err := os.OpenFile("import.csv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	for i := 0; i < 10; i++ {

		name := nameGenerator.Generate()

		n, err := w.WriteString(fmt.Sprintf("Klient_%s\n", name))
		_ = n
		if err != nil {
			log.Println(err)
		}
		w.Flush()
	}
}
