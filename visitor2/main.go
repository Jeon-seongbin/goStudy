package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left  Expression
	right Expression
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type ExpressionPointer struct {
	sb strings.Builder
}

func (e *ExpressionPointer) VisitDoubleExpression(de *DoubleExpression) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPointer) VisitAdditionExpression(ae *AdditionExpression) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionPointer() *ExpressionPointer {
	return &ExpressionPointer{strings.Builder{}}
}

func (e *ExpressionPointer) String() string {
	return e.sb.String()
}

func main() {
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{1},
			right: &DoubleExpression{1},
		},
	}
	ep := NewExpressionPointer()
	ep.VisitAdditionExpression(e)
	fmt.Println(ep.String())
}
