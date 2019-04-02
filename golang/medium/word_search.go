/*
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell,
where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/

package main

import "fmt"

func deepCopy(record map[[2]int]bool, n_v [2]int) map[[2]int]bool {
	n_record := make(map[[2]int]bool)
	for k, v := range record {
		n_record[k] = v
	}
	n_record[n_v] = true

	return n_record
}

func fDFS(board [][]byte, r, c, index int, word string, record map[[2]int]bool) bool {
	if index == len(word) { //都访问完了
		return true
	}

	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
		return false
	}

	if board[r][c] != word[index] || record[[2]int{r, c}] { //字母不相等或者已经访问过
		return false
	} else { //相等，但不是最后一个字母

		record[[2]int{r, c}] = true

		ret := fDFS(board, r+1, c, index+1, word, record) ||
			fDFS(board, r-1, c, index+1, word, record) ||
			fDFS(board, r, c+1, index+1, word, record) ||
			fDFS(board, r, c-1, index+1, word, record)

		record[[2]int{r, c}] = false

		return ret
	}
}

func exist(board [][]byte, word string) bool {

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if fDFS(board, i, j, 0, word, map[[2]int]bool{}) {
				return true
			}
		}
	}

	return false
}

func main() {
	ret := exist([][]byte{{'a', 'a'}}, "aaa")
	fmt.Println(ret)
}
