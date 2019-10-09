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

To test everything is working just call: 

```
http://localhost:8080/
```

## How to use

### Login

You will need to perform the following request in order to get the access token:

```
curl -H "Content-Type: application/json" --request POST --data '{"email":"$$YOUREMAIL$$","password":"$$YOURPASSWORD$$"}' http://localhost:8080/login
```

## Config

### Account

Now that everything seems to work, you must fill the application.properties file with 
your account details.

To obtain the userId, leagueId and biwengerVersion values, you have to access to your 
biwenger account through your browser and capture any request done by the browser. For 
this purpose you should use the browser console. For instance, when you click 'Home' on
your biwenger session, there's a request with the headers x-league, x-user and x-version.
Those are the values for your userId, leagueId and biwengerVersion respectively. 

Also, you can create a twitter account and set the values in the applications.properties
if you want to activate the tweet functionality.

At this point, you have to build again your docker image:

```
docker-compose up -d --build
```

### DB

For everything to work fine, you'll need some players information on your mongodb. For this
you should run the following call:

```
curl -H "Content-Type: application/json" --request PUT http://localhost:8080/players
```

Then connect to the *biwenger* database (you will need a mongodb client)
and check that there is data in the players collection:

```
> use biwenger
> db.players.find()
```


## Api Calls

Once you got the token, you can use it in a real call:

```
curl -H "Content-Type: application/json" --request GET -H "authorization":"Bearer $$YOURTOKEN$$" http://localhost:8080/squad
```

All the api calls examples are provided in resources/requestsExamples.txt

## Development

After doing any changes in the code, a new docker image should be created:
```
docker-compose up -d --build
```

## Current Features

* Login
* Get your entire squad
* Put your line up
* Get your money
* Get your max bid
* Get the market evolution
* Put your squad to the market
* Get all players in the market
* Place an offer to the market
* Get the offers from the market
* Accept an offer for a player
* Get all players
* Update player alias in your database
* Get the days left for the next round

## Licensing
MIT: http://rem.mit-license.org
