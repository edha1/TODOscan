package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Todo struct {
	FilePath string
	LineNum  int
	Text     string
	Date     time.Time
}

var todoRegex = regexp.MustCompile(`(?i)todo`)

func main() {

	// CLI flags
	path := flag.String("path", ".", "Path to scan for TODOs")
	ext := flag.String("ext", ".go", "File extension to include")
	sinceDays := flag.Int("since", 0, "Only show TODOs older than N days")

	flag.Parse()

	var todos []Todo

	// WalkDir takes the string search path and a callback function
	filepath.WalkDir(*path, func(pathName string, d os.DirEntry, err error) error {

		// base cases, ignore errors, directories and files not of the required type (.go/.java)
		if err != nil || d.IsDir() || !strings.HasSuffix(pathName, *ext) {
			return nil
		}

		file, err := os.Open(pathName) // open file

		// return if error opening file
		if err != nil {
			return nil
		}
		defer file.Close() // close the file when the surrounding block has been run

		scanner := bufio.NewScanner(file) // open a scanner
		lineNum := 1
		for scanner.Scan() {
			line := scanner.Text()
			if todoRegex.MatchString(line) {
				date := blameDate(pathName, lineNum)
				cutoff := time.Now().AddDate(0, 0, -*sinceDays)
				if *sinceDays > 0 && date.Before(cutoff) {
					// skip TODOs newer than the cutoff
					lineNum++
					continue
				}
				todos = append(todos, Todo{
					FilePath: pathName,
					LineNum:  lineNum,
					Text:     strings.TrimSpace(line),
					Date:     date,
				})
			}
			lineNum++
		}
		return nil
	})

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Date.Before(todos[j].Date)
	})

	for _, t := range todos {
		fmt.Printf("%s:%d [%s] %s\n", t.FilePath, t.LineNum, t.Date.Format("2006-01-02"), t.Text)
	}
}

// Function to get the date of a comment using 'git blame'
func blameDate(path string, line int) time.Time {
	cmd := exec.Command("git", "blame", "--line-porcelain", fmt.Sprintf("+%d,%d", line, line), path) // line porcelain gives key value pairs
	output, err := cmd.Output()                                                                      // get output
	if err != nil {
		return time.Time{} // return zero value
	}
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() { // while there is input lines to read
		line := scanner.Text()

		// see time and convert as required
		if strings.HasPrefix(line, "author-time ") {
			ts := strings.TrimPrefix(line, "author-time ")
			sec, _ := time.ParseDuration(ts + "s")
			return time.Unix(int64(sec.Seconds()), 0)
		}
	}
	return time.Time{}
}
