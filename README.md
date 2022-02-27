# ATPart1


02.21.2022 comments:



Please start working using PR requests. \
Look at   [GitHub flow - GitHub Docs](https://docs.github.com/en/get-started/quickstart/github-flow),
[Gitflow Workflow | Atlassian Git Tutorial](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) \
we are using gitlab flow https://docs.gitlab.com/ee/topics/gitlab_flow.html but I suggest to start with github flow cause its more straight forward.
Our primary expectation right now to start using PRs so I can leave comments for your PRs.

Right now I'll write down all comments just in the file:

**First**, lets start from directories:
[GitHub - golang-standards/project-layout: Standard Go Project Layout](https://github.com/golang-standards/project-layout#go-directories)	, right now we can use just this  three: `cmd`, `internal`, `pkg`. Please recognize your structure based on the explanation above.

**After**, lets go to the main.go:
https://github.com/Noone7238/ATPart1/blob/main/AT3/cmd/main.go

Huge problem - we are not checking if server been started properly or not, i.e. `network.Start()` does not return anything. \
Couple medium problems: \
- we should have a config folder and by using flag package we should be able to set up port and host for the http server
- we should create a logger instance (we are using rs/zerolog) to write down logs if something goes wrong
- network/network.go probably should be renamed to  `app/app.go` 

**Next**, lets talk about `network/`. Lets make structure more verbose and easy to read. Ideally we want to have those files \
`routes.go` - with all routes
`handlers.go`- with handlers
`app.go` - a compositor which get envs using flags, creates logger instance, invokes routes, handlers and etc. 

**Lets move** to the `product.go`.  My first idea, why don't we:
- move it to the separate folder, directly it has nothing to do with network
- Lets split this file into few ones, like structures.go, implementation.go
- Sample very bad naming. I think we should go with something like ToDo, or ItemHandler. Again to make it more easy to read for other engineers.

Couple serious issues: \
[ATPart1/products.go at main · Noone7238/ATPart1 · GitHub](https://github.com/Noone7238/ATPart1/blob/main/AT3/network/products.go#L59) its a bad habit do not report  about golang errors. SampleHandler should have an logger instance and in 99.995% cases you want to log this error. You can play with a log level and it might be a warn or even just a debug message. But good logs is a foundation of future maintainability of the code base.

[ATPart1/products.go at main · Noone7238/ATPart1 · GitHub](https://github.com/Noone7238/ATPart1/blob/main/AT3/network/products.go#L104) fmt.Prints should not go the production. I do really recommend to set up golang linter on save and on git commit events so you can catch them before pushing to production [Lint on Save With VS Code Official Golang Extension - Qvault](https://qvault.io/golang/lint-on-save-vs-code-golang/)

**Questions:**

[ATPart1/products.go at 2724c6f3f325778acb43a30e47b1f2d357c5853e · Noone7238/ATPart1 · GitHub](https://github.com/Noone7238/ATPart1/blob/2724c6f3f325778acb43a30e47b1f2d357c5853e/AT3/network/products.go#L115) I cannot figure out why you need to get `i` here at all? And it looks to me its to make this part more simple [ATPart1/products.go at 2724c6f3f325778acb43a30e47b1f2d357c5853e · Noone7238/ATPart1 · GitHub](https://github.com/Noone7238/ATPart1/blob/2724c6f3f325778acb43a30e47b1f2d357c5853e/AT3/network/products.go#L122) Why cannot we have samples just as an array? (which will crash whole app completely) \
Please take a look on golang documentation for arrays, slices in go and as I mention it looks like we can rewrite this part to be more easy and robust
