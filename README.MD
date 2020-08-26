# Mars Rover Challenge
Go implementation to operate a squad of mars rovers on a plateau.

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Customise](#customise)

## General info
This project has test files coverage which can be used to better understand how to interact with the packages.
To run tests  
```
$ go test ./... -v
```

## Technologies
Project is created with:
* Go 1.14.7
* [Testify]: http://github.com/stretchr/testify  

## Setup
To run this project, build it locally using go:

```
$ go build -o mars-rover 

```

Once built you can feed it files via the cli as the second arg:
```
$ ./mars-rover input.txt
```

A example test file called input.txt is included. And follows the format of 
```
5 5 -> plataue (x,y)
1 2 N -> rover starting location and way facing (x, y, "N")
LMLMLMLMM -> commands for the rover | L left, R right, M move
```

Extra rover can be added to the input via to the input with just the rover starting and commands. 
```
3 3 E
MMRMMRMRRM
```
## Customise

The input file can be changed to csv or another seperator file with new lines by changing the main file.

```
21 | lines, err := fileio.ParseFile(r, rune(' '))
```


