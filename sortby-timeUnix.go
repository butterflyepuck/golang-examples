package main                                                                   
import (                                                                       
  "time"                                                                       
  "fmt"                                                                        
  "sort"
  )                                                                              
                                                                               
func main() {                                                                  
  a := time.Now().Unix()                                                       
  time.Sleep(time.Duration(3)*time.Second)                                     
  b := time.Now().Unix()                                                       
                                                                               
  c := time.Unix(a,0)                                                          
  d := time.Unix(b,0)                                                          
                                                                               
  x := []time.Time{d, c}                                                       
                                                                               
  fmt.Println(x)                                                               
                                                                               
  sort.Slice(x, func(i,j int) bool {                                           
    return x[i].Before(x[j])                                                   
  })                                                                           
                                                                               
  fmt.Println(x)                                                               
}
