# godazo
Stupid simple presenter written in Go.

## Quick start
Install godazo 

`$ go get github.com/rutaihwa/godazo`

Run godazo example with default settings

`$ godazo`

On your browser visit [`localhost:8080`](http://localhost:8080/) to see your presentation.

## Useful use
On your terminal run `godazo [-port="PortNumber"] [-media="path-to-presentation-mediafiles"] [-presentation="path-to-json-presentation.json"`]

Example `$ godazo -port="8081" -presentation="./home/user/devcon/keynote.json"`

That tells godazo to run on port `8081` and do presentation `keynote.json` in user defined path.

On your browser visit [`localhost:8081`](http://localhost:8081/) to see your presentation.

## Thanks

The project is written in Go, and make heavy dependence on [Gin](https://github.com/gin-gonic/gin) and [Revealjs](https://github.com/hakimel/reveal.js/).

## Todo

* Don't be evil test your code - Writing tests

* Future Features?