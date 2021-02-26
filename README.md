# tsukae 🧑‍💻📊

> Tsukae, 使え - means <i>use</i> in Japanese (so it refers to commands that you <i>use</i>)
<p align="center"><img src="tsukae.gif"></p>


# Usage 🔬
<SHELL_NAME> - <b>zsh</b> or <b>bash</b>

### Flags
```-p, --piechart``` - use with <SHELL_NAME> (draws piechart only) <br>
```-b, --barchart``` - use with <SHELL_NAME> (draws barchart only) <br>
```-l, --list``` - use with <SHELL_NAME> (draws list only)

### Draw all widgets
```tsukae <SHELL_NAME>```
 
### Enter specific commands number (1-15) and draw all widgets
```tsukae <SHELL_NAME> 10```

### Draw specific widget
```tsukae <SHELL_NAME> -p ``` <br>
```tsukae <SHELL_NAME> -b```  <br>
```tsukae <SHELL_NAME> -l```

### Draw specific widget with commands number
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
- binaries for osx, linux

# License 📑 
(c) 2021 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)
