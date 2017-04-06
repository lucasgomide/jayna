package shapes

import (
	jayna "github.com/lucasgomide/jayna/types"
	"reflect"
	"testing"
)

func TestFileDoesNotExist(t *testing.T) {
	options := jayna.NewOptions("json", "not_exist.csv", ",", nil)

	csv := CSV{}
	_, err := csv.Convert(options)

	if err == nil {
		t.Error("The file has been found")
	}
}

func TestConvertToJSON(t *testing.T) {
	options := jayna.NewOptions("json", "../testdata/csv_to_json.csv", ",", nil)

	csv := CSV{}
	json, err := csv.Convert(options)

	if err != nil {
		t.Error(err)
	}

	expected := `[{"a":"1","b":"2"},{"a":"3","b":"4"},{"a":"5","b":"6"}]`
	if string(json) != expected {
		t.Errorf("Method Convert have been returned %s; but the expected %s", string(json), expected)
	}
}

func TestListOutputAvailable(t *testing.T) {
	csv := CSV{}
	list := csv.ListOutputAvailable()
	if !reflect.DeepEqual(list, []string{"json"}) {
		t.Error("List of available output has no match")
	}
}
