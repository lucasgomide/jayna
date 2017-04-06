package main

import (
	"flag"
	"fmt"
	"github.com/lucasgomide/jayna/shapes"
	jayna "github.com/lucasgomide/jayna/types"
	"log"
)

var (
	output    = flag.String("o", "", "Format to output, e.g: json")
	input     = flag.String("i", "", "Indicates input file, e.g: csv")
	filePath  = flag.String("f", "", "File of path to transform")
	delimiter = flag.String("d", "", "Indicates separator, e.g: ;")
)

func init() {
	flag.Parse()
}

func main() {
	var s jayna.Shape

	switch *input {
	case "csv":
		s = shapes.CSV{}
	default:
		log.Fatal("Input cannot be empty. Use flag -i to set it")
	}

	err := jayna.ValidateOutput(s, *output)

	if err != nil {
		log.Fatal(err)
	}

	result, err := convert(s)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}

func convert(s jayna.Shape) ([]byte, error) {
	return s.Convert(jayna.NewOptions(*output, *filePath, *delimiter, nil))
}
