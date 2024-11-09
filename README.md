# Two Fat Lads (Heavy Industries) - Bob Hoskins Excuse Boskins

A USB HID based on the RP2040 (though I'm sure any board will work really- there's nothing RP2040 special here), as found in the Raspberry Pico, which types out excuses for when you get an invite and it sounds shit.

The Bob Hoskins Excuse Boskins does this by emulating a USB keyboard; when plugged in the device looks like:

```
Bus 001 Device 080: ID 2e8a:000a Two Fat Lads (Heavy Industries) Bob Hoskins Excuse Boskins
```

## Building

You will need:

1. A board of some sort (honestly... the first generation pico costs 3 quid; you wont need headers for this project)
2. [tinygo](https://tinygo.org/getting-started/install/)
3. [go](https://go.dev/dl/)

From there, plug your device in and make it flashable (in the case of the pico, hold down the `BOOTSEL` button as you plug it in to your computer's USB drive). Once ready, you can run:

```bash
$ tinygo flash -target=pico
```

And you'll be good to go

### Caveat

After flashing the pico the board will restart and you'll suddenly find an excuse printed to your screeen. Be aware.

## Debugging

Errors and debug information can be found by connected to the USB over serial. Your best bet, on most things non-Windows, is to run:

```bash
$ cat /dev/ttyACM0
```

After the device starts up (you'll know when the on-board LED starts flashing on/off).

## Running

Once installed, plug the Excuse Boskins into your laptop, make sure the place you want your excuse to go (email client, chat client, calendar, whatsapp, slack, whatever) is active, and your excuse will appear as if typed out.
