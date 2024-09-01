package encoder

import (
	"bufio"
	"io"
	"os"

	"github.com/friarhob/cchuffman/internal/adt"
	"github.com/friarhob/cchuffman/internal/exitcodes"
)

func runeCounter(file *os.File) *adt.DefaultMap {
	reader := bufio.NewReader(file)

	res := adt.NewDefaultMap(0)

	for {
		curRune, _, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}
			os.Exit(int(exitcodes.ErrorReadingFile))
		}

		res.Set(curRune, res.Get(curRune).(int)+1)
	}

	return res
}
