using System;					
using System.Collections.Generic;


public class Queue
{
	public List<Todo> myList;
	public Queue()
	{
		this.myList = new List<Todo>();
	}
	
	public void enqueue(Todo value)
	{
		this.myList.Insert(0, value);
	}
	
	public Todo dequeue()
	{
		Todo toRemove = this.myList[this.myList.Count - 1];
		this.myList.RemoveAt(this.myList.Count-1);
		return toRemove;
	}
	
	public bool isEmpty()
	{
		return this.size() == 0;
	}
	
	public int size()
	{
		return this.myList.Count;
	}
}
