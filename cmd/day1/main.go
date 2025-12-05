package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    fmt.Println("password: ", processPassword("input"))
}

func processPassword(inputFile string) int {
    file, err := os.Open("input")
    if err != nil {
        panic(fmt.Sprintf("Error opening file: %v\n", err))
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    position := 50
    password := 0
    for scanner.Scan() {
        line := scanner.Text()
        
        direction := line[0]
        number, err := strconv.Atoi(line[1:])
        if err != nil {
            fmt.Printf("Error converting number: %v\n", err)
            continue
        }

         // if we get a number larger than 100, it's just doing
         //  a full rotation, so just reduce it to the relative movement
        number = number % 100

        switch direction {
        case 'L':
            fmt.Printf("Turn Left and move %d steps\n", number)
            position -= number
        case 'R':
            fmt.Printf("Turn Right and move %d steps\n", number)
            position += number
        default:
            fmt.Printf("Unknown direction: %c\n", direction)
        }

        if position > 99 {
            position = position - 100
        }

        if position < 0 {
            position = 100 + position
        }

        if position == 0 {
            password += 1
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error while reading file: %v\n", err)
    }
    return password
}