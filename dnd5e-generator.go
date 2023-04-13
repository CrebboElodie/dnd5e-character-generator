package main

import (
	"fmt"
	"math/rand"
)

var level int
var class string
var classes = [12]string{"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"}
var background string
var backgrounds = [6]string{"Acolyte", "Criminal", "Folk Hero", "Noble", "Sage", "Soldier"}
var race string
var races = [9]string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Half-Elf", "Halfling", "Half-Orc", "Human", "Teifling"}
var alignment string
var alignments = [3][3]string{{"Lawful Good", "Neutral Good", "Chaotic Good"}, {"Lawful Neutral", "True Neutral", "Chaotic Neutral"}, {"Lawful Evil", "Neutral Evil", "Chaotic Evil"}}
var health int
var stats = [6]int{}
var stats_name = [6]string{"Str", "Dex", "Con", "Int", "Wis", "Cha"}

func main() {
	level = rand.Intn(20-1) + 1
	fmt.Printf("Lv:%d\n", level)

	class = classes[rand.Intn(11-0)]
	fmt.Printf("Class:%s\n", class)

	background = backgrounds[rand.Intn(5-0)]
	fmt.Printf("Background:%s\n", background)

	race = races[rand.Intn(8-0)]
	fmt.Printf("Race:%s\n", race)

	alignment = alignments[rand.Intn(2-0)][rand.Intn(2-0)]
	fmt.Printf("Alignment:%s\n", alignment)

	health = rand.Intn(20-1) + 1
	fmt.Printf("Health:%d\n", health)

	for index := 0; index <= 5; index++ {
		stats[index] = (rand.Intn(6-1) + 1) + (rand.Intn(6-1) + 1) + (rand.Intn(6-1) + 1)
		fmt.Printf("%s:%d\n", stats_name[index], stats[index])
	}
}
