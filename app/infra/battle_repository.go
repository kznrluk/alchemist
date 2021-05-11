package infra

import (
	"fmt"
	"github.com/kznrluk/alchemist/app/domain"
)

// TODO: インメモリだとスケールしないのでRedisとかに移設する
var battleRepoSingleton domain.BattleRepository

type inMemoryBattleRepo struct {
	battles map[domain.BattleID]domain.Battle
}

func (i *inMemoryBattleRepo) GetFromBattleID(id domain.BattleID) (domain.Battle, error) {
	battle, ok := i.battles[id]
	if !ok {
		return nil, fmt.Errorf("infra: バトルがメモリ上に存在しませんでした")
	}

	return battle, nil
}

func (i *inMemoryBattleRepo) Save(battle domain.Battle) error {
	i.battles[battle.GetBattleID()] = battle

	return nil
}

func NewInMemoryBattleRepo() domain.BattleRepository {
	if battleRepoSingleton == nil {
		battleRepoSingleton = &inMemoryBattleRepo{
			battles: map[domain.BattleID]domain.Battle{},
		}
	}
	return battleRepoSingleton
}
