package main

import (
	"machine"
	"time"
)

type Flasher struct {
	p machine.Pin
}

func NewFlasher() (f *Flasher) {
	f = new(Flasher)

	f.p = machine.LED
	f.p.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return
}

func (f *Flasher) FlashyFlash() {
	for i := 0; i < 15; i++ {
		f.p.Low()
		time.Sleep(time.Millisecond * 250)

		f.p.High()
		time.Sleep(time.Millisecond * 250)
	}
}

func (f *Flasher) Off() {
	f.p.Low()
}
