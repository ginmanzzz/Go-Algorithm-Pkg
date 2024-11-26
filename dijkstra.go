package algo

import (
	"math"
)

const inf = math.MaxInt / 2

type Dijkstra struct {
	g [][]int
}

func NewDijkstra(edges [][]int, n int, isDirect bool) *Dijkstra {
	return &Dijkstra {
		g: buildGraph(edges, n, isDirect),
	}
}

func buildGraph(edges [][]int, n int, isDirect bool) [][]int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		w := edge[2]
		g[x][y] = w
		if !isDirect {
			g[y][x] = w
		}
	}
	return g
}

/*
  @params: 	g is the builed graph(directed graph or undirected graph).
  			k is the origin point.
  @brief:  	Pure dijkstra algorithm is used to find the shortest way between two points, or one point to others.
  			It can be used in positive weight graph and dense graph.
  @returns: distance is all distances start from the origin point to others.
  			can is whether the origin point can arrive at all others.
  @time:   	O(N^2)
  @space:  	O(N^2)
*/
func (dijkstra *Dijkstra) PlainAll(g [][] int, k int) ([]int,bool) {
	n := len(g)
	distance := make([]int, n)
	for i := range distance {
		distance[i] = inf
	}
	distance[k] = 0
	can := false
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || distance[i] < distance[x]) {
				x = i
			}
		}
		if x < 0 {
			can = true
			break
		}
		if distance[x] >= inf {
			can = false
			break
		}
		done[x] = true
		for y, d := range g[x] {
			distance[y] = MinInt(distance[y], distance[x] + d)
		}
	}
	return distance, can
}

func (dijkstra *Dijkstra) PlainE2E(g [][]int, start , end int) int {
	distance, _ := dijkstra.PlainAll(dijkstra.g, start)
	if distance[end] >= inf  {
		return -1
	}
	return distance[end]
}
