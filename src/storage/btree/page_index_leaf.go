package btree

/*
import (
	"github.com/aethne0/assert"
	"github.com/pierrec/xxHash/xxHash64"
)
*/

type PageIndexLeafHeader struct {}

// Leafpage is a "view" of the bytes in a page, it reads/modifies in place (other than the header, for convenience)
type PageIndexLeaf struct {
	header	PageIndexLeafHeader;
	raw		*[PAGE_SIZE]byte;
}

// Writeheader writes the header struct into the raw field of bytes.
// This must be called before writing out the page.
func (lp *PageIndexLeaf) writeHeader() {
	// todo

}


// HEADER LAYOUT
// 0x00:  0001 0203 0405 0607  0809 0a0b 0c0d 0e0f  1011 1213 1415 1617  1819 1a1b 1c1d 1e1f
// Entry  AAAA AAAA AAAA AAAA  BBBB BBBB BBBB BBBB	CCDD
// 0x20:  2021 2223 2425 2627  2829 2a2b 2c2d 2e2f  3031 3233 3435 3637  3839 3a3b 3c3d 3e3f
// Entry  -
// 
// A -> 64 Checksum
// B -> 64 PageID
// C ->  8 Page type
// D ->  8 Flags (todo)
// 

