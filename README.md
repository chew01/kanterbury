# Kanterbury

<img src="docs/princess.jpg" alt="Future Princess of Kanterbury" width="150">

Kanterbury is a HTTPS proxy server for Guardian Tales to receive game state updates, inspired by [Rhine](https://github.com/kyoukaya/rhine). At this moment in time, Kanterbury intercepts logging data that is sent from the client to the server on login, resource changes, and multiplayer gamemode activity, then updates a game state instance built into the proxy.

### Accessible game state data
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

### Applications

- [GT-RPC](https://github.com/chew01/gt-rpc), a Discord Rich Presence client for Guardian Tales that uses Kanterbury's game state

### Usage

Run `go build cmd/main.go && ./main.exe` to build and run the proxy server.

If there is no `cert.pem` and `key.pem` present in the folder, Kanterbury will generate a root CA to be installed in your emulator/device, so that HTTPS traffic can be intercepted.

### Background

As mentioned above, Kanterbury is inspired by [Rhine](https://github.com/kyoukaya/rhine) and was written to provide a simple way to record game state and other properties. However, the data that we are able to obtain is very,  very limited due to the fact that most game traffic does NOT go through the HTTPS protocol. It is likely that a TCP/UDP proxy will have to be used to intercept this traffic, and other ways of decryption may be required depending on such. This is something that I intend to work on in the future. 

Goroutines were not used in the current implementation of Kanterbury because there is very little need to do so due to the low frequency of updates. If it is required in the future, I will use them.