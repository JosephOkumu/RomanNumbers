package main

import (
    "fmt"
    "os"
    "strconv"
)

// Main function
func main() {
    // Check if the correct number of command-line arguments is provided
    if len(os.Args) != 2 {
        return
    }

    // Get the input from the command-line argument
    input := os.Args[1]

    // Parse the input number from the command-line argument
    number, err := strconv.Atoi(input)
    if err != nil || number <= 0 || number >= 4000 {
        fmt.Println("ERROR: cannot convert to roman digit")
        return
    }

    // Convert the number to a Roman numeral and print the breakdown and the Roman numeral
    roman, breakdown := toRoman(number)
    fmt.Println(breakdown)
    fmt.Println(roman)
}

// Function to convert an integer to a Roman numeral
func toRoman(num int) (string, string) {
    // Define the values and symbols for the Roman numerals
    vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    breakdowns := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

    // Initialize the Roman numeral and breakdown strings
    var roman string
    var breakdown string

    // Loop until the number is 0
    for i := 0; i < len(vals); i++ {
        // Loop until the number is greater than or equal to the current value
        for num >= vals[i] {
            // Subtract the value from the number and add the corresponding symbol to the Roman numeral
            num -= vals[i]
            roman += symbs[i]
            if breakdown != "" {
                breakdown += "+"
            }
            breakdown += breakdowns[i]
        }
    }

    // Return the Roman numeral and breakdown
    return roman, breakdown
}
