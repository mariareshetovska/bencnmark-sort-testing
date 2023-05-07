# Sorting Algorithm Benchmarks
This repository contains benchmark tests for two sorting algorithms implemented in Golang: Bubble Sort, Quick Sort and sorting from the standard package "sort". The benchmarks are designed to compare the performance of the two algorithms on random integer arrays of various sizes.

## Running the Benchmarks
To run the benchmark tests, navigate to the root directory of the project and execute the following command:

```go
go test -bench=.
```
This will run the Bubble Sort, Quick Sort, and sorting from the standard package "sort" benchmarks and display the results in the terminal.

## Profiling the Benchmarks

We can use the pprof tool to profile the benchmarks and analyze the performance. There are two steps to get the profiling output:

First, run the benchmarks with the `-memprofile` and -cpuprofile flags to generate memory and CPU profiling data:
1. Create a binary having profiling information with the command:
```
go test -run=XXX -cpuprofile cpu.prof -bench .
```
Profiling is run on the benchmark we wrote before. Also, the profile file is named "cpu.prof".

2. Visualize the output using the command:
```
go tool pprof cpu.prof
```
This opens a terminal. Type "web" and press enter to open the web version of the profiling report. You can view various statistics and call graphs to identify performance bottlenecks and optimize your code.
