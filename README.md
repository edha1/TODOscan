# 📝 TODOscan

TODOscan is a CLI tool built with **Go** that scans code and finds **TODO** comments. It also integrates with **git blame** and sorts based on the date that the comment was made, to prioritise older TODOS. 

## ⚠️ This Project is a Work in Progress.
> This project is a **work in progress**, I'm working on making the **fixing any bugs, adding features,** and allowing for the search of other comment types (like FIXME). I have made this project to practise **Go** programming and increase development efficiency. 
---

## ✨ Features

- Recursively scans files in a given path for `TODO` comments (file must be in a initialised git repo). 
- Uses `git blame` to fetch the **commit date** of each TODO.  
- Returns all the **TODO** comments in the codebase. 

---

## 🚀 How to use: 

1. Put the .exe in your git repo that you want to find the TODO/FIXME comments in. 
2. Run the command: ./TODOscan [flags] *see below* 

### 🚩 Available Flags

| Flag           | Description                                                                 | Default        |
|----------------|-----------------------------------------------------------------------------|----------------|
| `-path`        | Path to scan for TODOs                                                      | `.` (current directory) |
| `-ext`         | File extension to include (e.g., `.java`, `.go`)                            | `.java`       |
| `-olderthan`   | Only show TODOs older than N days                                           | `0`           |
| `-oldestFirst` | Show results in order of oldest first (true/false)                          | `true`        |
 