# GotHub

## What is it?

GotHub is a library for interfacing with GitHub, for the Go programming
language.

## How do I pronounce the name?

You can pronounce it like "got hub", "goth hub" for those darker times, or
really, however the heck you feel like pronouncing it.

"go thub" is also acceptable.

## Why are you working on this?

Well, simply because I wanted a client-side library with which to utilize
GitHub's APIs, in a similar fashion to the
[PyGithub](https://github.com/jacquev6/PyGithub) library for Python.

## How do I get it?

A simple `go get github.com/nesv/gothub` should be sufficient.

## I can haz docs plz?

Documentation is automatically generated, and can be found here:
[http://godoc.org/github.com/nesv/gothub](http://godoc.org/github.com/nesv/gothub)

kthxbai.

## Can I use this within my company's proprietary software?

You betcha. Check out the LICENSE.md file &mdash; the code and documentation for this project is 
kept under the MIT license. In other words, feel free to copy it, sell it, do whatever you want,
just make sure to include the provided LICENSE.md file if you incorporate it into your stuff (or
your company's stuff).

## Look, I need to run some tests on this to verify it even works...

And you can! No Go project can really be "complete" without tests.

To run the tests, you need to set two environment variables: `GITHUB_USERNAME` and
`GITHUB_PASSWORD`. Here is how you can run the tests without having to save these trinkets of
information anywhere:

    $ GITHUB_USERNAME=<login> GITHUB_PASSWORD=<password> go test -v

Afterwards, be sure to delete your shell's history file, so that your credentials aren't laying
around anywhere.

## I wanna help

Help is always appreciated!

If you have a bug or a feature request, please create an issue here in GitHub.

All of the development tasks are planned out in [Trello](https://trello.com). If you would like
to keep track of what is being worked on, check out the
[GotHub Development](https://trello.com/board/gothub-development/513689261eb6779e35005e29) board.
The tasks in Trello are broken down much further than they are in GitHub.

Once a feature, or bug, is completed in Trello, we will update the GitHub issue accordingly.

### …I actually wanna write code

Even better! Sign up for an account on Trello and assign yourself to a card so the work can be
appropriately coordinated, then fork the repository. After your code is done, tests have passed
and all that good stuff, submit a pull request and your code will be merged in.
