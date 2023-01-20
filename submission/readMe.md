# IOTechSystems Exercises Solution

## Problem statement:
Please write a program in the language of your choice (C or GO prefered) to complete the following tasks:

- Parse the data from exercise-02/data/data.json
- Discard the devices where the timestamp value is before the current time. The timestamps are in UNIX format
- Get the total of all value entries, values are base64 strings containing integer values
- Parse the uuid from the info field of each entry
Output the values total and the list of uuids in the format described by the JSON schema. Write this data to a file

## How to run the solution


1. Open terminal on Linux. Check if Golang is installed at your system:
```
$ which go
```
if this give you an output is some version of golang then you can move ahead with steps. If says  `go not found` just install it from [Go Dev](https://go.dev/doc/install) website based on your operating system.

2. Find the correct directory and unzip the package `submission.zip`
```
$ unzip submission.zip
```
3. There are two folders `data` and `output-schema` with `main.go`. The `main.go` is fetching the data from `data/data.json`. The `output-schema/schema.json` is the required schema based on the ouput is generated.
4. In order to run `main.go` run this command on the current directory on the terminal.
```
$ go run main.go
```
5. This generates a JSON file `output.json` based `output-schema/schema.json`.

## Developer logs
* It was great exercise and really focuses on the developer real world work.
* Learnt new things in GoLang but I think my Java experience helped me alot.
* Based on the assigment I think the job will be realted to IOT based software which is such a big deal right now.
* Developed the code on `go version go1.19.5 darwin/amd64`
* If I had a little more time I could have optimise the code or could have found more efficient ways to solve the problem.
* I really wanted to write tests for the my solution but I was not able to write a proper one so discarded the tests completely.
* I think it satisfy each requirement which was mentioned by the problem statement of Exercise-02