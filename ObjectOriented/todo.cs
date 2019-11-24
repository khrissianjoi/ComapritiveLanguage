using System;
using System.Collections.Generic;

public class Todo
{ 
    public Tuple<int, int, int> date = new Tuple <int, int, int>(0,0,0);
    public Tuple<int, int> startTime = new Tuple <int, int>(0,0);
    public Todo(Tuple<int, int, int> date, Tuple<int, int> startTime) {
      this.date = date;
      this.startTime = startTime;
    }
}

class Task : Todo
{ 

  private int duration;
  private List<string> assigned = new List<string>();

  public Task(Tuple<int, int, int> date, Tuple<int, int> startTime, int duration,List<string> assigned) 
    : base (date, startTime)
  {
    this.duration = duration;
    this.assigned = assigned;
  }

  public override string ToString() {
    string assignedPeople = String.Join(", ", this.assigned);
    return string.Format("Task-\nDate: {0}\nStartTime: {1}\nDuration: {2}\nAssigned To: {3}",date,startTime,this.duration, assignedPeople);
  }
}

class Event : Todo
{
  private string location;

  public Event(Tuple<int, int, int> date, Tuple<int, int> startTime, string location)
    : base (date, startTime)
  {
    this.location = location;
  }

  public override string ToString() {
    return string.Format("Event-\nDate: {0}\nStartTime: {1}\nLocation: {2}",date,startTime,this.location);
  }
}