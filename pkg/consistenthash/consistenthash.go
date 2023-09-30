package consistenthash

import (
	"hash/crc32"
	"sort"
)

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

// AddNode add the given nodes into consistent hash circle
func (m *NodeMap) AddNode(keys ...string) {

	for _, key := range keys {
		if key == "" {
			continue
		}
		hash := int(m.hashFunc([]byte(key)))
		m.nodeHashs = append(m.nodeHashs, hash)
		m.nodeHashMap[hash] = key
	}

	sort.Ints(m.nodeHashs)
}
