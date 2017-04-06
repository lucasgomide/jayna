package shapes

import (
	"encoding/json"
	"errors"
	jayna "github.com/lucasgomide/jayna/types"
	"strings"
)

type JSON struct{}

func (j JSON) Transform(content interface{}, options jayna.Options) ([]byte, error) {
	if ok, err := j.inputIsSupported(options.Input); !ok {
		return nil, err
	}

	var (
		output []byte
		err    error
	)

	switch options.Input {
	case "csv":
		output, err = transformFromCSV(content)
	}

	return output, err
}

func transformFromCSV(content interface{}) ([]byte, error) {
	lines := strings.Split(content.(string), "\n")

	headers := strings.Split(lines[0], ";")
	body := lines[1:len(lines)]

	var output []map[string]string
	for i := 0; i < len(lines)-1; i++ {
		line := strings.Split(body[i], ";")
		values := make(map[string]string)
		for j := 0; j < len(headers); j++ {
			values[headers[j]] = string(line[j])
		}
		output = append(output, values)
	}

	b, err := json.Marshal(output)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (j JSON) listInputSupported() []string {
	return []string{"csv"}
}

func (j JSON) inputIsSupported(input string) (bool, error) {
	for _, b := range j.listInputSupported() {
		if b == input {
			return true, nil
		}
	}
	return false, errors.New("Input " + input + " not supported")
}
