package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	matrix      [][]int
	rowCount    int
	columnCount int
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.rowCount {
		return false
	}
	if col < 0 || col >= m.columnCount {
		return false
	}

	if m.matrix[row] == nil {
		m.matrix[row] = make([]int, m.columnCount)
	}

	m.matrix[row][col] = val
	return true
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.rowCount)
	for row := 0; row < m.rowCount; row++ {
		rows[row] = make([]int, m.columnCount)
		copy(rows[row], m.matrix[row])
	}
	return rows
}

func (m *Matrix) Cols() [][]int {
	transposed := Init(m.columnCount, m.rowCount)
	for col := 0; col < m.columnCount; col++ {
		for row := 0; row < m.rowCount; row++ {
			transposed.Set(col, row, m.matrix[row][col])
		}
	}
	return transposed.Rows()
}

func Init(rows, cols int) (m *Matrix) {
	m = &Matrix{rowCount: rows, columnCount: cols}
	m.matrix = make([][]int, rows)
	return
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	m := Init(len(rows), len(strings.Fields(rows[0])))
	for row := 0; row < m.rowCount; row++ {
		cells := strings.Fields(rows[row])
		if len(cells) != m.columnCount {
			return nil, errors.New("not a rectangular matrix")
		}

		for col := 0; col < m.columnCount; col++ {
			cell, err := strconv.Atoi(cells[col])
			if err != nil {
				return nil, err
			}

			m.Set(row, col, cell)
		}
	}
	return m, nil
}
