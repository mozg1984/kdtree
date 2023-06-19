package kdtree

type Index struct {
	tree *Node
}

func NewIndex() *Index {
	return &Index{}
}

func (index *Index) Build(points []Point) {
	index.tree = &Node{}
	build(index.tree, points, 0)
}

func (index *Index) Within(p Point, r float64) (points []Point) {
	within(index.tree, p, r, &points, 0)
	return
}
