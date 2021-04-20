package main

import (
	"math/rand"
	"sync"
	"time"
	"fmt"
)




type Job struct {
	id int
	randomno int
}

type Result struct {
	job Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 把number的每一位相加
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 创作工程协程的函数
func worker(wg * sync.WaitGroup) {
	// 计算缓冲信道jobs里面的每一个job，结果传入results信道
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	// 完成一次工作，wg减一
	// wg.Wait()会等待所有go协程执行完毕，即计数器为0
	wg.Done()
}

// 创建Go协程的工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)		// WaitGroup计数器+1
		go worker(&wg)  // 此处传入指针
	}
	wg.Wait()			// 等待计数器为0，即所有任务完成
	close(results)		// 所有协程执行完毕，关闭results信道
}

// 分配job给worker
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999) // 最大值为998的随机数
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)	// 已写入所有job，关闭jobs信道
}

// 读取results信道和打印输出
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	// 把任务内容写进信道jobs
	go allocate(noOfJobs)
	done := make(chan bool)
	// 打印每一个任务执行后的结果，直到所有结果都被打印完。不能放在createWorkerPool后面执行，因为该函数执行完毕后close(results)，就无法再打印results了
	go result(done)
	// 确定worker数量，创建工作池，分配任务给每一个worker
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	// 判断任务结果打印完成
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time token", diff.Seconds(), "seconds")
}




