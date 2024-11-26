package main

import "fmt"


type Queue struct{
	element []int
}
func (q *Queue)Enqueue(data int){
	q.element=append(q.element, data)
}
func (q *Queue) deque(){
	if len(q.element)==0{
		fmt.Println("Queue is empty")
		return
	}
	q.element=q.element[1:]
}
func main(){
	que:=Queue{}
	que.Enqueue(4)
	que.Enqueue(5)
	que.Enqueue(5)
	que.deque()
	que.deque()
	que.deque()
	fmt.Println(que)
}