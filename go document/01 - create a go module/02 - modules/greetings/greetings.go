package greetings

import "fmt"

func Hello(name string) string {
  
  // var message string
  // message = fmt.Sprintf("G' Morning, %v!",name)
  
  message := fmt.Sprintf("G' Morning, %v!", name)    
  return message
}