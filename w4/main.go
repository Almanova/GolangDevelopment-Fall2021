package main

import (
	"fmt"
	"flag"
	"time"
	"io/ioutil"
	"encoding/json"
)

type Task struct {
	Name string
	StartTime string
	EndTime string
	Duration string
}

func (t Task) String() string {
	return fmt.Sprintf("Task '%s' was started at %s and finished at %s. It took %s to get it done.", 
		t.Name, t.StartTime, t.EndTime, t.Duration)
}

var task Task
var operationType string

func writeTask() {
	taskJson, err := json.Marshal(task)
	err = ioutil.WriteFile("task.json", taskJson, 0644)
    if err != nil {
        panic(err)
    }
}

func readTask() {
	taskJson, error := ioutil.ReadFile("task.json")
    if error != nil {
        panic(error)
    }
	error = json.Unmarshal(taskJson, &task)
	if error != nil {
        panic(error)
    }

	task.EndTime = time.Now().Format("2006-01-02 15:04:05")
	endTime, err := time.Parse("2006-01-02 15:04:05", task.EndTime)
	startTime, err := time.Parse("2006-01-02 15:04:05", task.StartTime)
	if err != nil {
		panic(err)
	}

	task.Duration = endTime.Sub(startTime).String()
}

func init() {
	flag.StringVar(&operationType, "operationType", "finish", "flag argument")
	flag.StringVar(&task.Name, "name", "", "first argument")
	task.StartTime = time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	flag.Parse()
	if operationType == "start" {
		writeTask()
		fmt.Println("Good Luck!")
	} else {
		readTask()
		fmt.Println(task)
	}
}
