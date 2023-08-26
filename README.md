# Hourly News Go (WIP)
## About
#### Purpose of the project
1. To provide a way to request news in JSON format that refreshes with news every hour between 7am and 8pm
2. To experiment with implementing RESTful APIs and to understand how they work
#### Technologies
* Gin Gonic
* Google's UUID library
* NewsAPI(REST API)
#### What has been implemented so far
[ ] Create Key
[ ] NewsAPI wrapper
[ ] Sends response
[ ] Admin privileges
[ ] Custom Query




## Development and Design
Originally, this project was going to be developed in rust; however, after trying to prototype an API Wrapper for getting web information with the Custom Search JSON API, I decided to move to Go and NewsAPI. 

#### Why did I switch?
I find that JSON parsing in Go is much more intuitive. Along with that, instead of trying to develop and maintain my own search engine with the Custom Search JSON API, which only returns the generic news sites not the actual news, I decided to use NewsAPI. It is maintained and updated by their developers, so it would be easier to create a more reliable search.
#### Back to Design
My goal for this project is to have a working REST API that will allow a user to retrieve the news every hour. The program will retrieve the news every hour between 7am and 8pm. To simply put why I chose these hours, nothing much happens outside of these hours. It also saves on the amount of queries being produced allowing for custom searches.
#### API Endpoints
* GET /news?=APIKEY - returns JSON Object
* GET /key - returns string
* GET /quit - Used by admin to shutdown the server remotely
* POST /reset - Used by admin to reset the server remotely. It also gets rid of all known API keys that are not admins
* POST /query?=QUERYTYPE

#### What does it do for the user?
This program/REST API is designed to be used by people trying to automate news collection. The API would allow people to generate a key and search by using the API's endpoints without creating a key through News API, parsing the information, and creating a search. Because the API is designed as an automated search, I plan to create a discord bot that the server would send the information to after searching for it. Because my planned discord bot's use would be contingent upon the state of this application, the /quit GET request would also send a request to the bot to shut down the bot.

## Future Plans
One of the first things I would like to do is to expand upon the functionality of this API. I would like to be able to add bots to some data structure so that when the bot is ready and is able to send data, it would automatically output the data to those bots.\
I am planning on also adding a neutral or blank endpoint(GET /) to allow users not using JSON retrival to get an HTML form of the page without using a key at all. This would allow people, who are not going to be using this as a component, to get the information through their browser.
