# goDo

Beginner Go CLI project based off of the tutorial from @SirTingling in this [post](https://dev.to/jordan_t/simple-go-cli-todo-app-19j6).

## v1
Maintained much of the original design and implementation, including the storage of tasks as a JSON file in local storage.
Small bash script compiles program, then moves and renames it to a directory that is in the Path. Allows running the program easily.

### Usage
The `compile.sh` script can be used to build the go app into an executable. It will need to be edited to meet your needs. Currently it is configured to send the executable to an existing path in the environment variables.
Once the script has finished running, simply run:
```
godo -help
```
to see the current list of flags.