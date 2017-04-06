package csv

import (
	"bytes"
	jayna "github.com/lucasgomide/jayna/types"
	"testing"
)

func TestShouldReturnsJSON(t *testing.T) {
	s := `a,b
				1,2
				3,4`

	reader := bytes.NewBufferString(s)
	json, err := ToJSON(reader, jayna.NewOptions("", "", "", nil))

	if err != nil {
		t.Error(err)
	}

	expected := `[{"a":"1","b":"2"},{"a":"3","b":"4"}]`

	if string(json) != expected {
		t.Errorf("Function ToJSON have been returned %s; but the expected %s", string(json), expected)
	}
}

func TestShouldReturnsJSONWithCustomColumns(t *testing.T) {
	s := `a,b
				1,2
				3,4`

	reader := bytes.NewBufferString(s)
	json, err := ToJSON(reader, jayna.NewOptions("", "", "", []string{"column1", "column 2"}))

	if err != nil {
		t.Error(err)
	}

	expected := `[{"column 2":"2","column1":"1"},{"column 2":"4","column1":"3"}]`

	if string(json) != expected {
		t.Errorf("Function ToJSON have been returned %s; but the expected %s", string(json), expected)
	}
}

func TestNumberOfCustomColumnGreatherThanCSVColumn(t *testing.T) {
	s := `a,b
				1,2
				3,4`

	reader := bytes.NewBufferString(s)
	json, err := ToJSON(reader, jayna.NewOptions("", "", "", []string{"column1", "column 2", "column3"}))

	if err != nil {
		t.Error(err)
	}

	expected := `[{"column 2":"2","column1":"1"},{"column 2":"4","column1":"3"}]`

	if string(json) != expected {
		t.Errorf("Function ToJSON have been returned %s; but the expected %s", string(json), expected)
	}
}
