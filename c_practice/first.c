using System;
					

public class Queue
{
	public string[] items = new string[1];  
	public Queue()
	{
		this.items = items;
	}
	
	public void enqueue(string item) {
		Array.Resize(ref this.items, this.items.Length + 1);
		this.items[this.items.GetUpperBound(0)] = item;
	}
	
	public string dequeue() {
		return this.items[-1];
	}
	public static void Main()
	{
		Queue q = new Queue();
		string val = Console.ReadLine();
		q.enqueue(val);
		val = Console.ReadLine();
		q.enqueue(val);
		val = Console.ReadLine();
		q.enqueue(val);
		val = Console.ReadLine();
		q.enqueue(val);
		foreach(var elem in q.items)
		{
    		Console.WriteLine(elem);
		}
	}
}


// https://dotnetfiddle.net/