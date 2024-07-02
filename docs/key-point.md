# Golang key point

## Package

- executable package: build out an executable, and must have a main func.

- reusable package:

## Data types

### Basic Type

bool, string, int, float, byte

Array, Slice

### struct：

struct with pointers: := &struct, define pointer receiver (name \*struct)

Refrence Type:int, float, string,bool, struct
Value Type: pointer, slice, map, channel,func

### Maps

- key must be same type, value must be same type

* vs struct: can iterate, no need predifine all the keys,

### Interfaces

- interface are implicit, the struct with same method are treaded implemented the interface. interface can embded each other

## Concurrency & Parallelism

- go {func}: start a new go routine.
- channel used to share value between routines, `{name} chan {type}` defined a channel variable, `{name} <- {value};` set a value, `<-{name}` get a value, we can ``range {channel}` to wait channel give back a channel

## Philosiphy

- Go lang has no class, not OO approach, we need to define new type and apply func to the new type as a receiver, like: func ({type}) name(){}
