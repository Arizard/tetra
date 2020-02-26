package tetra

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strings"
	"time"
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

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func mergeColumnsOp(transform Transform, csvData string) (string, error) {

	var recordsOut [][]string

	reader := readerFromData(csvData, transform)

	colSrcIndex := int(transform.KWArgs["index_from"].(float64))
	colDstIndex := int(transform.KWArgs["index_to"].(float64))

	recordsIn, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	for _, record := range recordsIn {
		newRecord := record
		if len(newRecord) - 1 >= max(colSrcIndex, colDstIndex) {
			srcCell := newRecord[colSrcIndex]
			if srcCell != "" {
				newRecord[colDstIndex] = srcCell
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

func ignoreRowsWhereColumnEqualsOp(transform Transform, csvData string) (string, error) {

	var recordsOut [][]string

	reader := readerFromData(csvData, transform)

	colIndex := int(transform.KWArgs["index"].(float64))
	colMatch := transform.KWArgs["match"].(string)

	recordsIn, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	for _, record := range recordsIn {
		newRecord := record
		if len(newRecord) - 1 < colIndex { continue }
		if newRecord[colIndex] == colMatch { continue }
		recordsOut = append(recordsOut, newRecord)
	}

	var b bytes.Buffer

	csvDataBuf := bufio.NewWriter(&b)

	csvWriter := csv.NewWriter(csvDataBuf)
	csvWriter.WriteAll(recordsOut)
	csvWriter.Flush()

	return b.String(), nil
}

func ignoreRowsWhereColumnInFuture(transform Transform, csvData string) (string, error) {
	format := transform.KWArgs["format"].(string)
	colIndex := int(transform.KWArgs["index"].(float64))

	timezoneName := transform.KWArgs["timezone"].(string)
	timezone, err := time.LoadLocation(timezoneName)
	if err != nil { panic(err) }

	var recordsOut [][]string

	reader := readerFromData(csvData, transform)

	recordsIn, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	for _, record := range recordsIn {

		if len(record) - 1 < colIndex { continue }

		cell := record[colIndex]

		cellTime, err := time.Parse(format, cell)
		if err != nil { panic(err) }
		cellTimeLocalised := cellTime.In(timezone)

		now := time.Now().In(timezone)

		if cellTimeLocalised.After(now) { continue }

		recordsOut = append(recordsOut, record)
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
	"merge_columns": mergeColumnsOp,
	"ignore_rows_where_column_equals": ignoreRowsWhereColumnEqualsOp,
	"ignore_rows_where_column_in_future": ignoreRowsWhereColumnInFuture,
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
