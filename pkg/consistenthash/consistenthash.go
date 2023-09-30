package consistenthash

// HashFunc defines function to generate hash code
type HashFunc func(data []byte) uint32

// NodeMap stores nodes and you can pick node from NodeMap
type NodeMap struct {
	hashFunc    HashFunc       // hash func
	nodeHashs   []int          // sorted
	nodeHashMap map[int]string // node
}
