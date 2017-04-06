package shapes

import (
	"github.com/lucasgomide/jayna/shapes/csv"
	jayna "github.com/lucasgomide/jayna/types"
	"os"
)

type CSV struct{}

func (c CSV) Convert(options *jayna.Options) ([]byte, error) {
	f, err := os.Open(options.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var output []byte

	switch options.Output {
	case "json":
		output, err = csv.ToJSON(f, options)
	}

	return output, err
}

func (c CSV) ListOutputAvailable() []string {
	return []string{"json"}
}
