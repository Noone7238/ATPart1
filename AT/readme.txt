Sample commands:
go test
curl -v localhost:8080/sample 
curl -v localhost:8080/sample/2 
curl -X POST localhost:8080/sample -H "Content-Type: application/json" -d '{"name": "bobby","description": "a cool dude"}'
curl -X PUT localhost:8080/sample/2 -H "Content-Type: application/json" -d '{"name": "bobby","description": "a cool dude"}'
curl -X DELETE localhost:8080/sample/2 

I was able to get echo working and a light change on the archetecture but there is still work to be done on that. A couple of things that are not exactly how I should have done them.
-I need to clean up error checking in unit test
-I have global variables that should be cleaned up
-The products file needs to be split up a bit as it has a good portion of the code, validator, data base could be split out.
-The way I'm passing in the echo pointer and back is sloppy, but I wanted to get stuff pushed up.

Happily, while not perfect, I feel some solid progress was made this weekend. I'll try to address the above notes next drop. Also, any feedback is ofcourse welcome.