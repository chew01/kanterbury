# Kanterbury

<img src="docs/princess.jpg" width="200">

Kanterbury is a proxy server for Guardian Tales to receive game state updates, inspired by [Rhine](https://github.com/kyoukaya/rhine).

**Work in progress!**

### Accessible game states
- Player Data (initialized upon log in)
  - Player ID
  - App ID (Region)
  - World Name
  - Start Time
- Display Character Data (updated on every item/resource update)
  - Character ID
  - Character Level
  - Character Name
- Activity Data (updated only on PvP)
  - Activity Name
  - Activity End Time (used to check if activity is over)

Game state is still rather limited as most of the other data can't be intercepted by the HTTP(S) proxy server...

### Usage

WIP