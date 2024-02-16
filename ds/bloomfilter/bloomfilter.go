package bloomfilter

import (
	"bytes"
	"encoding/binary"
	"goalds/al/hash"
	"goalds/ds/bitmap"
	"goalds/utils/locker"
	"math"
	"sync"
)

const salt = "goalds"

var defaultLocker locker.FakeLocker

type Options struct {
	locker locker.Locker
}

type Option func(opt *Options)

func WithGoroutineSafe() Option {
	return func(opt *Options) {
		opt.locker = &sync.RWMutex{}
	}
}

type Bloomfilter struct {
	m uint64
	k uint64
	b *bitmap.Bitmap
	l locker.Locker
}

func New(m uint64, k uint64, opts ...Option) *Bloomfilter {
	opt := &Options{
		locker: defaultLocker,
	}
	for _, o := range opts {
		o(opt)
	}
	return &Bloomfilter{
		m: m,
		k: k,
		b: bitmap.New(m),
		l: opt.locker,
	}
}

// NewWithEstimates creates a new Bloom filter with n and fp
// n is the capacity of the bloom filter
// fp is the tolerable error rate of the bloom filter
func NewWithEstimates(n uint64, fp float64, opts ...Option) *Bloomfilter {
	m, k := EstimateParameters(n, fp)
	return New(m, k, opts...)
}

func NewFromData(data []byte, opts ...Option) *Bloomfilter {
	opt := &Options{
		locker: defaultLocker,
	}
	for _, o := range opts {
		o(opt)
	}
	b := &Bloomfilter{
		l: opt.locker,
	}
	reader := bytes.NewReader(data)
	binary.Read(reader, binary.LittleEndian, &b.m)
	binary.Read(reader, binary.LittleEndian, &b.k)
	b.b = bitmap.NewFromBits(data[16:])
	return b
}

func EstimateParameters(n uint64, fp float64) (uint64, uint64) {
	m := uint64(math.Ceil(-1 * float64(n) * math.Log(fp) / (math.Ln2 * math.Ln2)))
	k := uint64(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return m, k
}

func (bf *Bloomfilter) Add(val string) {
	bf.l.Lock()
	defer bf.l.Unlock()
	hashs := hash.GetHashInts([]byte(salt+val), int(bf.k))
	for i := uint64(0); i < bf.k; i++ {
		bf.b.Set(hashs[i] % bf.m)
	}
}

func (bf *Bloomfilter) Test(val string) bool {
	bf.l.RLock()
	defer bf.l.RUnlock()
	hashs := hash.GetHashInts([]byte(salt+val), int(bf.k))
	for i := uint64(0); i < bf.k; i++ {
		if !bf.b.IsSet(hashs[i] % bf.m) {
			return false
		}
	}
	return true
}

func (bf *Bloomfilter) Data() []byte {
	bf.l.RLock()
	defer bf.l.RUnlock()
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, bf.m)
	binary.Write(buf, binary.LittleEndian, bf.k)
	buf.Write(bf.b.Data())
	return buf.Bytes()
}
