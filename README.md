# TODOscan

TODOscan is a CLI tool built with **Go** that scans code and finds **TODO** comments. It also integrates with **git blame** and sorts based on the date that the comment was made, to prioritise older TODOS. 

## ⚠️ This prject is a Work in Progress.   
> This project is a work in progress, I'm working on fixing any bugs, adding features (such as filtering by date), and allowing for the search of other comment types (like FIXME). I have made this project to practise **Go** programming and increase development efficiency for developers.

---

## Features

- Recursively scans files in a given path for `TODO` comments.  
- Uses `git blame` to fetch the **commit date** of each TODO.  
- Sorts TODOs by **oldest first** to help prioritise long-standing tasks.  

---
