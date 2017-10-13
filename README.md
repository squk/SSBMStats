# HopelessStats

## Task List

- [x] Read Dolphin memory in OSX/Linux
- [ ] Read Dolphin memory in Windows
- [x] Read Melee player information
- [x] Concurrency (almost all functionality should be concurrent)
- [ ] Usable UI
- [ ] Place importance on stats at high percent/low stock

## Secondary Task List (before MVP)
- [ ] Installable package on Windows
- [ ] Auto-Updating

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


