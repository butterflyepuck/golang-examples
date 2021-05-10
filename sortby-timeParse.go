package main

import (
  "time"
  "fmt"
  "sort"
)

func main() {
  layout := "2006-01-02 15:04"
  c,_ := time.Parse(layout, "2021-05-10 16:29")
  d,_ := time.Parse(layout, "2021-05-10 16:31")

  x := []time.Time{d, c}

  fmt.Println(x)

  sort.Slice(x, func(i,j int) bool {
	return x[i].Before(x[j])
  })

  fmt.Println(x)
}
