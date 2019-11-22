using System;					
using System.Collections.Generic;

public class Queue<T>
{
	private List<T> myList = new List<T>();
	public Queue()
	{
		this.myList = myList;
	}
	
	public void enqueue(T value);
	{
		this.myList.Insert(0, value);
	}
	
	public T dequeue()
	{
		T toRemove = this.myList[this.myList.Count - 1];
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
