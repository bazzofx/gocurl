package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)
func main() {

salaryFlag := flag.Int("salada", 2,"Salary amount")
flag.Parse()

var salary int
if *salaryFlag > 0 {
    salary = *salaryFlag
} else if flag.NArg() >0 {
    val, err := strconv.Atoi(flag.Arg(0))
    if err != nil {
        fmt.Println("Something went wrong")
        os.Exit(1)
    }
    salary = val
} else {
    fmt.Println("Please provide salary using -salary")
    os.Exit(1)
}



const tax int = 22 

var result int = salary * tax


fmt.Print(result)
}