package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	digitRegex := regexp.MustCompile(`\d\)`)

	if len(args) != 2 {
		fmt.Println("Usage: program <input_filename.txt> <output_filename.txt>")
		return
	}

	inputFileName := args[0]
	outputFileName := args[1]

	err := processFile(inputFileName, outputFileName, digitRegex)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func processFile(inputFileName, outputFileName string, digitRegex *regexp.Regexp) error {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer inputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		return fmt.Errorf("input file is empty")
	}

	scanner := bufio.NewScanner(inputFile)
	var modifiedText strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := processLine(line, digitRegex)
		modifiedText.WriteString(modifiedLine + "\n")
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning file: %v", err)
	}

	err = writeFile(outputFileName, modifiedText.String())
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	return nil
}

func processLine(line string, digitRegex *regexp.Regexp) string {
	words := split(line)

	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		words[i] = strings.ReplaceAll(words[i], " ", "")

		if i != 0 && (words[i] == "(cap)" ||
			words[i] == "(low)" ||
			words[i] == "(up)" ||
			words[i] == "(hex)" ||
			words[i] == "(bin)") {
			condition := words[i]
			applying(&words, condition, 1, i)
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if i != 0 && i != len(words) &&
			(words[i] == "(cap,") ||
			words[i] == "(low," ||
			words[i] == "(up," {

			if i+1 != len(words) && digitRegex.MatchString(words[i+1]) {
				condition := words[i]
				intchar := strings.ReplaceAll(words[i+1], ")", "")
				in, err := strconv.Atoi(intchar)
				if err != nil {
					fmt.Printf("%s is not a number\n", intchar)
					break
				}
				temp := append(words[:i])
				if in > len(temp) {
					fmt.Printf("you can't apply changes on the string\n(%s)\ncouse the number of words it contains is less than  %d\n", join(temp), in)
				} else {
					applying(&words, condition, in, i)
				}
				if i+1 == len(words)-1 {
					words = append(words[:0], words[:i]...)
				} else {
					words = append(words[:i], words[i+2:]...)
				}
			}
		}
	}

	joined := punctuation(join(words))
	joined = Correct(vowels(joined))
	return joined
}

func applying(strs *[]string, s string, a int, index int) {
	switch {
	case strings.Contains(s, "cap"):
		for i := index - a; i < index; i++ {
			Cap(&(*strs)[i])
		}
	case strings.Contains(s, "low"):
		for i := index - a; i < index; i++ {
			Low(&(*strs)[i])
		}
	case strings.Contains(s, "up"):
		for i := index - a; i < index; i++ {
			Up(&(*strs)[i])
		}
	case strings.Contains(s, "hex"):
		for i := index - a; i < index; i++ {
			Hex(&(*strs)[i])
		}
	case strings.Contains(s, "bin"):
		for i := index - a; i < index; i++ {
			Bin(&(*strs)[i])
		}
	}
}

// Other functions remain unchanged

func Bin(s *string) {
	a, err := strconv.ParseInt(*s, 2, 64)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	*s = strconv.Itoa(int(a))
}

func Cap(s *string) {
	r := []rune(*s)
	if r[0] <= 'z' && r[0] >= 'a' {
		r[0] -= 32
	}
	*s = string(r)
}

func Hex(s *string) {
	a, err := strconv.ParseInt(*s, 16, 64)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	*s = strconv.Itoa(int(a))
}

func Low(s *string) {
	r := []rune(*s)
	for i := 0; i < len(r); i++ {
		if r[i] <= 'Z' && r[i] >= 'A' {
			r[i] += 32
		}
	}
	*s = string(r)
}

func Up(s *string) {
	r := []rune(*s)
	for i := 0; i < len(r); i++ {
		if r[i] <= 'z' && r[i] >= 'a' {
			r[i] -= 32
		}
	}
	*s = string(r)
}

/* ---- punctuation ----*/

func punctuation(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if i != 0 && (r[i] == '.' || r[i] == ',' || r[i] == '!' || r[i] == '?' || r[i] == ':' || r[i] == ';') {
			if i != len(r)-1 && r[i-1] == ' ' && r[i+1] != ' ' {
				r[i-1], r[i] = r[i], r[i-1]
				i--
			}
			if i < len(r)-1 && r[i-1] != ' ' && r[i+1] != ' ' {
				r = append(r[:i+1], append([]rune{' '}, r[i+1:]...)...)
			}
			if r[i-1] == ' ' {
				r = append(r[:i-1], r[i:]...)
			}
		}
	}
	return string(r)
}

/* ---- removing space after and befor the single quote ------ */

func Correct(s string) string {
	count := 1
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if count%2 != 0 {
			if r[i] == '\'' && i < len(r)-1 {
				if r[i+1] == ' ' {
					r = append(r[:i+1], r[i+2:]...)
					count++
					continue
				}
			}
		}
		if count%2 == 0 {
			if r[i] == '\'' {
				if r[i-1] == ' ' {
					r = append(r[:i-1], r[i:]...)
					count++
				}
			}
		}

	}
	return string(r)
}

/* ------ adding (n) to a if there is a vauol after it ---- */

func vowels(s string) string {
	r := []rune(s)
	var result []rune

	for i := 0; i < len(r); i++ {
		result = append(result, r[i])
		if (r[i] == 'a' || r[i] == 'A') && i < len(r)-1 && (r[i+1] == ' ') && (i+2 < len(r)) &&
			(r[i+2] == 'a' || r[i+2] == 'e' || r[i+2] == 'i' || r[i+2] == 'o' || r[i+2] == 'u' || r[i+2] == 'A' || r[i+2] == 'E' || r[i+2] == 'I' || r[i+2] == 'O' || r[i+2] == 'U' || r[i+2] == 'H' || r[i+2] == 'h') {
			if i == 0 {
				result = append(result, 'n')
			} else if r[i-1] == ' ' || r[i-1] == '"' || r[i-1] == '\'' {
				result = append(result, 'n')
			}
		}
	}
	return string(result)
}

func split(s string) []string {
	var words []string
	word := ""
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] == '\n' && word == "" {
			words = append(words, "\n")
		} else if r[i] == ' ' && word != "" {
			words = append(words, word)
			word = ""
		} else if r[i] != ' ' && r[i] != '\n' {
			word += string(r[i])
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func join(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		if strs[i] == "\n" {
			s += "\n"
		} else {
			s += " " + strs[i]
		}
	}
	return s
}

func writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
