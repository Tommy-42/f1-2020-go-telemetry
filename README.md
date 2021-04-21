# F1-2020-Go-Telemetry

# Purpose
> The F1 series of games support the output of certain game data across UDP connections. This data can be used supply race information to external applications, or to drive certain hardware (e.g. motion platforms, force feedback steering wheels and LED devices).

F1-2020-Go-Telemetry will help you handle the output the F1 2020 game data sent through UDP connections, store the data into a repository ( Elasticsearch at the moment ), allowing you to create graphs ( Grafana ) with the data sourced from the repository

# Requirements
- Docker
- Go >= 1.16

# Game Setup

Setting up the F1 Game
In order to capture data, you'll need to set up your F1 2020 game.

Run the F1 2020 Game
From the main menu, open "Game Options"
Open Settings -> Telemetry Settings
Turn UDP Telemetry On


You can adapt the configuration to your preference
If you want more detailed telemetry, you can up the UDP send rate. Be warned, the higher the send rate, the larger storage you'll need


# Local Dev Setup

start the dependencies ( elasticsearch, kibana )
```bash
docker-compose up -d
```

```bash
go run -race main.go
```
