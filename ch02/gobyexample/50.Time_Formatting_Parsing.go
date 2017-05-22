/*
Title: 50.Time_Formatting_Parsing.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

Go supports time formatting and parsing via pattern-based layouts.

*/

package main 

import "fmt"
import "time"


func main() {

    fmt.Println("50.Time_Formatting_Parsing.go---------Start------------\n\n")

    p := fmt.Println
    t := time.Now()

    // Here’s a basic example of formatting a time according to RFC3339, 
    // using the corresponding layout constant.
    p("t.Format(time.RFC3339))", t.Format(time.RFC3339)) // 2017-05-23T05:58:32+09:00

    // Time parsing uses the same layout values as Format.
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1) // 2012-11-01 22:08:41 +0000 +0000

    // Format and Parse use example-based layouts. 
    // Usually you’ll use a constant from time for these layouts, 
    // but you can also supply custom layouts. 
    // Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 
    // to show the pattern with which to format/parse a given time/string. 
    // The example time must be exactly as shown: 
    // the year 2006, 15 for the hour, Monday for the day of the week, etc.
    p(t.Format("3.04PM")) // 5.58AM
    p(t.Format("Mon Jan _2 15:04:05 2006")) // Tue May 23 05:58:32 2017
    p(t.Format("2006-01-02T15:04:05.999999-07:00")) // 2017-05-23T05:58:32.278831+09:00
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)  // 0000-01-01 20:41:00 +0000 UTC

    // For purely numeric representations 
    // you can also use standard string formatting with the extracted components of the time value.
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:--\n",
        t.Year(), t.Month(), t.Day(), 
        t.Hour(), t.Minute(), t.Second()) //2017-05-23T05:58:32-00:--

    // Parse will return an error on malformed input explaining the parsing problem.
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    
    p(e) // parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"

    fmt.Println("\n\n50.Time_Formatting_Parsing.go-----------End------------")
}
