/*
Title: 54.SHA1_Hashes.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

SHA1 hashes are frequently used to compute short identities for binary or text blobs.
For example, the git revision control system uses SHA1s extensively to identify versioned files 
and directories. Here’s how to compute SHA1 hashes in Go.

You can compute other hashes using a similar pattern to the one shown above. 
For example, to compute MD5 hashes import crypto/md5 and use md5.New().

Note that if you need cryptographically secure hashes, 
you should carefully research hash strength!

*/

package main 

// Go implements several hash functions in various crypto/* packages.
import "crypto/sha1"
import "fmt"


func main(){
    fmt.Println("54.SHA1_Hashes.go---------Start------------\n\n")

    s := "sha1 this string"
    // The pattern for generating a hash is sha1.New(), sha1.Write(bytes), then sha1.Sum([]byte{}).
    // Here we start with a new hash.
    h := sha1.New()
    fmt.Println(s)
    fmt.Println(h)

    // Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
    h.Write([]byte(s))
    
    // This gets the finalized hash result as a byte slice. 
    // The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
    bs := h.Sum(nil)

    // SHA1 values are often printed in hex, for example in git commits. 
    // Use the %x format verb to convert a hash results to a hex string.
    fmt.Println(s)
    fmt.Printf("%x \n", bs)

    fmt.Println("\n\n54.SHA1_Hashes.go-----------End------------")
    // Running the program computes the hash and prints it in a human-readable hex format.
    
}

