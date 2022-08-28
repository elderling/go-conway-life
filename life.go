package main

import "bytes"

type PlayField [][]bool

func (p PlayField) String() string {
	var buf bytes.Buffer

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			if p[i][j] {
				buf.WriteByte('0')
			} else {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}
