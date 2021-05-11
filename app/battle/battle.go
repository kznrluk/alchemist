package battle

import (
	"fmt"
	"github.com/kznrluk/alchemist/app/domain"
)

type battle struct {
	alpha                          []domain.Character
	beta                           []domain.Character
	currentControllableCharacterID domain.CharacterID
}

func (b *battle) GetBattleID() domain.BattleID {
	panic("implement me")
}

func (b *battle) CurrentControllableCharacterID() domain.CharacterID {
	panic("implement me")
}

func (b *battle) Mix(invoker domain.CharacterID, cardIDList []domain.CardInstanceID) (domain.BattleState, error) {
	panic("implement me")
}

func (b *battle) UseCard(invoker domain.CharacterID, cardID domain.CardInstanceID, target domain.CharacterID) (domain.BattleState, error) {
	invokeChara := b.findCharacter(invoker)
	targetChara := b.findCharacter(target)
	if invokeChara == nil && targetChara == nil {
		return domain.BattleState{}, fmt.Errorf("battle: 要求されたキャラクターが存在しません")
	}

	act, err := invokeChara.UseCard(cardID, false)
	if err != nil {
		return domain.BattleState{}, fmt.Errorf("battle: キャラクターカードの使用に失敗 %w", err)
	}

	act.SetTarget(targetChara)
	for _, delta := range act.GetDeltaList() {
		targetChara.StatusChange(delta)
	}

	return b.GetState(), nil
}

func (b *battle) TurnEnd(invoker domain.CharacterID) error {
	chara := b.findCharacter(invoker)
	if chara == nil {
		return fmt.Errorf("battle: 要求されたキャラクターが存在しません")
	}

	if chara.GetCharacterID() == invoker {
		return fmt.Errorf("battle: 予期しないユーザからのターンエンド指示")
	}

	return nil
}

func (b *battle) findCharacter(id domain.CharacterID) domain.Character {
	for _, c := range b.alpha {
		if c.GetCharacterID() == id {
			return c
		}
	}

	for _, c := range b.beta {
		if c.GetCharacterID() == id {
			return c
		}
	}

	return nil
}

func (b *battle) GetState() domain.BattleState {
	var alphaState []domain.CharacterState
	var betaState []domain.CharacterState

	for _, c := range b.alpha {
		alphaState = append(alphaState, c.GetState())
	}

	for _, c := range b.beta {
		betaState = append(betaState, c.GetState())
	}

	return domain.BattleState{
		Alpha:                          alphaState,
		Beta:                           betaState,
		CurrentControllableCharacterID: b.currentControllableCharacterID,
		CurrentTurn:                    0,
	}
}

func InitializeNewBattle(alpha []domain.CharacterInitializeData, beta []domain.CharacterInitializeData) (domain.Battle, error) {
	var alphaChara []domain.Character
	var betaChara []domain.Character
	var currentControllable domain.CharacterID

	for _, cid := range alpha {
		chara, err := NewCharacter(cid)
		if err != nil {
			return nil, fmt.Errorf("battle: キャラクタインスタンスの生成に失敗 %w", err)
		}

		if currentControllable == "" {
			currentControllable = chara.GetCharacterID()
		}

		alphaChara = append(alphaChara, chara)
	}

	for _, cid := range beta {
		chara, err := NewCharacter(cid)
		if err != nil {
			return nil, fmt.Errorf("battle: キャラクタインスタンスの生成に失敗 %w", err)
		}
		betaChara = append(betaChara, chara)
	}

	return &battle{
		alpha:                          alphaChara,
		beta:                           betaChara,
		currentControllableCharacterID: currentControllable,
	}, nil
}
