package main

import "fmt"

type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (s *BaseState) On(sw *Switch) {
	fmt.Println("already on")
}

func (s *BaseState) Off(sw *Switch) {
	fmt.Println("already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turn on")
	return &OnState{BaseState{}}
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turnd off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("turning right on")
	sw.State = NewOnState()
}
func main() {
	sw := NewSwitch()
	fmt.Println("-")
	sw.Off()
	sw.Off()
	fmt.Println("-")
	sw.On()
	fmt.Println("-")
	sw.Off()
	fmt.Println("-")
	sw.Off()
}
