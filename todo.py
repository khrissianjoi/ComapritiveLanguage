class Queue:
    def __init__(self):
        self.items = []

    def enqueue(self, item):
        self.items.insert(0, item)

    def dequeue(self):
        return self.items.pop()

    def isEmpty(self):
        return self.items == []

    def size(self):
        return len(self.items)

    def __str__(self):
        if self.size():
            tokens = [str(data) for data in self.items if data]
            return "Your queue looks like this: {}".format(' -> '.join(tokens))
        else:
            return "Your queue is empty."


class Task:
    def __init__(self, data):
        self.date = data[0]
        self.startTime = data[1]
        self.duration = data[2]
        self.assigned = data[3]  # list of people

    def __str__(self):
        return "{}, {}, {}, {}".format(self.date, self.startTime,
                                       self.duration, self.assigned)


class Event:
    def __init__(self, data):
        self.date = data[0]
        self.startTime = data[1]
        self.location = data[2]

    def __str__(self):
        return "{}, {}, {}".format(self.date, self.startTime, self.location)


def main():
    myQueue = Queue()
    with open('ToDoList.txt', 'r') as contents:
        for content in contents:
            data = content.strip().split(',', maxsplit=3)
            # to avoid splitting inside list
            if len(data) == 4:
                myQueue.enqueue(Task(data))
            else:
                myQueue.enqueue(Event(data))
    print(myQueue)


main()
