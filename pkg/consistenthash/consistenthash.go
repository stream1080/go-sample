package consistenthash

import "hash/crc32"

// HashFunc defines function to generate hash code
type HashFunc func(data []byte) uint32

// NodeMap stores nodes and you can pick node from NodeMap
type NodeMap struct {
	hashFunc    HashFunc       // hash func
	nodeHashs   []int          // sorted
	nodeHashMap map[int]string // node
}

// NewNodeMap creates a new NodeMap
func NewNodeMap(fn HashFunc) *NodeMap {

	if fn == nil {
		fn = crc32.ChecksumIEEE
	}

	return &NodeMap{
		hashFunc:    fn,
		nodeHashMap: make(map[int]string),
	}
}

// IsEmpty returns if there is no node in NodeMap
func (m *NodeMap) IsEmpty() bool {
	return len(m.nodeHashs) == 0
}
