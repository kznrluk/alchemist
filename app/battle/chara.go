package battle

import (
	"fmt"
	"github.com/kznrluk/alchemist/app/config"
	"github.com/kznrluk/alchemist/app/domain"
)

type character struct {
	id      domain.CharacterID
	ownerId domain.PlayerID

	isCPU  bool
	deck   domain.CardDeck
	status domain.Statuses

	mix domain.Mix

	reDrawled bool
}

func (c *character) GetCharacterID() domain.CharacterID {
	return c.id
}

func (c *character) IsAlive() bool {
	return c.status[domain.STATUS_HP] > 0
}

func (c *character) StatusChange(delta domain.StatusDelta) {
	c.status[delta.Status] += delta.Delta
}

func (c *character) IsCPU() bool {
	return c.isCPU
}

func (c *character) DrawDeck() (domain.CardDeckState, error) {
	if c.reDrawled {
		return c.deck.GetState(), fmt.Errorf("battle: すでに一度デッキを引き直しています")
	}

	c.deck.Draw(config.DRAW_COUNT_AT_ONCE)
	c.reDrawled = true

	return c.deck.GetState(), nil
}

func (c *character) CreateMixCase() []domain.CardMixCase {
	return c.mix.FindMix(c.deck.GetHand())
}

func (c *character) DoMix(cardList []domain.Card) error {
	cardMixCase, err := c.mix.ExecMix(cardList)
	if err != nil {
		return fmt.Errorf("battle: カードのミックスができませんでした %w", err)
	}

	for _, mc := range cardMixCase.From {
		err = c.deck.CardUsed(mc.GetCardInstanceId())
		if err != nil {
			return fmt.Errorf("battle: 使用できないカードが使用されました %w", err)
		}
	}

	c.deck.AddCard(NewCard(cardMixCase.To, cardMixCase.SuccessRate))
	// TODO: 他選択肢の削除(メモリリーク解消)

	return nil
}

func (c *character) UseCard(cardInstanceId domain.CardInstanceID, isDryRun bool) (domain.Action, error) {
	cd := c.deck.FindFromHand(cardInstanceId)
	if cd == nil {
		return nil, fmt.Errorf("battle: カードが手札に見つかりませんでした")
	}

	if !isDryRun {
		c.reDrawled = false

		err := c.deck.CardUsed(cd.GetCardInstanceId())
		if err != nil {
			return nil, fmt.Errorf("battle: 使用できないカードが使用されました %w", err)
		}
	}

	return NewAction(c, cd), nil
}

func (c *character) GetState() domain.CharacterState {
	return domain.CharacterState{
		CharacterID: c.id,
		PlayerID:    c.ownerId,
		Status:      c.status,
		Deck:        c.deck.GetState(),
	}
}

func NewCharacter(cid domain.CharacterInitializeData) (domain.Character, error) {
	return &character{
		id:        cid.CharacterID,
		ownerId:   cid.PlayerID,
		isCPU:     cid.PlayerID == 0,
		deck:      NewDeck(cid.DeckCardIdList),
		status:    map[domain.Status]domain.StatusValue{},
		mix:       NewMix(),
		reDrawled: false,
	}, nil
}
