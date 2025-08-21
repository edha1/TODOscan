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
	"strconv"
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
var fixmeRegex = regexp.MustCompile(`(?i)fixme`)

func main() {

	// CLI flags
	path := flag.String("path", ".", "Path to scan for TODOs")
	ext := flag.String("ext", ".java", "File extension to include (e.g., .java, .go)")
	olderThan := flag.Int("olderthan", 0, "Only show TODOs older than N days")
	oldestFirst := flag.Bool("oldestFirst", true, "Show in order of oldest first")

	flag.Parse()

	normalizedExt := strings.TrimSpace(*ext)
	if !strings.HasPrefix(normalizedExt, ".") {
		normalizedExt = "." + normalizedExt
	}

	var todos []Todo

	// WalkDir takes the string search path and a callback function
	filepath.WalkDir(*path, func(pathName string, d os.DirEntry, err error) error {

		if err != nil || d.IsDir() || (*ext != "" && !strings.HasSuffix(pathName, strings.ToLower(*ext))) {
			return nil
		}

		file, err := os.Open(pathName)

		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			line := scanner.Text()
			if todoRegex.MatchString(line) || fixmeRegex.MatchString(line) {
				date := blameDate(pathName, lineNum)
				cutoff := time.Now().AddDate(0, 0, -*olderThan)
				if *olderThan > 0 && date.Before(cutoff) {
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

		// sort in order required
		if *oldestFirst {
			return todos[i].Date.Before(todos[j].Date)
		} else {
			return todos[i].Date.After(todos[j].Date)
		}
	})

	for _, t := range todos {
		fmt.Printf("%s:%d [%s] %s\n", t.FilePath, t.LineNum, t.Date.Format("2006-01-02"), t.Text)
	}
}

func blameDate(path string, line int) time.Time {
	cmd := exec.Command("git", "blame", "--porcelain", "-L", fmt.Sprintf("%d,%d", line, line), path)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("git blame error:", err)
		return time.Time{}
	}

	lines := bytes.Split(output, []byte("\n"))
	for _, line := range lines {
		if bytes.HasPrefix(line, []byte("author-time ")) {
			tsStr := string(bytes.TrimPrefix(line, []byte("author-time ")))
			ts, err := strconv.ParseInt(tsStr, 10, 64)
			if err != nil {
				return time.Time{}
			}
			return time.Unix(ts, 0)
		}
	}
	return time.Time{}
}
