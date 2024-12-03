# Advent of Code 2024

https://adventofcode.com/2024

## Running Locally

- Written in golang
- Run it on the root folder, eg

```
go run day1/main.go
```

## Scripts/Make commands

Makefile in root folder governs the targets/commands to use

1. For help, refer to help message in

```
make help
```

2. For creating skeleton files for problem of the day

- creates a directory called day2
- creates files main.go, example.txt and input.txt

```
make skeletion DIR=day2
```

3. For removing code in a directory

```
make clean DIR=day2
```
