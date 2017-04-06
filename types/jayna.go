package jayna

import "errors"

type Shape interface {
	Convert(options *Options) ([]byte, error)
	ListOutputAvailable() []string
}

type Options struct {
	Input     string   // -i csv
	FilePath  string   // -f ~/file_to/path/barzu.csv
	Output    string   // -o json
	Data      string   // "a;b\n1;2;3.4"
	Delimiter rune     // ;
	Columns   []string //
}

func NewOptions(output string, filePath string, delimiter string, columns []string) *Options {

	if delimiter == "" {
		delimiter = ","
	}

	return &Options{
		Output:    output,
		FilePath:  filePath,
		Delimiter: []rune(delimiter)[0],
		Columns:   columns,
	}
}

func ValidateOutput(s Shape, output string) error {
	for _, b := range s.ListOutputAvailable() {
		if b == output {
			return nil
		}
	}

	return errors.New("Output " + output + " not supported yet")
}
