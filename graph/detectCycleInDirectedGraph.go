package graph

//DetectCycle : Given a directed graph we need to detect cycle in it
func DetectCycle(directedGraph map[int][][]int) {
	vertices := make(map[int]bool)
	for key := range directedGraph {
		vertices[key] = false
	}
}
