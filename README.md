# biwenger-api
Api for the football manager game biwenger. It will allow you to do every operation 
you need like get all your players, get all the info from an specific player, 
set your lineup, get all players from the market, make a bid on a player...

Built with docker, go and mongodb

## Getting started

In the project root run:

```
docker-compose up -d
```

To test everything is working just call: `http://localhost:8080/v1/`

## How to use

### Login

You will need to perform the following request in order to get the access token:

```
curl --header "Content-Type: application/json" --request POST --data '{"email":"$$YOUREMAIL$$","password":"$$YOURPASSWORD$$"}' http://localhost:8080/v1/login
```

### Api Call Example

Once you got the token, you can use it in a real call:

```
curl --header "Content-Type: application/json" --request GET -H "authorization":"Bearer $$YOURTOKEN$$" http://localhost:8080/v1/getMyPlayers
```

### Test your Setup

To test that the database is storing data you can run the following call:

```
curl --header "Content-Type: application/json" --request GET -H "authorization":"Bearer $$YOURTOKEN$$" http://localhost:8080/v1/updatePlayersAlias
```

Then connect to the *biwenger* database (you will need a mongodb client installed)
and check that there is data in the players collection:

```
> use biwenger
> db.players.find()
```

## Development

After doing any changes in the code, a new docker image should be created:
docker-compose build
docker-compose up

## Current Features

* Login
* UpdatePlayersAlias
* GetPlayerById
* GetMyPlayers
* SetLineUp

## Licensing
MIT: http://rem.mit-license.org
