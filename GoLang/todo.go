package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "time"
)

var reader *bufio.Reader
var dateFormat string = "02/01/2006"
var timeFormat string = "15:04"

type Task struct {
	date time.Time
	startTime time.Time
	duration time.Time
	assigned []string
}

type Event struct {
	date time.Time
	startTime time.Time
	location string
}

type ToDo struct {
	task *Task
	event *Event
}

func getDate() time.Time {
	// prompt user for date in dd/mm/yyyy format
	fmt.Println("Please enter a date: dd/mm/yyyy")
	dateInput, _ := reader.ReadString('\n')
	date, _ := time.Parse(dateFormat, strings.TrimSpace(dateInput))
	// date := strings.Split(strings.TrimSpace(dateInput),"/")
	return date
}

func getStartTime() time.Time {
	// prompt user for date in hh:mm format
	fmt.Println("Please enter a start time: hh:mm")
	startTimeInput, _ := reader.ReadString('\n')
	startTime, _ := time.Parse(timeFormat, strings.TrimSpace(startTimeInput))
	// startTime := strings.Split(strings.TrimSpace(startTimeInput),":")
	return startTime
}

func getLocation() string {
	// prompt user for location
	fmt.Println("Please enter a location")
	locationInput, _ := reader.ReadString('\n')
	location := strings.TrimSpace(locationInput)
	return location
}

func getDuration() time.Time {
	// prompt user for duration in hh:mm format
	fmt.Println("Please enter a the duration: hh:mm")
	// var durationInput float64
	// _, err := fmt.Scanf("%f", &durationInput)
	durationInput, _ := reader.ReadString('\n')
	duration, _ := time.Parse(timeFormat, strings.TrimSpace(durationInput))
	// fmt.Println(err)
	return duration
}

func getAssigned() []string {
	// prompt user for list of people seperated by comma
	fmt.Println("Please Enter the people assigned to this task, seperated my a comma ','")
	assignedInput, _ := reader.ReadString('\n')
	assigned := strings.Split(strings.TrimSpace(assignedInput),",")
	return assigned
}

func event() Event {
	fmt.Print("You have chosen to add event\n")
	date := getDate()
	startTime := getStartTime()
	location := getLocation()
	fmt.Println("The date is", date.Format(dateFormat))
	fmt.Println("The start time is", startTime.Format(timeFormat))
	fmt.Println("The location is", location)
	eventEntry := Event{
		date:date,
		startTime:startTime,
		location:location,
	}
	return eventEntry
}

func task() Task {
	fmt.Print("You have chosen to add a task\n")
	date := getDate()
	startTime := getStartTime()
	duration := getDuration()
	assigned := getAssigned()
	fmt.Println("The date is", date.Format(dateFormat))
	fmt.Println("The start time is", startTime.Format(timeFormat))
	fmt.Println("The duration is", duration.Format(timeFormat))
	fmt.Println("The people assigned to the task are", assigned)
	taskEnter := Task{
		date: date,
		startTime: startTime,
		duration: duration,
		assigned: assigned,
	}
	return taskEnter
}


var myQueue []ToDo
func getToDo() []ToDo {
	if len(myQueue) == 0 {
		fmt.Println("The To-Do list is currently empty")
		return myQueue
	}
	firstToDo, myQueue := myQueue[0], myQueue[1:]
	if firstToDo.event != nil {
		fmt.Println("Event: on", firstToDo.event.date.Format(dateFormat), "at", firstToDo.event.startTime.Format(timeFormat), "location", firstToDo.event.location)
	}
	if firstToDo.task != nil {
		fmt.Println("Task: ", firstToDo.task.date.Format(dateFormat), "at", firstToDo.task.startTime.Format(timeFormat), "duration", firstToDo.task.duration, "assigned to", firstToDo.task.assigned)
	}
	return myQueue
}

func viewToDo() {
	if len(myQueue) == 0 {
		fmt.Println("The To-Do list is currently empty")
	}
	for i, todo := range myQueue {
		if todo.event != nil {
			fmt.Println(i+1,": Event: on", todo.event.date.Format(dateFormat), "at", todo.event.startTime.Format(timeFormat), "location", todo.event.location)
		}
		if todo.task != nil {
			fmt.Println(i+1, ": Task: ", todo.task.date.Format(dateFormat), "at", todo.task.startTime.Format(timeFormat), "duration", todo.task.duration, "assigned to", todo.task.assigned)
		}
	}
}

func main() {
	fmt.Println("Hi there!")
	fmt.Println("Welcome to the To-Do-List manager.")
	fmt.Println("To exit enter 'quit'.")
	fmt.Println("If you would like to check your first ToDo, enter 'next to-do'")
	fmt.Println("If you would like to check your first ToDo, enter 'view to-do'")
	reader = bufio.NewReader(os.Stdin)
	state := true
	for state {
		fmt.Println("Would you like to enter an Event or a Task?")
		userInput, _ := reader.ReadString('\n')
		toDo := strings.ToLower(strings.TrimSpace(userInput))
		if toDo == "event" {
			eventEntry := event()
			myQueue = append(myQueue, ToDo{
				event: &eventEntry,
				task: nil,
			})
		} else if toDo == "task" {
			taskEntry := task()
			myQueue = append(myQueue, ToDo{
				task: &taskEntry,
				event: nil,
			})
		} else if toDo == "quit" {
			state = false
		}else if toDo == "next to-do" {
			myQueue = getToDo()
		} else if toDo == "view to-do"{
			viewToDo()
		} else {
			fmt.Println("What you have entered is invalid")
		}
	}
	fmt.Println("Thank you for using the To-Do-List manager")
}