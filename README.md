# Advent of Code 2024

This repository includes my Go code for [Advent of Code 2024](https://adventofcode.com/2024).

## Running From Terminal

To run this code from the terminal, first clone this repository then move into the project's directory.

Add an input.txt in the day's directory.

Then run the following code in the terminal:

```
go run . [day] [part]
```

For example, if you want to run the code for day 1 part 2, you would run:

```
go run . 1 2
```

### Running Tests

To run all tests for all days and parts, add in test .txt files to each day's directory then run the following in the terminal:

```
go test ./... -v
```

To run tests for day 01 , run the following in the terminal (replace 01 with the day number you wish to run):

```
go test ./day01 -v
```
