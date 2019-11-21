using System;
using System.Collections.Generic;

class Task
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
    Console.WriteLine(this.date);
  }

  public bool work()
  {
    return true;
  }
}