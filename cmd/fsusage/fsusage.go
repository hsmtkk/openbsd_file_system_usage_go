package main

import(
  "fmt"
  "log"
  "os"
  "strconv"
  "github.com/hsmtkk/openbsd_file_system_usage_go/pkg/fsusage"
)

func main(){
  if len(os.Args) != 2 {
    log.Fatalf("Usage: %s threshold", os.Args[0])
  }
  threshold, err := strconv.Atoi(os.Args[1])
  if err != nil {
    log.Fatalf("failed to parse threshold as int: %s", os.Args[1])
  }

  usages, err := fsusage.GetFsUsages()
  if err != nil {
    log.Fatal(err)
  }
  for _, usage := range usages {
    if usage.Capacity > threshold {
      fmt.Println(usage)
    }
  }
}
