# Course Mixer

[![Build](https://github.com/shoriwe/course-mixer/actions/workflows/build.yaml/badge.svg)](https://github.com/shoriwe/course-mixer/actions/workflows/build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoriwe/course-mixer)](https://goreportcard.com/report/github.com/shoriwe/course-mixer)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shoriwe/course-mixer)

A quickly thrown together tool for trying courses schedules (Linux, Macintosh, Windows).

![image](https://github.com/Woynert/course-mixer/assets/61242172/a0aa3db5-0a5c-4729-ad05-759a09c3cbcf)

## Quick run

### UI

> Godot 3.5.x ([Get it here](https://godotengine.org/download/3.x))

- Open Godot 3.5.x and scan this project's root directory.
- Run the project.
- Click the "Import JSON" button at the top and select your `data.json` file.
- Now you can start course-mixing!

### Web-scraper

> Go 1.19+

You can download pre-compiled binaries [here](releases/latest). Or compile it from source code:

```shell
go install github.com/Woynert/course-mixer@latest
```

Then to generate the JSON for the UI tool use:

```shell
course-mixer courses.json
```
