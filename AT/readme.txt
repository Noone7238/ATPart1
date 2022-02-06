Sample commands:
curl -v localhost:8080/sample/ 
^^print all

curl -v localhost:8080/sample/1 
^^print ID 1

curl localhost:8080/sample/ -X POST -d '{"name": "Car", "description": "A vehicle for transportation", "id": "3"}'
^^add new ID is currently used

curl localhost:8080/sample/3 -X POST -d '{"name": "Car"}'
^^partial update ID field is not updatable and is ignored here

curl -X DELETE http://localhost:8080/sample/2
^^Delete ID 2

-I added those just incase I was doing them in a slightly weird way. I'll continue to read up on API's to improve.


Currently this is a first pass at the assignment, I am going to be uploading them piece by piece. That way yall can see how things are going, also assist communication. Please give feedback as you desire, any feedback is welcome. I'm hoping to learn and prepare, hopefully for us eventually working together. Sorry if progress on the weekdays is a bit slow. I still work about 8-9 hours a day monday-friday.

This hits all the primary todo items. My next planned steps are as follows, the order is however loose and is subject to change via feedback:

- Add Testing.
- Add logging and some sort of Printed Log file.
- Saved file for initial database, read in at opening, writen out at closing. (or maybe as added and removed items occur)
- Reconfigure code architecture to actually utilze packages and be overall cleaner.
- Better error checking and code stability.
- Address any feedback given.

Current Questions:
-You requested the data structure be an array. I can swap it over, I used a map because it was easier to use and remove from but I can swap it if you'd like.

-What Packages should I be using? I know Gorilla Mux is common, I could add that if you'd like. I tried to use less packages as I wasn't sure which ones you wanted.

-Is there any order you want the above items? I plan on trying to tackle Testing tomorrow. Then I want to start reading on logging and code architecture next as I want to get better at following the starndard go style.

-Are the above items all necessary? I am planning on doing them for my own personal growth, eventually. If some are less important to you for the matter at hand then I'll move them off the list until after.

