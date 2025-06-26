package main

import "time"

func ioBoundTask(t time.Duration) {
	time.Sleep(t)
}

func cpuBoundTask(t time.Duration) {
	sum := 0
	i := 1
	curT := time.Now()

	for time.Since(curT) < t {
		sum += i
		i++
	}
}

func memoryHeavyTask(t time.Duration, do func(t time.Duration)) {
	_ = make([]byte, 1024*1024) // 1 MB
	do(t)
}

func getTaskCnts() []int {
	return []int{
		10, 50, 100, 200, 500,
		1000, 2000, 5000,
		10000, 20000, 50000,
		100000, 200000, 300000,
		500000, 750000, 1000000,
	}
}

func getPoolSizes() []int {
	return []int{
		10, 50, 100, 200, 500,
		1000, 2000, 5000,
		10000, 20000, 50000,
		100000, 200000, 300000,
		500000, 750000, 1000000,
	}
}

func getTaskDuations() []int {
	return []int{
		1, 5, 10, 50, 100, 500, 1000, 5000, 10000, 50000, 100000,
	}
}
