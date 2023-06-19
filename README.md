There is one of the kD-tree implementation for only 2 dimension case (X, Y).
It uses splitting the data into two parts by a median point which allows to maintain balance when building a kD-tree.

```go
points := []Point{{48.381124, 54.310801}}
p := Point{48.396230, 54.312595}

index := NewIndex()
index.Build(points)

// Search for neighboring points in a given point and radius (m)
found := index.Within(p, 1010)
```
