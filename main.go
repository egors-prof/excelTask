package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func writeToOut(resMap map[string]int) {

}
func task1() {
	file, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := fmt.Sprint(string(file))
	data = strings.TrimSpace(data)
	data = strings.ToLower(data)
	data = strings.ReplaceAll(data, "\n", " ")
	fmt.Printf("data: %v\n", data)
	strSlice := strings.Fields(data)
	resMap := make(map[string]int)
	for i := 0; i < len(strSlice); i++ {
		resMap[strSlice[i]] = strings.Count(data, strings.ToLower(strSlice[i]))
	}
	fmt.Println(resMap)
	func(map[string]int) {
		file, err := os.Create("output.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.Write([]string{"word", "count"})
		for k, v := range resMap {
			write := []string{k, strconv.Itoa(v)}
			fmt.Println("writing", write)
			writer.Write(write)
		}
	}(resMap)

}
func task2() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()

	info, err := os.Stat("app.log")
	fmt.Printf("Name:%s\nSize:%v\nisDir:%v\n", info.Name(), info.Size(), info.IsDir())
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	log.SetFlags(log.LstdFlags)
	log.SetOutput(logFile)
	log.Println("started logging inputs")
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			log.Println("exiting log program")
			break
		} else {
			log.Println(scanner.Text())
			fmt.Println(scanner.Text())
		}
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var digits = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func task3() {
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	users := make([]User, 0)
	json.Unmarshal(data, &users)
	file := excelize.NewFile()
	defer file.Close()
	file.SetCellValue("Sheet1", "A1", "Name")
	file.SetCellValue("Sheet1", "B1", "Age")
	for i := 0; i < len(users); i++ {
		for j := 0; j < 2; j++ {
			switch j {
			case 0:
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", letters[j], digits[i+1]), users[i].Name)
			case 1:
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", letters[j], digits[i+1]), users[i].Age)
			}

		}

	}
	file.SaveAs("report.xlsx")
}

func main() {
	// task1()
	// task2()
	task3()
}
