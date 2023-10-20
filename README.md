# Kaurapuuro

[Here](https://github.com/0mlml/js-tanks) I wrote a tank game inspired by Wii Play's tanks game that works good enough, but I always wanted to network it. This repository is just my experimentation with hand-crafting packets per-byte and what goes into designing the netcode for a real time multiplayer game.

## Design
The repository is split into two folders: client and server. The client folder is a static HTML page (so it can be hosted on GitHub pages or opened as a file), but I might choose to serve the client from the server directly to streamline connecting to the server. The server is written in Go with the only external dependency (for now) being the [gorilla/websocket](https://github.com/gorilla/websocket) package to streamline creating and using websocket connections.

The client and server both have different packets defined in what I think is the best way (so not ES6 classes) per language. Both client and server have a deserialize and serialize function defined per packet and a deserialize and send(packet) function that looks up the type of packet based on the first byte.

I'm also partial to the "Register packet listener" paradigm, as I think that it's an elegant way to make the code modular.


[to be continued as I progress]

