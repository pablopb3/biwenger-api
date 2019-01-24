# biwenger-api
Api for the football manager game biwenger. It will allow you to do every operation 
you need like get all your players, get all the info from an specific player, 
set your lineup, get all players from the market, make a bid on a player...

Built with docker, go and mongodb

## Getting started

docker-compose up in the project root

To test everything is working just call: http://localhost:8080/

## How to use

### Login

You will need to perform the following request in order to get the access token:
curl --header "Content-Type: application/json"   --request POST   
--data '{"email":"$$YOUREMAIL$$","password":"$$YOURPASSWORD$$"}' http://localhost:8080/login

### Api Call Example

Once you got the token, you can use it in a real call:
curl --header "Content-Type: application/json" --request GET  
-H "authorization":"Bearer $$YOURTOKEN$$" 
http://localhost:8080/getMyPlayers

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
