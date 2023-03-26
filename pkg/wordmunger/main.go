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

func (wm *WordMunger) Munging() []string {
	variations := []string{""}
	for _, x := range wm.WordTarget {
		for _, p := range x {
			lowers := []string{}
			for _, v := range variations {
				lowers = append(lowers, v+string(p))
			}
			if subs, ok := commonSubs[string(p)]; ok {
				for _, s := range subs {
					x := []string{}
					for _, v := range variations {
						x = append(x, v+s)
					}
					lowers = append(lowers, x...)
				}
			}
			variations = lowers
		}
	}
	return variations
}
