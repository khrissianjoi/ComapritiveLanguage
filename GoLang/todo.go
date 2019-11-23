package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func getDate() []string {
	fmt.Println("Please enter a date: dd/mm/yyyy")
	reader := bufio.NewReader(os.Stdin)
	var dateInput string
	dateInput, _ = reader.ReadString('\n')
	date := strings.Split(strings.TrimSpace(dateInput),"/")
	return date
}

func getStartTime() []string {
	fmt.Println("Please enter a start time: hh:mm")
	reader := bufio.NewReader(os.Stdin)
	var startTimeInput string
	startTimeInput, _ = reader.ReadString('\n')
	startTime := strings.Split(strings.TrimSpace(startTimeInput),":")
	return startTime
}

func getLocation() string {
	fmt.Println("Please enter a location")
	reader := bufio.NewReader(os.Stdin)
	var locationInput string
	locationInput, _ = reader.ReadString('\n')
	location := strings.TrimSpace(locationInput)
	return location
}

func getDuration() float64 {
	fmt.Println("Please enter a the duration: hh.mm")
	// reader := bufio.NewReader(os.Stdin)
	var durationInput float64
	_, err := fmt.Scanf("%f", &durationInput)
	fmt.Println(err)
	return durationInput

}
func getAssigned() []string {
	fmt.Println("Please Enter the people assigned to this task, seperated my a comma ','")
	var assignedInput string
	reader := bufio.NewReader(os.Stdin)
	assignedInput, _ = reader.ReadString('\n')
	assigned := strings.Split(strings.TrimSpace(assignedInput),",")
	return assigned

}
func event() {
	fmt.Print("You have chosen to add event\n")
	date := getDate()
	startTime := getStartTime()
	location := getLocation()
	fmt.Println("The date is", date)
	fmt.Println("The start time is", startTime)
	fmt.Println("The location is", location)
}

func task() {
	fmt.Print("You have chosen to add a task\n")
	date := getDate()
	startTime := getStartTime()
	duration := getDuration()
	assigned := getAssigned()
	fmt.Println("The date is", date)
	fmt.Println("The start time is", startTime)
	fmt.Println("The duration is", duration)
	fmt.Println("The people assigned to the task are", assigned)

}

func main() {
  fmt.Println("Hi there!")
  fmt.Println("Welcome to the To-Do-List manager.")
  fmt.Println("Would you like to enter an Event or a Task?")
  reader := bufio.NewReader(os.Stdin)
  var userInput string
  userInput, _ = reader.ReadString('\n')
  toDo := strings.ToLower(strings.TrimSpace(userInput))
  if toDo == "event" {
	event()
  } else {
	task()
  }
  fmt.Print("DONE")
}