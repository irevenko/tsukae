# tsukae 🧑‍💻📊

> Tsukae, 使え - means <i>use</i> in Japanese (so it refers to commands that you <i>use</i>)
<p align="center"><img src="tsukae.gif"></p>

<p align="center">Built on top of <a href="https://github.com/gizak/termui">termui</a> and <a href="https://github.com/spf13/cobra">cobra</a> </p> <br>
<p align="center">Big shoutout to <a href="https://gitlab.com/jokerj40">jokerj40</a> for suggesting this project idea (the original thought was to parse the history file) </p>


# Usage 🔬
```<SHELL_NAME>``` - <b>zsh</b> or <b>bash</b>

### Flags 🎌
```-p, --piechart``` - use with ```<SHELL_NAME>``` (draws only piechart widget) <br>
```-b, --barchart``` - use with ```<SHELL_NAME>``` (draws only barchart widget) <br>
```-l, --list``` - use with ```<SHELL_NAME>``` (draws only list widget)

### Draw all widgets
```tsukae <SHELL_NAME>``` <br>
<img src="tsukae.gif" width="300"> 

### Draw specific widget
```tsukae <SHELL_NAME> -p ``` <br>
<img src="https://i.imgur.com/FsXDwsC.jpg" width="300"> <br>
```tsukae <SHELL_NAME> -b```  <br>
<img src="https://i.imgur.com/O9q0SVx.jpg" width="500"> <br>
```tsukae <SHELL_NAME> -l``` <br>
<img src="https://i.imgur.com/FEgw8wp.jpg" width="300">

### Pass certain commands number (1-15) and draw all widgets
```tsukae <SHELL_NAME> 10```

### Draw specific widget with certain commands number
```tsukae <SHELL_NAME> 10 -p ``` <br>
```tsukae <SHELL_NAME> 5 -b``` <br> 
```tsukae <SHELL_NAME> 3 -l``` 

# Contributing 🤝
Contributions, issues and feature requests are welcome! 👍 <br>
Feel free to check [open issues](https://github.com/irevenko/tsukae/issues).

# Quick Start 🚀
```git clone https://github.com/irevenko/tsukae.git``` <br>
```cd tsukae``` <br>
```go get -d ./...``` <br>
```go run main.go``` <br>

# What I Learned 🧠
- Parsing Text Files using Go
- Drawing termui

# ToDo
- auto detect shell, manual shell path
- maybe fish, sh, etc... support
- make proper PieChart render
- binaries for osx, linux

# License 📑 
(c) 2021 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)
