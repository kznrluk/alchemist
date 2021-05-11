package battle

import "github.com/kznrluk/alchemist/app/domain"

type action struct {
	invoker domain.Character
	target  domain.Character

	card            domain.Card
	statusDeltaList []domain.StatusDelta
}

func (a *action) GetActionType() domain.ActionType {
	panic("implement me")
}

func (a *action) GetInvoker() domain.Character {
	panic("implement me")
}

func (a *action) GetCard() domain.Card {
	panic("implement me")
}

func (a *action) SetTarget(target domain.Character) {
	panic("implement me")
}

func (a *action) SetStatusDelta(delta domain.StatusDelta) {
	panic("implement me")
}

func (a *action) GetDeltaList() []domain.StatusDelta {
	panic("implement me")
}

func (a *action) GetResult() domain.ActionResult {
	panic("implement me")
}

func NewAction(invoker domain.Character, card domain.Card) domain.Action {
	return &action{
		invoker: invoker,
		card:    card,
	}
}
