# Go Goroutine Benchmark: Raw vs. Pool

This project provides a benchmarking tool to compare the performance of creating raw goroutines versus using a goroutine pool (`ants`) in Go. It is designed to demonstrate the performance benefits of using a goroutine pool, especially for managing a large number of I/O-bound or CPU-bound tasks concurrently.

## Overview

In Go, goroutines are a lightweight and efficient way to handle concurrency. However, creating a very large number of goroutines can still lead to high memory consumption and performance overhead. A goroutine pool can help mitigate these issues by reusing a fixed number of goroutines to execute tasks.

This benchmark tool runs a series of tests to compare these two approaches. It measures key performance metrics, including:

- **Nanoseconds per operation (ns/op):** The average time taken to complete a single task.
- **Bytes per operation (B/op):** The average memory allocated per task.
- **Allocations per operation (allocs/op):** The average number of memory allocations per task.

The results are printed to the console and saved in a `detailed_results.csv` file for more detailed analysis.

## Features

- **Benchmark Raw Goroutines:** Measures the performance of creating a new goroutine for each task.
- **Benchmark Goroutine Pool:** Measures the performance of using the `ants` goroutine pool to manage tasks.
- **Configurable Tasks:** Includes different types of tasks for benchmarking:
  - `ioBoundTask`: Simulates tasks that wait for I/O operations (e.g., network requests, file I/O).
  - `cpuBoundTask`: Simulates tasks that are computationally intensive.
  - `memoryHeavyTask`: Simulates tasks that allocate a significant amount of memory.
- **Customizable Benchmarks:** Easily configure benchmark parameters such as the number of tasks, pool size, and task duration.
- **Detailed Results:** Generates a CSV file with detailed performance metrics for each benchmark run.

## How to Run

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the benchmarks:**
   ```bash
   go run .
   ```

The benchmark results will be printed to the console, and a `detailed_results.csv` file will be created in the project's root directory.

## Benchmark Results

The output provides a side-by-side comparison of the raw goroutine and goroutine pool approaches for different numbers of tasks.

### Console Output

The console output will look something like this:

```
TaskCount: 50000 | Raw: 60473ns/op 25B/op 0.00allocs/op | Pool: 31139ns/op 50B/op 0.00allocs/op
```

- **TaskCount:** The number of tasks executed in the benchmark.
- **Raw:** The performance metrics for creating raw goroutines.
- **Pool:** The performance metrics for using the `ants` goroutine pool.

### `detailed_results.csv`

The `detailed_results.csv` file contains a more detailed breakdown of the results, including total execution time and total memory allocations. This file can be imported into a spreadsheet program for further analysis and visualization.

| taskCount | raw_ns/op | raw_B/op | ... | pool_ns/op | pool_B/op | ... |
|-----------|-----------|----------|-----|------------|-----------|-----|
| 50000     | 60473     | 25       | ... | 31139      | 50        | ... |


## Customization

You can customize the benchmarks by modifying the following files:

- **`main.go`:**
  - `sizes`: Change the values in this slice to run benchmarks with different numbers of tasks.
  - `poolSize`: Adjust the size of the `ants` goroutine pool.
  - `ioTask()`: Modify this function to switch between different task types (e.g., `cpuBoundTask`, `memoryHeavyTask`) or to change the task duration.

- **`task.go`:**
  - `getTaskCnts()`, `getPoolSizes()`, `getTaskDuations()`: These functions provide lists of values that can be used for more extensive benchmark sweeps. You can modify these or use them in `main.go` to run a wider range of tests.
  - `ioBoundTask()`, `cpuBoundTask()`, `memoryHeavyTask()`: Modify the implementation of these tasks to better simulate your specific use case.