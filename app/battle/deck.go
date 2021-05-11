package battle

import "github.com/kznrluk/alchemist/app/domain"

type deck struct {
}

func (d *deck) Draw(count int) {
	panic("implement me")
}

func (d *deck) FindFromHand(id domain.CardInstanceID) domain.Card {
	panic("implement me")
}

func (d *deck) GetHand() []domain.Card {
	panic("implement me")
}

func (d *deck) AddCard(card domain.Card) {
	panic("implement me")
}

func (d *deck) CardUsed(id domain.CardInstanceID) error {
	panic("implement me")
}

func (d *deck) GetState() domain.CardDeckState {
	panic("implement me")
}

func NewDeck(initCards []domain.CardID) domain.CardDeck {
	return &deck{}
}
