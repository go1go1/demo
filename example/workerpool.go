package main

import (
	"fmt"
	"math/rand"
)

/**
WorkerPool的实现

计算一个数字的各个位数之和，比如123， 和等于 1+2+3=6
需要计算的数字使用随机生成
*/

type Job struct {
	Number int
	Id     int
}

type Result struct {
	job    *Job
	result int
}

func calc(job *Job, result chan *Result) {
	var sum int
	number := job.Number
	for number != 0 {
		tmp := number % 10
		sum += tmp
		number /= 10
	}

	r := &Result{
		job:    job,
		result: sum,
	}

	result <- r
}

func printResult(resChan chan *Result) {
	for res := range resChan {
		fmt.Printf("job id:%v number: %v, result: %d\n", res.job.Id, res.job.Number, res.result)
	}
}

func Worker(jobChan chan *Job, resChan chan *Result) {
	for job := range jobChan {
		calc(job, resChan)
	}
}

func startWorkerPool(num int, jobChan chan *Job, resChan chan *Result) {
	for i := 0; i < num; i++ {
		go Worker(jobChan, resChan)
	}
}

func main() {
	jobChan := make(chan *Job, 1000)
	resChan := make(chan *Result, 1000)

	startWorkerPool(128, jobChan, resChan)

	go printResult(resChan)

	var id int
	for {
		id++
		number := rand.Int()
		job := Job{
			Id:     id,
			Number: number,
		}
		jobChan <- &job
	}
}
