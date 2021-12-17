package main

import (
	"log"
	"strings"
)

type Worker int
type Input struct {
	Text       string
	WordToGrep string
}

func (w *Worker) Map(in Input, reply *string) error {

	lines := strings.Split(in.Text, "\n")
	count := 0
	for _, v := range lines {
		if strings.Count(v, in.WordToGrep) != 0 {
			*reply += v + "\n"
			count++ // debug
		}
	}
	if count == 0 {
		log.Printf("Word not founded")
	} else {
		log.Printf("Founded {%s} in %d lines", in.WordToGrep, count) //debug
	}
	*reply = strings.TrimSuffix(*reply, "\n")

	return nil
}
