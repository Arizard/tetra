package tetra

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func readerFromData(s string, t Transform) *csv.Reader {
	reader := csv.NewReader(strings.NewReader(s))

	reader.Comma = t.Config.Comma
	reader.FieldsPerRecord = t.Config.FieldsPerRecord

	return reader
}

func noneOp(_ Transform, csvData string) (string, error) {
	reader := csv.NewReader(strings.NewReader(csvData))
	for {
		_, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}
	return csvData, nil
}

func sliceRowsOp(transform Transform, csvData string) (string, error) {

	var recordsOut [][]string

	var startIndex int
	var endIndex int

	reader := readerFromData(csvData, transform)

	recordsIn, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	tempStart := int(transform.KWArgs["start"].(float64))
	tempEnd := int(transform.KWArgs["end"].(float64))

	startIndex = tempStart
	if tempStart < 0 {
		startIndex = len(recordsIn) + tempStart
	}

	endIndex = tempEnd
	if tempEnd < 0 {
		endIndex = len(recordsIn) + tempEnd
	}

	for i := startIndex; i <= endIndex; i++ {
		record := recordsIn[i]

		recordsOut = append(recordsOut, record)
	}

	var b bytes.Buffer
	csvDataBuf := bufio.NewWriter(&b)
	csvWriter := csv.NewWriter(csvDataBuf)

	csvWriter.WriteAll(recordsOut)

	csvWriter.Flush()

	return b.String(), nil
}

func reverseRowsOp(transform Transform, csvData string) (string, error) {

	var recordsOut [][]string

	reader := readerFromData(csvData, transform)

	recordsIn, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	for i := len(recordsIn) - 1; i >= 0; i-- {
		record := recordsIn[i]

		recordsOut = append(recordsOut, record)
	}

	var b bytes.Buffer

	csvDataBuf := bufio.NewWriter(&b)

	csvWriter := csv.NewWriter(csvDataBuf)
	csvWriter.WriteAll(recordsOut)
	csvWriter.Flush()

	return b.String(), nil
}

func titleCaseColumnOp(transform Transform, csvData string) (string, error) {

	var recordsOut [][]string

	reader := readerFromData(csvData, transform)

	columnIndex := int(transform.KWArgs["index"].(float64))

	recordsIn, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	for _, record := range recordsIn {
		newRecord := record
		for j, cell := range newRecord {
			if columnIndex == j {
				newRecord[j] = strings.Title(strings.ToLower(cell))
			}
		}
		recordsOut = append(recordsOut, newRecord)
	}

	var b bytes.Buffer

	csvDataBuf := bufio.NewWriter(&b)

	csvWriter := csv.NewWriter(csvDataBuf)
	csvWriter.WriteAll(recordsOut)
	csvWriter.Flush()

	return b.String(), nil
}

var operationMap = map[string](func(Transform, string) (string, error)){
	"none":             noneOp,
	"slice_rows":       sliceRowsOp,
	"reverse_rows":     reverseRowsOp,
	"titlecase_column": titleCaseColumnOp,
}

func operate(transform Transform, csvData string) (string, error) {

	if operationMap[transform.Operation] == nil {
		return "", fmt.Errorf(
			"error: No such operation defined (%s)",
			transform.Operation,
		)
	}

	newCsvData, err := operationMap[transform.Operation](transform, csvData)

	return newCsvData, err
}
