using System;
using System.Collections.Generic;

public class Todo
{
    public Todo() {

    }
}

class Task : Todo
{ 
  public Tuple<int, int, int> date = new Tuple <int, int, int>(0,0,0);
  public Tuple<int, int> startTime = new Tuple <int, int>(0,0);
  public int duration;
  public List<string> assigned = new List<string>();

  public Task(Tuple<int, int, int> date, Tuple<int, int> startTime, int duration,List<string> assigned)
  {
    this.date = date;
    this.startTime = startTime;
    this.duration = duration;
    this.assigned = assigned;
  }

  public override string ToString()
  {
    return string.Format("Task-\nDate: {0}\nStartTime: {1}\nDuration: {2}\nAssigned To: {3}",this.date,this.startTime,this.duration,this.assigned);
  }
}

class Event : Todo
{
  public Tuple<int, int, int> date = new Tuple <int, int, int>(0,0,0);
  public Tuple<int, int> startTime = new Tuple <int, int>(0,0);
  public string location;

  public Event(Tuple<int, int, int> date, Tuple<int, int> startTime, string location)
  {
    this.date = date;
    this.startTime = startTime;
    this.location = location;
  }

  public override string ToString()
  {
    return string.Format("Task-\nDate: {0}\nStartTime: {1}\nLocation: {2}",this.date,this.startTime,this.location);
  }
}