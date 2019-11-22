package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  fmt.Println("Hi there!")
  fmt.Println("Welcome to the To-Do-List manager.")
  fmt.Println("Would you like to enter an Event or a Task?")
  reader := bufio.NewReader(os.Stdin)
  var userInput string
  userInput, _ = reader.ReadString('\n')
  toDo := strings.ToLower(strings.TrimSpace(userInput))
  if toDo == "event" {
    fmt.Print("You have chosen to add event\n")
  } else {
    fmt.Print("You have chosen to add something\n")
  }
}