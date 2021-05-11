package battle

import "github.com/kznrluk/alchemist/app/domain"

type card struct{}

func (c *card) GetCardInstanceId() domain.CardInstanceID {
	panic("implement me")
}

func (c *card) GetCurrentPower() {
	panic("implement me")
}

func (c *card) GetState() domain.CardState {
	panic("implement me")
}

func NewCard(definition domain.CardDefinition, successRate int) domain.Card {
	return &card{}
}
