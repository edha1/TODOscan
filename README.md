# 📝 TODOscan

TODOscan is a CLI tool built with **Go** that scans code and finds **TODO** comments. It also integrates with **git blame** and sorts based on the date that the comment was made, to prioritise older TODOS. 

## ⚠️ This Project is a Work in Progress.
> This project is a **work in progress**, I'm working on making the **output in date order as required, fixing any bugs**, adding features, and allowing for the search of other comment types (like FIXME). I have made this project to practise **Go** programming and increase development efficiency. 
---

## ✨ Features

- Recursively scans files in a given path for `TODO` comments (file must be in a initialised git repo). 
- Uses `git blame` to fetch the **commit date** of each TODO.  
- Returns all the **TODO** comments in the codebase. 

---

## 🚀 How to use: 

1. Clone the repo
2. Build the executable: 

```
go build -o todo.exe main.go // on windows 
GOOS=windows GOARCH=amd64 go build -o todo.exe main.go // on Linux/Mac 

```
3. Use the command: .\todo.exe --path <directory> --ext <file-extension> --since <days> to get all TODO comments in line order 

*NOTE:* 
If no paths is given, the default is **.**, which scans all files in the current directory. Default for ext is **.go** files. Default for since is **0 days**.  