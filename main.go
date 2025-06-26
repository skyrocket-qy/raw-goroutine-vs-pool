package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"time"

	"testing"

	"github.com/panjf2000/ants/v2"
)

const poolSize = 100
const maxGoChanCnt = 100

var sizes = []int{50000}

func ioTask() {
	ioBoundTask(30 * time.Millisecond)
}

func benchmarkRaw(taskCount int) (nsPerOp int64, bytesPerOp uint64, allocsPerOp float64) {
	done := make(chan struct{}, taskCount)

	maxGorChan := make(chan struct{}, maxGoChanCnt)
	f := func() {
		for i := 0; i < taskCount; i++ {
			maxGorChan <- struct{}{}
			go func() {
				ioTask()
				done <- struct{}{}
				<-maxGorChan
			}()
		}
		for i := 0; i < taskCount; i++ {
			<-done
		}
	}

	allocs := testing.AllocsPerRun(1, f)

	var mStart, mEnd runtime.MemStats
	runtime.ReadMemStats(&mStart)
	start := time.Now()
	f()
	duration := time.Since(start)
	runtime.ReadMemStats(&mEnd)

	memUsed := mEnd.TotalAlloc - mStart.TotalAlloc
	nsPerOp = duration.Nanoseconds() / int64(taskCount)
	bytesPerOp = memUsed / uint64(taskCount)

	return nsPerOp, bytesPerOp, allocs
}

func benchmarkPool(taskCount, poolSize int) (nsPerOp int64, bytesPerOp uint64, allocsPerOp float64) {
	p, _ := ants.NewPool(poolSize)
	defer p.Release()

	done := make(chan struct{}, taskCount)
	maxGorChan := make(chan struct{}, maxGoChanCnt)

	f := func() {
		for i := 0; i < taskCount; i++ {
			maxGorChan <- struct{}{}
			_ = p.Submit(func() {
				ioTask()
				done <- struct{}{}
				<-maxGorChan
			})
		}
		for i := 0; i < taskCount; i++ {
			<-done
		}
	}

	allocs := testing.AllocsPerRun(1, f)

	var mStart, mEnd runtime.MemStats
	runtime.ReadMemStats(&mStart)
	start := time.Now()
	f()
	duration := time.Since(start)
	runtime.ReadMemStats(&mEnd)

	memUsed := mEnd.TotalAlloc - mStart.TotalAlloc
	nsPerOp = duration.Nanoseconds() / int64(taskCount)
	bytesPerOp = memUsed / uint64(taskCount)

	return nsPerOp, bytesPerOp, allocs
}

func main() {
	f, _ := os.Create("detailed_results.csv")
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{
		"taskCount",
		"raw_ns/op", "raw_B/op", "raw_allocs/op", "raw_total_ns", "raw_total_allocs",
		"pool_ns/op", "pool_B/op", "pool_allocs/op", "pool_total_ns", "pool_total_allocs",
	})

	for _, count := range sizes {
		rawNs, rawB, rawA := benchmarkRaw(count)
		poolNs, poolB, poolA := benchmarkPool(count, poolSize)

		rawTotalNs := rawNs * int64(count)
		poolTotalNs := poolNs * int64(count)

		rawTotalAllocs := rawA * float64(count)
		poolTotalAllocs := poolA * float64(count)

		fmt.Printf("TaskCount: %d | Raw: %dns/op %dB/op %.2fallocs/op | Pool: %dns/op %dB/op %.2fallocs/op\n",
			count, rawNs, rawB, rawA, poolNs, poolB, poolA)

		w.Write([]string{
			fmt.Sprintf("%d", count),
			fmt.Sprintf("%d", rawNs),
			fmt.Sprintf("%d", rawB),
			fmt.Sprintf("%.2f", rawA),
			fmt.Sprintf("%d", rawTotalNs),
			fmt.Sprintf("%.0f", rawTotalAllocs),
			fmt.Sprintf("%d", poolNs),
			fmt.Sprintf("%d", poolB),
			fmt.Sprintf("%.2f", poolA),
			fmt.Sprintf("%d", poolTotalNs),
			fmt.Sprintf("%.0f", poolTotalAllocs),
		})
	}
}
