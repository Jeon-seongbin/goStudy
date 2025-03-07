package main

import (
	"sync"
)

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(*Query)
}

type Observable interface {
	Subscribe(o Observer)
	UnSubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game
	Name            string
	attack, defense int
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{
		game:    game,
		Name:    name,
		attack:  attack,
		defense: defense,
	}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

type CreatureModifier struct {
	game     *Game
	creatrue *Creature
}

func (c *CreatureModifier) Handle(*Query) {}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{
			game:     g,
			creatrue: c,
		},
	}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creatrue.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func main() {
	game := &Game{sync.Map{}}
	g := NewCreature(game, "string goblin", 2, 2)
	// fmt.Println(g.String())

	m := NewDoubleAttackModifier(game, g)
	m.Close()
}
