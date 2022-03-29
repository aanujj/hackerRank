package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "log"
)

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */
//10 5 20 20 4 5 2 25 1
func breakingRecords(scores []int32) []int32 {
    
    highScore := int32(0) 
    lowScore := int32(0)
    max := scores[0]
    min := scores[0]
    for _,value := range scores{
        
         if max < value{
            max = value 
            highScore+=1
        } else if min > value{
            min = value 
            lowScore+=1
        }
        
    }
    
    log.Println(highScore)
    log.Println(lowScore)
    
    // Write your code here
    return []int32{highScore,lowScore}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    scoresTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var scores []int32

    for i := 0; i < int(n); i++ {
        scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
        checkError(err)
        scoresItem := int32(scoresItemTemp)
        scores = append(scores, scoresItem)
    }

    result := breakingRecords(scores)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
