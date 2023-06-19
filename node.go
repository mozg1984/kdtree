package kdtree

import "sort"

type Node struct {
	point Point
	left  *Node
	right *Node
}

func (node *Node) compare(p Point, depth int) int {
	dim := depth % 2
	if p[dim] <= node.point[dim] {
		return -1
	}
	return 1
}

func within(node *Node, p Point, r float64, points *[]Point, depth int) {
	if node == nil {
		return
	}

	if node.compare(Point{p[0] - r, p[1] - r}, depth) > 0 {
		within(node.right, p, r, points, depth+1)
		return
	}

	if node.compare(Point{p[0] + r, p[1] + r}, depth) < 0 {
		within(node.left, p, r, points, depth+1)
		return
	}

	if p.Distance(node.point) <= r {
		*points = append(*points, node.point)
	}

	within(node.left, p, r, points, depth+1)
	within(node.right, p, r, points, depth+1)
	return
}

func build(tree *Node, points []Point, depth int) {
	if len(points) == 0 {
		return
	}

	axis := depth % 2

	sort.Slice(points, func(i, j int) bool {
		return points[i][axis] < points[j][axis]
	})

	pivot := len(points) / 2
	for pivot < len(points)-1 && points[pivot][axis] == points[pivot+1][axis] {
		pivot++
	}

	tree.point = points[pivot]
	build(tree.left, points[:pivot], depth+1)
	build(tree.right, points[pivot+1:], depth+1)
}
