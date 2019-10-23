package toy

//go:generate mockgen -self_package github.com/dbyington/actiony/toy -package toy -source box.go -destination mock_box_test.go

import "errors"

type box struct {
	height, width, length int64
	isFull bool
}

var (
	ErrFull = errors.New("box already full")
	ErrEmpty = errors.New("box already empty")
)

type Box interface {
	Empty() error
	Fill() error
	Size() (height, width, length int64)
}

// NewBox returns a box with the optional arguments of `height`, `width`, `length` (in that order).
func NewBox (dim ...int64) *box {
	if len(dim) == 3 {
		return &box{
			height: dim[0],
			width: dim[1],
			length: dim[2],
		}
	}
	return &box{}
}

func (b *box) Fill() error {
	if b.isFull == true {
		return ErrFull
	}
	b.isFull = true
	return nil
}

func (b *box) Empty() error {
	if b.isFull == false {
		return ErrEmpty
	}
	b.isFull = false
	return nil
}

func (b *box) Size() (height, width, length int64) {
	return b.height, b.width, b.length
}