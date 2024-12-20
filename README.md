## Advent Of Code 2024

> This year I thought I'd give it another ***go*** 


### TOC

* [Installation](#installation)
* [Usage](#usage)
    * [Puzzle Input](#puzzle-input)
    * [Flags](#flags)
* [Testing](#testing)
* [Building](#building)
* [Results](#results)

### Installation
- Install [`go`](https://go.dev/doc/install)

- Clone the repo && cd into the directory

     ```
     git clone https://github.com/justjcurtis/2024
     cd AdventOfCode2024
     ```

- Run the project

    ```
    go run .
    ```

### Usage

#### Puzzle Input
The files in `./puzzleInput/` are read in and passed to each solution via `./main.go`. If you want to replace those files with your own input data to ensure correct solutions / compare runtimes just replace the file for the corresponding day in the `./puzzleInput/` dir & follow the naming convention in there (`day_{n}.txt`).

#### Flags

| Flag | Description | Example |
| ---- | ----------- | ------- |
| `-n` | How many times to run each solution. | `go run . -n 1000` Run each solution 1000 times and report the average runtime for each solution + the total average runtime. |
| `-min` | Report the minimum time instead of the average. | `go run . -min` Set `-n` to 5000 by default & report the minimum time for each day and the total minimum runtime. |
| `-d` | Only run a single day | `go run . -d 8` Only run day 8. This will only run the solution once unless other flags are set. |

### Testing
Each day is unit tested using the example input from the puzzle fpr that day on adventofcode.com

- Run the unit tests with go:

    ```
    go test ./... -v
    ```

### Building

To build locally:
- Follow the [installation instructions](#installation)
- Then run

    ```
    go build .
    ```
- A new file will be create in the root dir (`AdventOfCode2024`)
- Run the build with

    ```
    ./AdventOfCode2024 
    ```

### Results
Results show are the min runtime for each soltuion taken over 100 runs as reported by the github actions runner. Reading the input data from disk is not included as part of the solution here so the runtime you see is the runtime of any parsing & logic requried to solve the puzzle.
>*Your results may vary*

| # | Runtime (both parts) |
| - | -------------------- |
| Day 1 | 193µs |
| Day 2 | 226µs |
| Day 3 | 581µs |
| Day 4 | 587µs |
| Day 5 | 509µs |
| Day 6 | 223ms |
| Day 7 | 19ms |
| Day 8 | 112µs |
| Day 9 | 1ms |
| Day 10 | 555µs |
| Day 11 | 292ns |
| Day 12 | 4ms |
| Day 13 | 109µs |
| Day 14 | 1s |
| Day 15 | 39ms |
| ------- | ----------------------------- |
| **Total** | **1s** |


##### [Take Me To The TOP!](#top)
