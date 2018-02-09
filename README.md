# Factorigo
Tech demo of reproducing Factorio using Go. Requires a local Factorio install to run.

![Current Progress](https://github.com/PeterBooker/factorigo/blob/master/factorigo.gif)

## Purpose

This is purely a learning exercise, a chance to explore systems and methods related to game design. It will never reproduce anywhere near the full Factorio functionality.

## Current Task

I have run into the challenge of only rendering relevant parts of the map for performance. While splitting the map into chunks and only rendering the closest seemed easy at first, I have now run into the problem of how to generate new chunks on demand and have them 'fit' with the current map.

## TODO

* Explore how GUI systems work and implement a proper menu system ingame.

* Write my own procedural generation libs and improve the map by combining lots of layers for surface, foliage, resources, trees, etc.

* Potentially explore asset loading and a loading screen.