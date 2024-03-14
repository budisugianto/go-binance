package common

import (
	"bytes"
	"math"
	"sync"
)

// AmountToLotSize converts an amount to a lot sized amount
func AmountToLotSize(lot float64, precision int, amount float64) float64 {
	return math.Trunc(math.Floor(amount/lot)*lot*math.Pow10(precision)) / math.Pow10(precision)
}

// ToJSONList convert v to json list if v is a map
func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}

type Int64c struct {
	sync.RWMutex
	data int64
}

func (v *Int64c) Set(d int64) {
	v.Lock()
	defer v.Unlock()
	v.data = d
}

func (v *Int64c) Get() int64 {
	v.RLock()
	defer v.RUnlock()
	return v.data
}
