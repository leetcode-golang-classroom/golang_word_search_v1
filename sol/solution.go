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
