package domain

type Status string

const (
	STATUS_HP Status = "HP"
)

type StatusValue float32

type Statuses map[Status]StatusValue

type StatusDelta struct {
	Status Status
	Delta  StatusValue
}
