package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

    b, err := os.ReadFile("input")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return 
    }

    fmt.Println("password: ", processPassword(string(b)))
}

func processPassword(input string) int {
    lines := strings.Split(input, "\n")

    position := 50
    password := 0
    for _, line := range lines {
        if line == "" {
            continue
        }
        fmt.Println(line)
        direction := line[0]
        number, err := strconv.Atoi(line[1:])
        if err != nil {
            fmt.Printf("Error converting number: %v\n", err)
            continue
        }

        // if the number is larger than 100, we do a full rotation, and count how many rotaions to see how often we pass 0
        if number > 100 {
            fullRotation := number / 100
            password += fullRotation
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

        // at or passed 0, normalize position and increase password
        if position > 99 {
            position = position - 100
            password += 1 
        } else if position < 0 {
            position = 100 + position
            // crossing position 0 going left, 
            // but check that the previous position wasn't 0
            if position + number != 100 {
                password += 1
            }
        } else if position == 0 {
            password += 1
        }

        fmt.Println("New position: ", position, " password: ", password)
    }


    return password
}