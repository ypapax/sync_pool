package main

import (
	"fmt"
	"sync"
	"testing"
)

//https://stackoverflow.com/questions/72893095/how-to-reuse-slice-with-sync-pool-in-golang/72893375#72893375

func TestSliceWithPool(t *testing.T) {
	var slicePool = sync.Pool{
		New: func() interface{} {
			t.Log("Created")
			s := make([]interface{}, 0)
			return s
		},
	}

	s, _ := slicePool.Get().([]interface{})
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
	// ^ output: [0 1 2 3 4 5 6 7 8 9]

	s = s[:0]
	slicePool.Put(s)

	s2, _ := slicePool.Get().([]interface{})
	fmt.Println(s)
	// ^ output: []

	for i := 0; i < 5; i++ {
		s2 = append(s2, i*10)
	}
	fmt.Println(s2)
	// ^ output: [0 10 20 30 40]

	fmt.Println(s2[:10])
	// ^ output: [0 10 20 30 40 5 6 7 8 9]
}
