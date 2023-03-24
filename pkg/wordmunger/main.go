package wordmunger

import (
	"bufio"
	"log"
	"os"
)

type WordMunger struct {
	WordTarget []string
	Concurrent int
	InputFile  string
	OutputFile string
}

func (wm *WordMunger) ReadFile() {
	curdir, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	var filePath string
	if len(wm.InputFile) < len(curdir) {
		filePath = curdir + "/" + wm.InputFile
	} else {
		filePath = wm.InputFile
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	// for each line in file add it to WordTarget slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// check if the line is not empty then add it to WordTarget slice
		if len(scanner.Text()) > 0 {
			wm.WordTarget = append(wm.WordTarget, scanner.Text())
		}
	}
}
