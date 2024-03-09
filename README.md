# pa-mute-to-awtrix

This script mutes/unmutes a PulseAudio source and indicates source status via a few pixels on an [awtrix](https://github.com/Blueforcer/awtrix-light) pixel clock.

## Usage

```shell
# Toggle mute on default PA source, send status to awtrix at 10.0.0.2

./pa-mute-to-awtrix -host=10.0.0.2

# Don't toggle, only send mute status to awtrix

./pa-mute-to-awtrix -host=10.0.0.2 -onlyCheck

# Use blue indicator

./pa-mute-to-awtrix -host=10.0.0.2 -color=#0000FF

# Use specific PA source

./pa-mute-to-awtrix -host=10.0.0.2 -source=alsa_input.pci-0000_0e_00.4.analog-stereo
```

## ...why?

Very niche, but since switching to Wayland, my global PTT hotkeys would not work with certain applications. Muting/unmuting the driver via global hotkeys is the next best thing, but I was missing a prominent indicator.