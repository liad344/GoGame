package main

import "liad/game/Controllers"

type item struct {
}

type Character struct {
	Items      item
	Controller Controllers.Controller
}
