package main

import "fmt"

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    // Maps for character mapping
    mappingST := make(map[byte]byte)
    mappingTS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        charS := s[i]
        charT := t[i]

        // Check if there's an existing mapping
        if mappedT, exists := mappingST[charS]; exists {
            if mappedT != charT {
                return false
            }
        } else {
            mappingST[charS] = charT
        }

        // Check reverse mapping
        if mappedS, exists := mappingTS[charT]; exists {
            if mappedS != charS {
                return false
            }
        } else {
            mappingTS[charT] = charS
        }
    }

    return true
}

func main() {
    s := "egg"
    t := "add"
    fmt.Println(isIsomorphic(s, t)) // Output: true

    s = "foo"
    t = "bar"
    fmt.Println(isIsomorphic(s, t)) // Output: false

    s = "paper"
    t = "title"
    fmt.Println(isIsomorphic(s, t)) // Output: true
}
