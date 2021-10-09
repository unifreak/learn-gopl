// Echo1 prints its command-line arguments.
package main

import ( // import mutiple packages with list form
    "fmt"
    "os"
)

func main() {
    var s, sep string
    // := is part of short variable declaration:  declares
    //    one or more variables and gives them appropriate types
    //    based on the initializer values
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    // for is the only loop in Go
    //
    // form 1
    //  	for init; condition; post {
    //  	}
    // form 2: while loop
    // 		for condition {
    // 		}
    // form 3: infinite loop
    // 		for {
    // 		}
    fmt.Println(s)
}

// i++ is statement, not expression. hence j = i++ is illegal
// ++/-- are postfix only, so --i is illegal, too

