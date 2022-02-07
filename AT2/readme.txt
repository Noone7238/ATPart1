Sample commands:
go test

I was able to get the tests to run and pass. I now see having a map is an odd choice as I couldn't control (at least to my knowledge as of now) the order in which the data came back. I tried a 
    solution. I plan on looking at other ways to test and compare values. I didn't stick with one "pass" type condition as I was still trying things out. When I am for a restructure I will dig
    into which is the prefered style. I made this seperate, I'll likely do this again with each major pass, to keep it seperate. I suppose I could use source control (as I'm uploading it to git) 
    but for now, I plan on keeping in piecemeal for easy digging back in to. 

I added some fail states in there, I would imagine if I were aiming for robust unit testing, I'd want to test each possible branch. If you want I can expand upon this, I will likely do that
    When I do the overall restructure and cleanup.

As always, any feed back is greatly appreciated. I plan on potientially looking into some sort of storage or logging as my next exercises. Then moving on from there to something else, maybe the 
    restructure. I imagine that will be the most complex. 


    TODO:
    -inital api                 x
    -testing                    x
    -logging
    -storage
    -code re-architecture
    -error checking/stability