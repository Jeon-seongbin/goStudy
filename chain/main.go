package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d, %d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{name, attack, defense}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
		return
	}
	c.next = m
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifie(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{
			creature: c,
		}}
}

type IncreasedDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefanceModifier(c *Creature) *IncreasedDefenseModifier {
	return &IncreasedDefenseModifier{
		CreatureModifier: CreatureModifier{
			creature: c,
		}}
}

func (i *IncreasedDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("increasing", i.creature.Name)
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

func (d *DoubleAttackModifier) Handle() {
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{
		CreatureModifier: CreatureModifier{
			creature: c,
		}}
}

func (n *NoBonusesModifier) Handle() {

}

func main() {
	g := NewCreature("G", 1, 1)
	root := NewCreatureModifier(g)

	root.Add(NewDoubleAttackModifie(g))
	root.Add(NewIncreaseDefanceModifier(g))
	root.Add(NewDoubleAttackModifie(g))

	root.Handle()
	fmt.Println(g.String())
}
