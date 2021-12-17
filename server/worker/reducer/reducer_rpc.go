package main

import "log"

type Reducer int

func (r *Reducer) Reduce(in string, reply *string) error {
	log.Printf("Reducer received %s", in)
	*reply = in
	return nil
}
