package merkle


import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)
/*
func TestNewMerkleNode(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
	}

	// Level 1

	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	// Level 2
	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	// Level 3
	n7 := NewMerkleNode(n5, n6, nil)

	assert.Equal(
		t,
		"64b04b718d8b7c5b6fd17f7ec221945c034cfce3be4118da33244966150c4bd4",
		hex.EncodeToString(n5.Data),
		"Level 1 hash 1 is correct",
	)
	assert.Equal(
		t,
		"08bd0d1426f87a78bfc2f0b13eccdf6f5b58dac6b37a7b9441c1a2fab415d76c",
		hex.EncodeToString(n6.Data),
		"Level 1 hash 2 is correct",
	)
	assert.Equal(
		t,
		"4e3e44e55926330ab6c31892f980f8bfd1a6e910ff1ebc3f778211377f35227e",
		hex.EncodeToString(n7.Data),
		"Root hash is correct",
	)
}*/

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
		[]byte("node4"),
		[]byte("node5"),
	}
	// Level 1
	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[3])
	n5 := NewMerkleNode(nil, nil, data[4])


	// Level 2
	n21 := NewMerkleNode(n1, n2, nil)
	n22 := NewMerkleNode(n3, n4, nil)
	n23 := NewMerkleNode(n5, n5, nil)

	// Level 3
	n31 := NewMerkleNode(n21, n22, nil)
	n32 := NewMerkleNode(n23, n23, nil)

	// Level
	n41 := NewMerkleNode(n31, n32, nil)


	fmt.Printf("rootHash: %x\n", n41.Data)

	mTree := NewMerkleTree(data)
	fmt.Printf("mTree: %x\n", mTree.RootNode.Data)
	assert.Equal(t, fmt.Sprintf("%x", n41.Data), fmt.Sprintf("%x", mTree.RootNode.Data), "Merkle tree root hash is correct")
}