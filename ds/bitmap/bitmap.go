package bitmap

type Bitmap struct {
	data []byte
	size uint64
}

func New(size uint64) *Bitmap {
	size = (size + 7) / 8 * 8
	bitmap := &Bitmap{
		size: size,
		data: make([]byte, size/8),
	}
	return bitmap
}

func NewFromBits(data []byte) *Bitmap {
	bitmap := &Bitmap{
		size: uint64(len(data)) * 8,
		data: make([]byte, len(data)),
	}
	copy(bitmap.data, data)
	return bitmap
}

func (b *Bitmap) Set(pos uint64) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] |= 1 << (pos & 0x07)
	return true
}

func (b *Bitmap) Unset(pos uint64) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] &= ^(1 << (byte(pos) & 0x07))
	return true
}

func (b *Bitmap) IsSet(pos uint64) bool {
	if pos >= b.size {
		return false
	}
	if b.data[pos>>3]&(1<<(pos&0x07)) > 0 {
		return true
	}
	return false
}

func (b *Bitmap) Resize(size uint64) bool {
	size = (size + 7) / 8 * 8
	if b.size == size {
		return false
	}
	data := make([]byte, size/8)
	copy(data, b.data)
	b.data = data
	b.size = size
	return true
}

func (b *Bitmap) Size() uint64 {
	return b.size
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size/8)
}

func (b *Bitmap) Data() []byte {
	return b.data
}
