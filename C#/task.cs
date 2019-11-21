using System;
using System.Collections.Generic;

class Task
{ 
  public int date;
  public int startTime;
  public int duration;
  public List<string> assigned = new List<string>();

  public Task(int date, int startTime, int duration,List<string> assigned)
  {
    this.date = date;
    this.startTime = startTime;
    this.duration = duration;
    this.assigned = assigned;
  }

  public bool work()
  {
    return true;
  }
}