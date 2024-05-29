package ui

import "strings"

func Table(header []string, content [][]string) string {
	t := newTable(header, content)
	return t.build()
}

type table struct {
	header  []string
	content [][]string
}

func newTable(header []string, content [][]string) table {
	for _, c := range content {
		if len(c) != len(header) {
			panic("header and content do not have the same length")
		}
	}

	filteredHeader := []string{}
	filteredContent := make([][]string, len(content))
	lenEachCol := getMaxLenEachCol(content)
	for i, length := range lenEachCol {
		if length == 0 {
			continue
		}

		filteredHeader = append(filteredHeader, header[i])
		for j := range content {
			filteredContent[j] = append(filteredContent[j], content[j][i])
		}
	}

	return table{
		header:  filteredHeader,
		content: filteredContent,
	}
}

func (t table) build() string {
	result := ""

	maxLenEachCol := getMaxLenEachCol(append(t.content, t.header))

	alignedHeader := getAlignedRow(t.header, maxLenEachCol)
	result += strings.Join(alignedHeader, "  ") + "\n"

	divider := getDivider('-', getMaxLenEachCol([][]string{t.header}))
	alignedDivider := getAlignedRow(divider, maxLenEachCol)
	result += strings.Join(alignedDivider, "  ") + "\n"

	for _, row := range t.content {
		alignedRow := getAlignedRow(row, maxLenEachCol)
		result += strings.Join(alignedRow, "  ") + "\n"
	}

	return result
}

func getDivider(char rune, lenEachCol []int) []string {
	result := []string{}

	for i := 0; i < len(lenEachCol); i++ {
		col := ""
		for j := 0; j < lenEachCol[i]; j++ {
			col += string(char)
		}

		result = append(result, col)
	}

	return result
}

func getAlignedRow(row []string, lenEachCol []int) []string {
	if len(row) != len(lenEachCol) {
		panic("row and lenEachCol do not have the same length")
	}

	result := make([]string, len(row))

	for i, col := range row {
		result[i] = col

		spaceToFill := lenEachCol[i] - len(col)
		if spaceToFill < 0 {
			panic("col is greater than lenEachCol[i]")
		}

		for j := 0; j < spaceToFill; j++ {
			result[i] += " "
		}
	}

	return result
}

func getMaxLenEachCol(rows [][]string) []int {
	if len(rows) == 0 {
		return []int{}
	}

	result := make([]int, len(rows[0]))

	for _, row := range rows {
		for i, c := range row {
			if len(c) > result[i] {
				result[i] = len(c)
			}
		}
	}

	return result
}
