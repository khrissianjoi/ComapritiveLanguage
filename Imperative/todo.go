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
	duration string
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
	return date
}

func getStartTime() time.Time {
	// prompt user for date in hh:mm format
	fmt.Println("Please enter a start time: hh:mm")
	startTimeInput, _ := reader.ReadString('\n')
	startTime, _ := time.Parse(timeFormat, strings.TrimSpace(startTimeInput))
	return startTime
}

func getLocation() string {
	// prompt user for location
	fmt.Println("Please enter a location")
	locationInput, _ := reader.ReadString('\n')
	location := strings.TrimSpace(locationInput)
	return location
}

func getDuration() string {
	// prompt user for duration in hh:mm format
	fmt.Println("Please enter a the duration: hh.mm")
	durationInput, _ := reader.ReadString('\n')
	duration := strings.TrimSpace(durationInput)
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
	// gathers information an event requires
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
	fmt.Println("Event-\nDate:", eventEntry.date.Format(dateFormat), "\nStart Time:", eventEntry.startTime.Format(timeFormat), "\nLocation", eventEntry.location)
	fmt.Println("Event added")
	return eventEntry
}

func task() Task {
	// gathers information a task requires
	fmt.Print("You have chosen to add a task\n")
	date := getDate()
	startTime := getStartTime()
	duration := getDuration()
	assigned := getAssigned()
	fmt.Println("The date is", date.Format(dateFormat))
	fmt.Println("The start time is", startTime.Format(timeFormat))
	fmt.Println("The duration is", duration)
	fmt.Println("The people assigned to the task are", assigned)
	taskEnter := Task{
		date: date,
		startTime: startTime,
		duration: duration,
		assigned: assigned,
	}
	fmt.Println("Task-\nDate:", taskEnter.date.Format(dateFormat), "\nStartTime:", taskEnter.startTime.Format(timeFormat), "\nDuration", taskEnter.duration, "\nAssigned To:", taskEnter.assigned)
	fmt.Println("Task added")
	return taskEnter
}

// slice structure to store to-do
var myQueue []ToDo

func nextToDo() []ToDo {
	if len(myQueue) == 0 {
		fmt.Println("The To-Do list is currently empty")
		return myQueue
	}
	firstToDo, myQueue := myQueue[0], myQueue[1:]
	if firstToDo.event != nil {
		fmt.Println("Event-\nDate:", firstToDo.event.date.Format(dateFormat), "\nStart Time:", firstToDo.event.startTime.Format(timeFormat), "\nLocation", firstToDo.event.location)
	}
	if firstToDo.task != nil {
		fmt.Println("Task-\nDate:", firstToDo.task.date.Format(dateFormat), "\nStartTime:", firstToDo.task.startTime.Format(timeFormat), "\nDuration", firstToDo.task.duration, "\nAssigned To:", firstToDo.task.assigned)
	}
	return myQueue
}

func viewToDo() {
	if len(myQueue) == 0 {
		fmt.Println("The To-Do list is currently empty")
	}
	for i, todo := range myQueue {
		if todo.event != nil {
			fmt.Println(i+1,": Event-\nDate:", todo.event.date.Format(dateFormat), "\nStart Time:", todo.event.startTime.Format(timeFormat), "\nLocation", todo.event.location)
		}
		if todo.task != nil {
			assignedString := strings.Join( todo.task.assigned, ",")
			fmt.Println(i+1, ": Task-\nDate:", todo.task.date.Format(dateFormat), "\nStartTime:", todo.task.startTime.Format(timeFormat), "\nDuration:", todo.task.duration, "\nAssigned To:", assignedString)
		}
	}
}

func promptUserMessages() {
	fmt.Println("To exit enter 'quit'.")
	fmt.Println("If you would like to view and remove the first item in your To-Do list, enter 'next'")
	fmt.Println("If you would like to view all your items in your To-Do list, enter 'view'")
	fmt.Println("To add an Event, enter 'event'")
	fmt.Println("To add an Task, enter 'task'")
}

func main() {
	fmt.Println("Hi there!")
	fmt.Println("Welcome to the To-Do-List manager.")
	reader = bufio.NewReader(os.Stdin)
	promptUserMessages()
	state := true
	for state {
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
		}else if toDo == "next" {
			myQueue = nextToDo()
		} else if toDo == "view"{
			viewToDo()
		} else {
			promptUserMessages()
			fmt.Println("What you have entered is invalid")
		}
	}
	fmt.Println("Thank you for using the To-Do-List manager")
}
