# Course Mixer
A quickly thrown together tool for trying courses schedules. (Linux, Macintosh, Windows)

![image](https://github.com/Woynert/course-mixer/assets/61242172/a0aa3db5-0a5c-4729-ad05-759a09c3cbcf)

## Requirements
- Godot 3.5.x
- Go

## Usage (Rough)
1. Input data:
    - Select a topic at [HorariosUPB](https://horariosupb.bucaramanga.upb.edu.co/) and download the html file by right click > "save page as...".
    - Place the html file inside the `go-crawler/datain` directory.
    - Repeat this process for as many topics as you want.
2. Output data:
    - Edit `go-crawler/main.go` and edit the "queries" variable replacing the **file name (File)** and **column number (Cols)** according to your html.
    - Run the go program with `go run .`
    - A file called `data.json` should be created under `go-crawler/dataout`.
3. Graphic interface:
    - Open Godot 3.5.x and scan the root project directory.
    - Run the project.
    - Click the "Import JSON" button and select your `data.json`.
    - Now you can start course-mixing!
