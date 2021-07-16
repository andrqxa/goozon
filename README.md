# Ozon Go school

## 1. Compiling project
There is MakeFile, so all of you need is just execute `make`

## 2. Running tests
There is MakeFile, so all of you need is just execute `make test`

## 3. Start application
After compiling the application (see #1) you can start the program e.g. `./schedule.exe` at Windows There are 2 options for starting application:

- `./schedule.exe` then you enter Start Date and End Date of each automobile in format `DD.MM.YYYY HH:MM:SS; DD.MM.YYYY HH:MM:SS` plus "Enter" When all data were added you should enter "Ctrl-Z"

- `./schedule.exe path-to-file-with-data` File with data must exist and has data for each automobile in the format: `DD.MM.YYYY HH:MM:SS; DD.MM.YYYY HH:MM:SS`

In the end, you get information about the maximum quantity of automobiles 

`The maximum number of vehicles at the same time in the parking lot equals` plus number

## 4. Aditional dependencies
To compile a project on Windows you need to install make choco install make
To run tests on Windows you should install gcc: for example, TDM