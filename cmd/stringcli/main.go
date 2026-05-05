package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"petgo/internal/stringapp"
)

func romanToInt(s string) int {
	srune := []rune(s)
	alphabet := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	count := 0

	for i:=0; i < len(srune); i++ {
		n := alphabet[string(srune[i])]

		if i+1 < len(srune) && n < alphabet[string(srune[i+1])] {
			count -= n
		} else {
			count += n
		}
	}

	fmt.Println("count roman to int", count)
	return 0
}

// {"flower","flow","flight"}
func longestCommonPrefix(strs []string) string {
	prefix := ""
	count := 0

	for i := 0; i < len(strs); i++ {
		if count > len(strs[i]) {
			return prefix
		}

		for _,value := range strs[i] {
			if i+1 < len(strs) && count < len(strs[i]) && count < len(strs[i+1]) && strs[i][count] == strs[i+1][count] {
				prefix += string(value)
				count++
			}
		}
	}
	fmt.Println("prefix", prefix)
	fmt.Println("count", count)
	return prefix
}

func concat(values []string) string {
    s := ""
    for _, v := range values {
        s += v
    }
    return s
}

func main() {
	var (
		input      string
		daemonMode bool
	)

	flag.StringVar(&input, "input", "", "input string")
	flag.BoolVar(&daemonMode, "daemon", false, "run in daemon mode")
	flag.Parse()

	if daemonMode {
		runDaemon()
		return
	}

	if input == "" {
		exitWithError(fmt.Errorf("передайте строку через --input или запустите --daemon"))
	}

	result, err := stringapp.Unpack(input)
	if err != nil {
		exitWithError(err)
	}

	fmt.Println(result)

	s := "hello"
	s1 := ""
	r := []rune(s)
	// r1 := []rune(s1)
	fmt.Printf("length s %c\n", r[0])
	fmt.Printf("length s1 %d\n", len(s1))
	fmt.Printf("concat %v\n", concat([]string{"b", "a", "r"}))

	stringsPrefix :=[]string{"flower","flow","flight"}
	// v := "hello"
	// vrune := []rune(v)
	// symbol := []rune("12")[0]
	// fmt.Println("result 222", vrune[symbol])
	// result = CountFirstSymbol("gbbccdda")
	fmt.Println("result", result)
	romanToInt("XXI")
	fmt.Printf("stringsPrefix=%T\n", stringsPrefix[0])
	longestCommonPrefix(stringsPrefix)
}

func runDaemon() {
	fmt.Println("Ctrl+C для завершения")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите строку: ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				exitWithError(err)
			}
			return
		}

		result, err := stringapp.Unpack(scanner.Text())
		if err != nil {
			fmt.Printf("ошибка: %v\n", err)
			continue
		}

		fmt.Println(result)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}