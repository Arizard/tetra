package tetra

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func noneOperation(transform Transform, csvData string) (string, error) {
	reader := csv.NewReader(strings.NewReader(csvData))
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		fmt.Printf("%s", record)
	}
	return csvData, nil
}

var operationMap = map[string](func(Transform, string) (string, error)){
	"none": noneOperation,
}

func operate(transform Transform, csvData string) (string, error) {

	if operationMap[transform.Operation] == nil {
		return "", fmt.Errorf(
			"error: No such operation defined (%s)",
			transform.Operation,
		)
	}

	newCsvData := operationMap[transform.Operation](transform, csvData)

	return newCsvData, nil
}
