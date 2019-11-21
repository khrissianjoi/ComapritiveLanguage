using System;
using System.Collections.Generic;

class Event
{
  public int date;
  public int startTime;
  public int location;

  public Event(int date, int startTime, int location)
  {
    this.date = date;
    this.startTime = startTime;
    this.location = location;
  }

  public bool work()
  {
    return true;
  }
}