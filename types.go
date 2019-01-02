package main

type Error int

const (
	unknown Error = 0
	console Error = 1
	file    Error = 2
)

type Lang int

const (
	en Lang = -1
	ru Lang = 0
	de Lang = 1
)
