package main

import "fmt"

func main() {
    size := 15 
    sandglass(size)
}

func sandglass(size int) {
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if i == 0 || i == size-1 || i == j || i+j == size-1     {
                fmt.Print("\033[33mX")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
