[![GoDoc](https://godoc.org/github.com/thomaso-mirodin/go-shorten?status.svg)](http://godoc.org/github.com/thomaso-mirodin/go-shorten)
[![Go Report Card](https://goreportcard.com/badge/github.com/thomaso-mirodin/go-shorten)](https://goreportcard.com/report/github.com/thomaso-mirodin/go-shorten)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/thomaso-mirodin/go-shorten/master/LICENSE)

## go-shorten: URL Shortener Service

This service stores and serves URL redirects. I stole the idea from when I worked at LinkedIn who apparently stole it from Google.

### Why?

By itself, a shared URL shortening service is stupid useful. Want to visit a thing? Type "go/thing" into your browser and hit enter.

It is honestly little things like this make everyone's life far easier.

### How do people use this?

Where I work we generally follow a pattern of `go/$thing` should take you to whatever $thing is. Here are a few examples:
* Jira is at `go/jira`
* Github is at `go/gh`, `go/github`, `go/g`, etc
* Our monitoring dashboards are at `go/graphs`, `go/dashboards`, etc

It is perfectly fine to have duplication, the goal is that all the different ways folks think will all be taken to the right place.

You can also do fancier things like:
* `go/pr/1234` or `go/#1234` will take you directly to the PR #1234
* `go/jira/ABC-1234`, `go/ABC-1234` to go directly to a specific jira issue

### Okay, how do I set this up?

Roughly, to make this work:
1. Build and host go-shorten somewhere
2. Setup a DNS entry to point to it (e.g. `go.corp.example.com`)
3. Configure any clients you have to include `corp.example.com` in their DNS search suffix list
4. Troubleshoot :P

## Credits

I forked this project from <https://github.com/didip/shawty> because I liked how they laid out their project but I wanted to add a bunch more features and productionize it a bit more than was within scope for the original project.
