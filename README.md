# Awaluddin - parkinglot

## Installation

1. install golang (ref: [https://golang.org/doc/install](https://golang.org/doc/install))
2. clone repository `git clone git@github.com:awaluddin/parkinglot.git somewhere`

## Running
1. Running `./parking_lot.go` 

## keys
1. Create car slot `create <number of slot>` `-> create 5`
2. Add car to park slot `park <Car ID> <Car Color>` `-> park B12345VB white`
3. Car leave or remove car from slot `leave <slot number>` `-> leave 2`
4. Show slot list `status` `-> status`
5. Find cars by ID `find id <Car ID>` `-> find id B12345VB`
6. Find cars by Color `find color <Car Color>` `-> find color white`
