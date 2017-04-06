package csv

import (
	"encoding/csv"
	"encoding/json"
	jayna "github.com/lucasgomide/jayna/types"
	"io"
	"log"
)

func ToJSON(reader io.Reader, options *jayna.Options) ([]byte, error) {
	rows := make([]map[string]string, 0)
	r := csv.NewReader(reader)
	r.TrimLeadingSpace = true
	r.Comma = options.Delimiter

	for count := 0; ; count++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if options.Columns == nil {
			options.Columns = record
		}

		if count == 0 {
			continue
		}

		log.Print(options.Columns)

		row := make(map[string]string)
		for index, value := range options.Columns {
			if index > len(record)-1 {
				break
			}
			row[value] = record[index]
		}
		rows = append(rows, row)
	}

	data, err := json.Marshal(rows)
	if err != nil {
		return nil, err
	}
	return data, nil
}
