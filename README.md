# [Gotromino](https://github.com/nikarh/gotromino)

[![GitHub](https://img.shields.io/github/license/nikarh/gotromino)](https://github.com/nikarh/gotromino)
[![GitHub Repo stars](https://img.shields.io/github/stars/nikarh/gotromino)](https://github.com/nikarh/gotromino)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/nikarh/gotromino)](https://hub.docker.com/r/nikarh/gotromino)

A console Tetris™-like game written in go.
Kill your time, while you are waiting for a CI to build your branch.
Impress your colleges with your gaming skills during a video call with screen sharing
when you have already fixed a critical production bug, and you still have 10 minutes
for your E2E test to finish before trying to deploy it.


## Installation
There are multiple ways to get this game
- If you are a QA engineer, don't waste your time trying to compile this one from sources, just download [the latest release](https://github.com/nikarh/gotromino/releases/latest).
- If you are a DevOps engineer, just run this game from a Docker container

  ```shell
  docker run --rm -it nikarh/gotromino
  ```
- If you are a software developer, you can build this package from sources
  
  ```shell
  go get -u github.com/nikarh/gotromino
  ```
- If you are a senior DevOps engineer, you can build it from sources in a Docker container
  
  ```shell
  git clone https://github.com/nikarh/gotromino
  docker build -t nikarh/gotromino gotromino
  docker run --rm -it nikarh/gotromino
  ```

## Game
![Gameplay](https://github.com/nikarh/gotromino/blob/master/game.gif?raw=true)

Have fun!
