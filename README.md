SSBMStats
============

## Description
A Go program using my Golphin library for SSBM. Reads updates in the VRAM of
the Dolphin emulator and provides verbose statistics on approaches, neutral
game, successful punishes, tech skill 'perfect-ness' and more.

## Usage
### Installation
Requires the [Golphin](https://www.github.com/ctnieves/Golphin) package.
```
# Prerequisites
go get github.com/ctnieves/Golphin


# Download
git clone https://github.com/ctnieves/SSBMStats.git
cd SSBMStats

#Run
go run main.go
```


## Task List
- [x] Read Dolphin memory in OSX/Linux
- [x] Read Melee player information
- [x] Concurrency (almost all functionality should be concurrent)
- [ ] Place importance on stats at high percent/low stock

## WIP
- [ ] Read Dolphin memory in Windows
- [ ] Usable UI (WIP)

## Secondary Task List (before MVP)
- [ ] Installable package on Windows
- [ ] Auto-Updating
- [ ] APM Calculation

## Questionable Tasks
- [ ] Web UI instead of CLI
- [ ] Electron UI instead of Web/CLI
- [ ] Completely verbose tracking. (Memory limitations. Confusing to encode all frame data and pre-marked flags. Could be useful to 'upgrade' old stats.)


## Link Dump(Stuff to be used maybe?)
Deploying the app on windows machines
https://github.com/mh-cbon/go-msi

Cross platform auto-updates(choose one eventually)
https://github.com/inconshreveable/go-update
https://github.com/sanbornm/go-selfupdate
https://github.com/keybase/go-updater

Win32 bindings. Will use ReadProcessMemory for Dolphin mem-reading on Windows
https://github.com/AllenDang/w32


