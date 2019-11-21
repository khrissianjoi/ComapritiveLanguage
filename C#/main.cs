using System;
using System.Collections.Generic;

namespace com.queue
{
  class Queue {};
}

namespace com.task
{
  class Task {};
}

namespace com.event1
{
  class Event {};
}

class MainClass {
  public static void test() {
    Queue q = new Queue();
		q.enqueue(1);
		q.enqueue(2);
		q.enqueue(3);
		Console.WriteLine(q.myList[1]==2);
		Console.WriteLine(q.dequeue()==1);
		Console.WriteLine(q.isEmpty()==false);
		Queue p = new Queue();
		Console.WriteLine(true==p.isEmpty());
    List<string> assigned = new List<string>();
    Task t = new Task(1,1,1,assigned);
    Console.WriteLine(true==t.work());
    Event e = new Event(1,1,"hello");
    Console.WriteLine(e.work()==true);
  }

  public bool isValidDate() {
    return true;
  }

  public static Task taskSelected() {
    Console.WriteLine("Please Enter a date dd/mm/yyyy");
    int date = Convert.ToInt32(Console.ReadLine());
    Console.WriteLine("Please Enter a start time hh:mm");
    int startTime = Convert.ToInt32(Console.ReadLine());
    Console.WriteLine("Please Enter the duration");
    int duration = Convert.ToInt32(Console.ReadLine());
    Console.WriteLine("Please Enter the people assigned to this task, seperated my a comma ','");
    List<string> assignedTemp = new List<string>();
    // var assigned = Console.ReadLine();
    Task validTask = new Task(date,startTime,duration,assignedTemp);
    return validTask;
  }

    public static Event eventSelected() {
    Console.WriteLine("Please Enter a date dd/mm/yyyy");
    int date = Convert.ToInt32(Console.ReadLine());
    Console.WriteLine("Please Enter a start time hh:mm");
    int startTime = Convert.ToInt32(Console.ReadLine());
    Console.WriteLine("Please Enter the location");
    string location = Console.ReadLine();
    Event validEvent = new Event(date,startTime,location);
    return validEvent;
  }
  public static void Main(){
    // test();
    Console.WriteLine("Hi there!");
    Console.WriteLine("Welcome to the To-Do-List manager.");

    bool matchToDo = false;
    while(!(matchToDo)) {
      Console.WriteLine("Would you like to enter an Event or a Task?");
      string userInput = Console.ReadLine().ToLower();
      if(userInput == "event"){
        Event theEvent = eventSelected();
        matchToDo = true;
      }
      else if (userInput == "task"){
        Task theTask = taskSelected();
        matchToDo = true;
      }
      else {
        Console.WriteLine("The input is invalid");
      }
    }
    Console.WriteLine("SUCCESS");
    System.Environment.Exit(1);
  }
}