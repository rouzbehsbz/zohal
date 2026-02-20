package world

import "time"

type System interface {
	Stage() Stage
	Update(w *World, dt time.Duration)
}
