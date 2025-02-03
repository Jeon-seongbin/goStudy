package main

import "fmt"

type Person struct {
	Name    string
	Room    *Chatroom
	chatlog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s : %s", sender, message)
	fmt.Println(s)
	p.chatlog = append(p.chatlog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type Chatroom struct {
	people []*Person
}

func (c *Chatroom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *Chatroom) Join(p *Person) {
	joinMsg := p.Name + " join the chat"
	c.Broadcast("room", joinMsg)
	p.Room = c
	c.people = append(c.people, p)
}

func (c *Chatroom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func main() {
	room := Chatroom{}
	john := NewPerson("john")
	jane := NewPerson("jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi")
	jane.Say("hi")

}
