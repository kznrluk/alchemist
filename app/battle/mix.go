package battle

import "github.com/kznrluk/alchemist/app/domain"

type mix struct {
}

func (m *mix) FindMix(cardList []domain.Card) []domain.CardMixCase {
	panic("implement me")
}

func (m *mix) ExecMix(cardList []domain.Card) (domain.CardMixCase, error) {
	panic("implement me")
}

func NewMix() domain.Mix {
	return &mix{}
}
