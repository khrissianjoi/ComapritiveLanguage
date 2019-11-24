using System;
using System.Collections.Generic;

namespace com.todo
{
  class Todo {};
  class Task {};
  class Event {};
}

namespace com.queue
{
  class Queue {};
}

class ToDoManager {
  public Queue toDoList;
  public ToDoManager() {
    this.toDoList = new Queue();
  }

  public static Tuple<int, int, int> isValidDate() {
    while(true) {
      try {
        Console.WriteLine("Please Enter a date dd/mm/yyyy");
        string dateEntered = Console.ReadLine();
        String[] dateList = dateEntered.Split("/");
        int date = Convert.ToInt32(dateList[0]);
        int month = Convert.ToInt32(dateList[1]);
        int year = Convert.ToInt32(dateList[2]);
        return new Tuple <int, int, int>(date,month,year);
      } catch {
        Console.WriteLine("Invalid");
      }
    }
  }

  public static Tuple<int,int> isValidTime() {
    while(true) {
      try {
        Console.WriteLine("Please Enter a start time hh:mm");
        string startTime = Console.ReadLine();
        String[] timeList = startTime.Split(":");
        int hour = Convert.ToInt32(timeList[0]);
        int minute = Convert.ToInt32(timeList[1]);
        return new Tuple <int,int>(hour, minute);
      } catch {
        Console.WriteLine("Invalid");
      }
      }
  }

  public static int isDuration() {
    while(true) {
      try {
        Console.WriteLine("Please Enter the duration");
        int duration = Convert.ToInt32(Console.ReadLine());
        return duration;
      } catch {
        Console.WriteLine("Invalid");
      }
    }
  }

  public static List<string> isAssigned() {
    Console.WriteLine("Please Enter the people assigned to this task, seperated my a comma ','");
    string[] userInput = Console.ReadLine().Split(",");
    List<string> assigned = new List<string>(userInput);
    return assigned;
  }
  public static Task taskSelected() {
    Tuple<int, int, int> date = isValidDate();
    Tuple<int,int> startTime = isValidTime();
    int duration = isDuration();
    List<string> assigned = isAssigned();
    Task validTask = new Task(date,startTime,duration,assigned);
    Console.WriteLine(validTask);
    Console.WriteLine("Task added");
    return validTask;
  }

  public static Event eventSelected() {
    Tuple<int, int, int> date = isValidDate();
    Tuple<int,int> startTime = isValidTime();
    Console.WriteLine("Please enter the location");
    string location = Console.ReadLine();
    Event validEvent = new Event(date,startTime,location);
    Console.WriteLine(validEvent);
    Console.WriteLine("Event added");
    return validEvent;
  }

  public void promptUserMessages() {
    Console.WriteLine("To exit enter 'quit'.");
    Console.WriteLine("If you would like to view and remove the first item in your To-Do list, enter 'next'");
    Console.WriteLine("If you would like to view all your items in your To-Do list, enter 'view'");
    Console.WriteLine("To add an Event, enter 'event'");
    Console.WriteLine("To add an Task, enter 'task'");
  }

  public void viewToDo() {
    if (this.toDoList.size() != 0) {
      for (int i = 0; i < this.toDoList.size(); i++) {
        string toDoInternation = string.Format("{0}:{1}",i+1, this.toDoList.myList[i]);
        Console.WriteLine(toDoInternation);
      }
    } else {
      Console.WriteLine("The To-Do list is currently empty");
    }
  }

  public void nextToDo() {
    if (this.toDoList.size() != 0) {
      Todo currentOne = this.toDoList.dequeue();
      Console.WriteLine(currentOne);
    } else {
      Console.WriteLine("The To-Do list is currently empty");
    }
  }

  public void run() {
    Console.WriteLine("Hi there!");
    Console.WriteLine("Welcome to the To-Do-List manager.");
    bool state = true;
    this.promptUserMessages();
    while (state) {
        Console.WriteLine("Would you like to enter an Event or a Task?");
        string userInput = Console.ReadLine().ToLower();
        if (userInput == "event"){
          Event theEvent = eventSelected();
          this.toDoList.enqueue(theEvent);
        } else if (userInput == "task"){
          Task theTask = taskSelected();
          this.toDoList.enqueue(theTask);
        } else if (userInput == "quit") {
			    state = false;
        } else if (userInput == "view") {
          this.viewToDo();
        } else if (userInput == "next") {
          this.nextToDo();
        } else {
          Console.WriteLine("The input is invalid");
        }
    }
    Console.WriteLine("Thank you for using the To-Do-List manager");
    System.Environment.Exit(1);
  }
}

class MainClass {
  public static void Main() {
    ToDoManager toDoManager = new ToDoManager();
    toDoManager.run();
  }
}