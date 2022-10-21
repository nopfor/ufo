
# UFO

Float a UFO across the screen. Windows, Mac, and Linux.

Created for [An Artist’s History of Computer Viruses and Malware](https://sfpc.study/sessions/fall-22/artists-history-of-malware).

Built on [Ebitengine](https://ebitengine.org/).

## "Infection" Instructions

Follow these instructions to run the UFO binary every so often. Keep in mind that most of the time the program will exit immediately.

### Windows

Run the following to execute the program every 7 minutes, changing the path and binary name where appropriate:

```psh
schtasks.exe /create /tn "_UFO" /tr "C:\PATH\TO\ufo.exe" /sc minute /mo 7
```

To clean up:

```psh
schtasks /delete /tn "_UFO"
```

### Linux / Mac

Run `crontab -e` and enter the following, changing the path and binary name where appropriate:

```sh
*/7 * * * * /PATH/TO/ufo
```

To clean up, re-run `crontab -e` and remove the line that executes the binary.

## Build Instructions

Follow these instructions to build from source instead of running the included binaries. This is currently necessary for Macs with ARM processors (M1, M2).

`go` must be installed: https://go.dev/doc/install

### Windows

```
go mod tidy
go build -o ufo.exe
```

### Mac

```sh
go mod tidy
go build -o ufo-mac
```

### Linux

First, install dependencies:

```sh
sudo apt install -y build-essential libalut-dev libasound2-dev libc6-dev libgl1-mesa-dev libglu1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev mesa-utils pkg-config xorg-dev xvfb
```

Then build:

```sh
go mod tidy
go build -o ufo-mac
```
