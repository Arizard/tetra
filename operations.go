package tetra

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func noneOp(transform Transform, csvData string) (string, error) {
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

	reader := csv.NewReader(strings.NewReader(csvData))

	reader.Comma = transform.Config.Comma
	reader.FieldsPerRecord = transform.Config.FieldsPerRecord

	recordsIn, err := reader.ReadAll()

	if tempStart, err := strconv.Atoi(transform.KWArgs["start"].(string)); err != nil {
		return "", fmt.Errorf("error: start not integer (%s)", err)
	} else {
		startIndex = tempStart
	}
	if tempEnd, err := strconv.Atoi(transform.KWArgs["end"].(string)); err != nil {
		return "", fmt.Errorf("error: end not integer (%s)", err)
	} else {
		if tempEnd >= 0 {
			endIndex = tempEnd
		}
		if tempEnd < 0 {
			endIndex = len(recordsIn) + tempEnd
		}
	}

	for i := startIndex; i <= endIndex; i++ {
		record := recordsIn[i]
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		fmt.Printf("%s\n", record)
		recordsOut = append(recordsOut, record)
	}

	var b bytes.Buffer
	csvDataBuf := bufio.NewWriter(&b)
	csvWriter := csv.NewWriter(csvDataBuf)

	csvWriter.WriteAll(recordsOut)

	csvWriter.Flush()

	fmt.Printf("%v\n", recordsOut)

	return b.String(), nil
}

var operationMap = map[string](func(Transform, string) (string, error)){
	"none":       noneOp,
	"slice_rows": sliceRowsOp,
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
