# PokeTelnet

This application is a CTF Challenge that was used for Metared's CTF in 2021

It requires users to automate connections to a telnet server and decode messages in a made up alphabet dynamically created with emojis.

## Running the game

1. The server needs to execute having the emoji.json and the pokedex.json files in the same directory. After executing the binary, the server will start listening on port 5555

```
./poketelnet
```

2. Clients need to connect to the server via nc or telnet to port 5555, then they can start playing

```
nc 127.0.0.1 5555


	██╗███╗░░██╗░██████╗███████╗░█████╗░██╗░░░██╗██████╗░██╗████████╗██╗░░░██╗
	██║████╗░██║██╔════╝██╔════╝██╔══██╗██║░░░██║██╔══██╗██║╚══██╔══╝╚██╗░██╔╝
	██║██╔██╗██║╚█████╗░█████╗░░██║░░╚═╝██║░░░██║██████╔╝██║░░░██║░░░░╚████╔╝░
	██║██║╚████║░╚═══██╗██╔══╝░░██║░░██╗██║░░░██║██╔══██╗██║░░░██║░░░░░╚██╔╝░░
	██║██║░╚███║██████╔╝███████╗╚█████╔╝╚██████╔╝██║░░██║██║░░░██║░░░░░░██║░░░

	I'm sorry, my responses are limited. You must ask the right question.
	Prove you can speak my language fluently and quickly.
	Do it and you shall receive a flag.

a = 👨🏿‍❤‍💋‍👨🏾
b = 🍾
c = ♐
d = 👨‍❤️‍👨
e = 🤽🏿‍♂
f = 👨🏽‍❤️‍👨🏾
g = 🇵🇸
h = 👩🏾‍❤️‍💋‍👩🏿
i = 👨🏾‍❤️‍👨🏾
j = 👩🏿‍❤‍💋‍👩🏼
k = 👨🏾‍🤝‍👨🏻
l = 🛻
m = 🧍🏼‍♂
n = 🚣🏽‍♂️
o = 🤸🏾‍♂
p = 👸🏻
q = 🅾️
r = 🇵🇾
s = 🏓
t = 🙍‍♀️
u = 🟩
v = ❄
w = 🚐
x = 🤾🏼‍♂
y = 🧜🏽‍♂️
z = 🧘
🧍🏼‍♂👨🏿‍❤‍💋‍👨🏾♐👩🏾‍❤️‍💋‍👩🏿🤸🏾‍♂👸🏻 =>
```
