package model

import (
	"fmt"
)

type GameItem struct {
	Id     string
	Title  string
	Amount float64
	Target int
}

type GamePlayLog struct {
	Id        string
	GameID    string
	UserId    string
	UserEntry string
	Amount    float64
	PlayTime  string
}

func (playLog GamePlayLog) StringifyGame() string {
	var result = fmt.Sprintf("%v,%v,%v,%v,%v \n", playLog.PlayTime, playLog.GameID, playLog.UserId, playLog.UserEntry, playLog.Amount)
	return result
}



// Test collection
func LoadGames() []GameItem {
	return []GameItem{
		{
			Id:     "1",
			Title:  "xxxx",
			Amount: 10,
			Target: 100,
		},{
			Id:     "2",
			Title:  "yyyy",
			Amount: 20,
			Target: 100,
		},
		{
			Id:     "3",
			Title:  "zzzz",
			Amount: 30,
			Target: 100,
		},
	}
}