package util

import "os"

type funcLog func(v ...any)

func FecharArquivoL(f *os.File, l funcLog) {
	err := f.Close()
	if err != nil && l != nil {
		l(err)
	}
}
