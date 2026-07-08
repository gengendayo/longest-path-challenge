package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// エッジ（路線）の構造体
type Edge struct {
	To     int
	Weight float64
}

var graph map[int][]Edge
var maxDist float64
var bestPath []int

func main() {
	graph = make(map[int][]Edge)

	// 1. 標準入力からの読み込み
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// 始点, 終点, 距離のフォーマットを処理
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		// 任意の数のホワイトスペースが含まれることを考慮してトリム
		uStr := strings.TrimSpace(parts[0])
		vStr := strings.TrimSpace(parts[1])
		wStr := strings.TrimSpace(parts[2])

		u, err1 := strconv.Atoi(uStr)
		v, err2 := strconv.Atoi(vStr)
		w, err3 := strconv.ParseFloat(wStr, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		// 無向グラフとしてエッジを追加（双方向移動可能という推論に基づく）
		graph[u] = append(graph[u], Edge{To: v, Weight: w})
		graph[v] = append(graph[v], Edge{To: u, Weight: w})
	}

	nodes := make([]int, 0, len(graph))
	for k := range graph {
		nodes = append(nodes, k)
	}

	maxDist = -1.0
	bestPath = nil

	// 2. 任意の始点から探索を開始
	for _, startNode := range nodes {
		visited := make(map[int]bool)
		visited[startNode] = true
		currentPath := []int{startNode}
		dfs(startNode, 0.0, visited, currentPath)
	}

	// 3. 結果の出力
	var out []string
	for _, nodeID := range bestPath {
		out = append(out, strconv.Itoa(nodeID))
	}
	fmt.Print(strings.Join(out, "\r\n"))
	fmt.Print("\r\n") // 末尾の改行
}

func dfs(current int, currentDist float64, visited map[int]bool, currentPath []int) {
	if currentDist > maxDist {
		maxDist = currentDist
		bestPath = make([]int, len(currentPath))
		copy(bestPath, currentPath)
	}

	for _, edge := range graph[current] {
		// 同じ点を2回通らないようにチェック
		if !visited[edge.To] {
			visited[edge.To] = true
			currentPath = append(currentPath, edge.To)

			dfs(edge.To, currentDist+edge.Weight, visited, currentPath)

			// バックトラッキング
			visited[edge.To] = false
			currentPath = currentPath[:len(currentPath)-1]
		}
	}
}