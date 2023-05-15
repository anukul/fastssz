// Code generated by fastssz. DO NOT EDIT.
// Hash: 4ebf31c9db2d40eb24702116641d6e98d75ebeb401ec69e41b7f351b81e25aa2
// Version: 0.1.3
package testcases

import (
	ssz "github.com/anukul/fastssz"
	"github.com/anukul/fastssz/sszgen/testcases/other"
)

// MarshalSSZ ssz marshals the Case3B object
func (c *Case3B) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case3B object to a target array
func (c *Case3B) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	return
}

// UnmarshalSSZ ssz unmarshals the Case3B object
func (c *Case3B) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 0 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case3B object
func (c *Case3B) SizeSSZ() (size int) {
	size = 0
	return
}

// HashTreeRoot ssz hashes the Case3B object
func (c *Case3B) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case3B object with a hasher
func (c *Case3B) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case3B object
func (c *Case3B) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}

// MarshalSSZ ssz marshals the Case3A object
func (c *Case3A) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case3A object to a target array
func (c *Case3A) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'A'
	if dst, err = c.A.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'B'
	if c.B == nil {
		c.B = new(Case3B)
	}
	if dst, err = c.B.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'C'
	if dst, err = c.C.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'D'
	if c.D == nil {
		c.D = new(other.Case3B)
	}
	if dst, err = c.D.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Case3A object
func (c *Case3A) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 0 {
		return ssz.ErrSize
	}

	// Field (0) 'A'
	if err = c.A.UnmarshalSSZ(buf[0:0]); err != nil {
		return err
	}

	// Field (1) 'B'
	if c.B == nil {
		c.B = new(Case3B)
	}
	if err = c.B.UnmarshalSSZ(buf[0:0]); err != nil {
		return err
	}

	// Field (2) 'C'
	if err = c.C.UnmarshalSSZ(buf[0:0]); err != nil {
		return err
	}

	// Field (3) 'D'
	if c.D == nil {
		c.D = new(other.Case3B)
	}
	if err = c.D.UnmarshalSSZ(buf[0:0]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case3A object
func (c *Case3A) SizeSSZ() (size int) {
	size = 0
	return
}

// HashTreeRoot ssz hashes the Case3A object
func (c *Case3A) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case3A object with a hasher
func (c *Case3A) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'A'
	if err = c.A.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'B'
	if c.B == nil {
		c.B = new(Case3B)
	}
	if err = c.B.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'C'
	if err = c.C.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'D'
	if c.D == nil {
		c.D = new(other.Case3B)
	}
	if err = c.D.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case3A object
func (c *Case3A) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
