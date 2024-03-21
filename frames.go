package main

import (
    "fmt"
    "image"
    _"image/png"
    _"image/jpeg"
    "os"
    "strconv"
    "time"
)

func printFrame(img image.Image, condenseFactor uint32) {
    // getting the bounds for the image
    bounds := img.Bounds()
    // Iterating through each pixel
    for y := bounds.Min.Y; y < bounds.Max.Y; y+= int(condenseFactor) {
        line := ""
        for x := bounds.Min.X; x < bounds.Max.X; x+= int(condenseFactor) {
            var sum uint32 = 0
            for i := 0; i < int(condenseFactor); i++ {
                pixel := img.At(x + i, y + i)
                r, g, b, _ := pixel.RGBA()
                r = r >> 8
                g = g >> 8
                b = b >> 8
                sum += r + g + b
            }
            if sum > 700*condenseFactor {
               line += "@"
            } else if sum > 600*condenseFactor {
               line += "&"
            } else if sum > 550*condenseFactor {
               line += "%"
            } else if sum > 500*condenseFactor {
               line += "#"
            } else if sum > 400*condenseFactor {
               line += "9"
            } else if sum > 250*condenseFactor {
               line += "{"
            } else if sum > 100*condenseFactor {
               line += "*"
            } else if sum > 50*condenseFactor {
               line += "^"
            } else {
               line += "`"
            }
        }
        fmt.Println(line)
    }
}

func main() {
    // Checking a command-line argument was passed
    if len(os.Args) < 2 {
        fmt.Println("Please input file name")
        os.Exit(1)
    }

    // Reading the file name
    fileName := os.Args[1]
    i := 1
    for true {
        // Opening the files
        strLen := len(fileName)
        file, err := os.Open(fileName[:strLen-4] + " (" + fmt.Sprint(i) + ").jpg")
        if err != nil {
            // fmt.Println("Error opening frames:", err)
          //  fmt.Println(i)
            os.Exit(1)
        }

        inputNumber := 5
        if len(os.Args) > 2 {
            inputNumber, err = strconv.Atoi(os.Args[2])
            if err != nil {
                fmt.Println("Error reading condense factor:", err)
                os.Exit(1)
            }
        }
        condenseFactor := uint32(inputNumber)
        
        // Decoding the file
        img, _, err := image.Decode(file)
        if err != nil {
            fmt.Println("Error decoding file:", err)
            os.Exit(1)
        }

        file.Close()

        printFrame(img, condenseFactor)
        time.Sleep(30 * time.Millisecond)
        fmt.Printf("\033[0;0H")
        i++
    }
}   


