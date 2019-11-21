using System;
using System.Collections.Generic;
using Task;					

public class Queue
{
	public List<int> myList = new List<int>();
	public Queue()
	{
		this.myList = myList;
	}
	
	public void enqueue(int n)
	{
		this.myList.Insert(0, n);
	}
	
	public int dequeue()
	{
		int toRemove = this.myList[this.myList.Count - 1];
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
	
	public static void Main()
	{
		Queue q = new Queue();
		q.enqueue(1);
		q.enqueue(2);
		q.enqueue(3);
		Console.WriteLine(q.myList[1]);
		Console.WriteLine(q.dequeue());
		Console.WriteLine(q.isEmpty());
		Queue p = new Queue();
		Console.WriteLine(p.isEmpty());
		
	}
}