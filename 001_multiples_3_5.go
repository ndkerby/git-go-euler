package main;

import (
    "fmt"
)

func main() {
    var max_number = 1000
    var numbers []int
    var sum = 0
    for i := 1; i < max_number; i++ {
        if i%3==0 {
            numbers = append(numbers,i)
        } else if i%5==0 {
            numbers = append(numbers,i)
        }
    }
    
    for index := range numbers {
        
        sum += numbers[index]
    }
    fmt.Println("The sum = ", sum)
}
