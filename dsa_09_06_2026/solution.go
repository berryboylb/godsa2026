// DSA Daily — 2026-06-09
// Problem: https://leetcode.com/problems/word-search/description/

package dsa09062026

func backtrack(row int, col int, word string, index int, rows int, cols int, board [][]byte) bool {
	if index >= len(word) {
		return true
	}

	if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] {
		return false
	}

	rowDir, colDir, ret := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}, false
	board[row][col] = '#'

	for i := 0; i < 4; i++ {
		ret = backtrack(row+rowDir[i], col+colDir[i], word, index+1, rows, cols, board)
		if ret {
			break
		}
	}

	board[row][col] = word[index]

	return ret

}

func solution(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if backtrack(row, col, word, 0, rows, cols, board) {
				return true
			}
		}
	}

	return false
}

func main() {}
