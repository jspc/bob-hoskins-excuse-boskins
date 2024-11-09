package main

import (
	"math/rand"
	"time"
)

var excuses = []string{
	"I'm sorry, I can't make it. I think you're a cunt",
	"My mum's dog's neighbour's cat's dog's going to be licking up sick at that time, and I don't want to miss it",
	"I'm not available then, I'll be playing Bob Hoskins' Hot Boskins",
	"I'd love to, but to put it bluntly, I'm shitting through the eye of a needle",
	"I would but I've just realised I fucking hate everything about you",
	"I need to see a man about shagging your mum",
}

func excuse() string {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))

	return excuses[r.Intn(len(excuses))]
}
