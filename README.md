# godazo
Stupid simple presenter written in Go.

## Quick start
Install godazo

Runs godazo example with default settings.

On your terminal run `godazo`

On your browser visit [`localhost:8080`](http://localhost:8080/) to see your presentation.

## Useful use
On your terminal run `godazo [-port="PortNumber"] [-media="path-to-presentation-mediafiles"] [-presentation="path-to-json-presentation.json"`]

Example `godazo -port="8081" -presentation="./home/user/devcon/keynote.json"`

That tells godazo to run on port `8081` and do presentation `keynote.json` in user defined path.

On your browser visit [`localhost:8081`](http://localhost:8081/) to see your presentation.

## Thanks

The project is written in Go, and make heavy dependence on Gin and Revealjs.

## Todo

* Don't be evil test your code - Writing tests

* Future Features?