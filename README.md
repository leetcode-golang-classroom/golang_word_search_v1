# golang_word_search_v1

Given an `m x n` grid of characters `board` and a string `word`, return `true` *if* `word` *exists in the grid*.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/11/04/word2.jpg](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg](https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg)

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

```

**Example 3:**

![https://assets.leetcode.com/uploads/2020/10/15/word3.jpg](https://assets.leetcode.com/uploads/2020/10/15/word3.jpg)

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

```

**Constraints:**

- `m == board.length`
- `n = board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` and `word` consists of only lowercase and uppercase English letters.

## 解析

題目給了一個字元矩陣 board 與一個字串 word

要求實作一個演算法去判斷這個 word 是否有存在這個矩陣，注意的是每個字元組成字串的形式是水平或垂直鄰近的字元做連接並且只能使用一次。

矩陣為 m by n ，連結的方向有水平跟垂直

所以對於從每個矩陣的字元為開始，總共有 4 種接續的可能

因為最遭的狀況是 $(4^m)^n$

以有找到的狀況一般case 假設找到 k 字元 ，且字元長度是 p

則總搜尋時間是 k + $4^p$

k 介於 1 ~ m*n - 1

相鄰的查找方式需要使用 DFS 來做實作

如下圖

![](https://i.imgur.com/7Ji3ucf.png)

## 程式碼

```go
package sol

type Pair struct {
	row, col int
}

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])

	visit := make(map[Pair]struct{})
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		if idx == len(word) { // match all prefix
			return true
		}
		if r < 0 || r >= row || c < 0 || c >= col {
			return false
		}
		// revisited
		if _, visited := visit[Pair{row: r, col: c}]; visited {
			return false
		}
		// not match
		if board[r][c] != word[idx] {
			return false
		}
		// marked visited
		visit[Pair{row: r, col: c}] = struct{}{}

		if dfs(r-1, c, idx+1) || dfs(r+1, c, idx+1) || dfs(r, c-1, idx+1) || dfs(r, c+1, idx+1) {
			return true
		}
		// delete marked
		delete(visit, Pair{row: r, col: c})
		return false
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

```
## 困難點

1. 理解如何使用 DFS
2. 了解終止條件

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity