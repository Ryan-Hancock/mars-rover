# Mars Rover Mission Command


## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
    * [Tests](#tests)
    * [Inputs](#inputs)

## General info
This is a Golang (Go) implementation to operate a squad of rovers on a plateau on Mars. 

## Technologies
These are the technologies used to create this project:
* Go 1.14.7
* [Testify]: http://github.com/stretchr/testify  

## Setup
To run this project, build it locally using Go:

```
$ go build -o mars-rover 

```

Once built you can feed it files via the CLI as the second argument:
```
$ ./mars-rover input.txt
```

An example test file called input.txt is included and follows the format of: 
```
5 5 -> plateau (x,y)
1 2 N -> the starting location of the rover and its facing direction (x, y, "N")
LMLMLMLMM -> commands for the rover | L left, R right, M move
```

Additional rovers can be added to the file via the input with additional starting locations and commands on subsequent lines: 
```
3 3 E
MMRMMRMRRM
```
### Tests
In order to run tests:  
```
$ go test ./... -v
```

### Inputs

The input file can be changed to .csv or another seperator file with new lines by changing the `main.go` file.

```
21 | lines, err := fileio.ParseFile(r, rune(' '))
```


