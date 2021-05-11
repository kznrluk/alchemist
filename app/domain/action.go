package domain

type ActionType string

const (
	ACTION_TYPE_ATTACK = "ATTACK"
)

type ActionResult struct {
	ActionType            ActionType
	Invoker               CharacterID
	Target                CharacterID
	TargetStatusDeltaList []StatusDelta
}

type Action interface {
	GetActionType() ActionType
	GetInvoker() Character
	GetCard() Card

	SetTarget(target Character)
	SetStatusDelta(delta StatusDelta)
	GetDeltaList() []StatusDelta

	GetResult() ActionResult
}
