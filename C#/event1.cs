using System;
using System.Collections.Generic;

class Event
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

  public bool work()
  {
    return true;
  }
}