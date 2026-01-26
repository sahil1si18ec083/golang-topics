package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	Id    string
	Name  string
	Email string

	Age int
}

func Process(jobchan <-chan Job, resultchan chan<- Job, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobchan {
		resultchan <- job
	}

}

func main() {
	fmt.Println("hello")
	log.Print("hi")

	start := time.Now()

	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	recordcount := len(records[1:])

	workerscount := 50

	jobchan := make(chan Job, recordcount)
	resultchan := make(chan Job, recordcount)
	var wg sync.WaitGroup
	wg.Add(workerscount)

	for i := 0; i < workerscount; i++ {

		go Process(jobchan, resultchan, i, &wg)

	}

	for _, record := range records[1:] {
		age, err := strconv.Atoi(record[4])
		if err != nil {
			continue
		}

		job := Job{
			Id:    record[0],
			Name:  record[2],
			Email: record[3],
			Age:   age,
		}
		jobchan <- job

	}
	close(jobchan)

	wg.Wait()

	// for val := range resultchan {
	// 	// fmt.Print(val.Email)
	// }
	close(resultchan)

	fmt.Print(time.Since(start))
}
