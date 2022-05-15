package filelayer

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"lottoengine/model"
	"os"
	"time"
)

var games []model.GameItem

type FileLayer struct {
	File os.File
}

func (f FileLayer) Initialize(workstation string) error {
	games = model.LoadGames()
	WORKSTATION = workstation
	for _, v := range games {
		err := ensureAllPathsExist(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// CurrentFilePath - Current logs of an active game
func CurrentFilePath(gameId string) string {
	return fmt.Sprintf("%v/%v_%v.txt", GAMEPLAYDIR, CURRENT, gameId)
}

// DumpedFilePath - Archived files of past games that have been won or cancelled
func DumpedFilePath(gameId string) string {
	return fmt.Sprintf("%v/%v/%v_%v.log", GAMEPLAYDIR, ARCHIVED, gameId, time.Now().UTC().Local())
}

// FindGame - Find a single Game
func FindGame(id string) model.GameItem {
	_games := make(map[string]model.GameItem)
	for _, v := range games {
		_games[v.Id] = v
	}
	return _games[id]
}

// GamePath
func GamePath(playLog model.GamePlayLog) string {
	return CurrentFilePath(playLog.GameID)
}

// DumpPath
func DumpPath(playLog model.GamePlayLog) string {
	return DumpedFilePath(playLog.GameID)
}

// GameFile
func GameFile(playLog model.GamePlayLog) *os.File {
	var _file *os.File
	_file, err := os.Open(CurrentFilePath(playLog.GameID))
	if err != nil {
		game := FindGame(playLog.GameID)
		ensureAllPathsExist(game)
		_file, err := os.Open(CurrentFilePath(playLog.GameID))
		if err != nil {
			log.Panic(err)
		}
		return _file
	}
	return _file
}

// ensureAllPathsExist - Creates all necessary files and directory for a game
func ensureAllPathsExist(v model.GameItem) error {
	var dir_file_mode = os.ModeDir | (OS_USER_RWX | OS_ALL_R)
	if _, err := os.Stat(GAMEPLAYDIR); os.IsNotExist(err) {
		_err := os.Mkdir(GAMEPLAYDIR, dir_file_mode)
		return _err

	}

	if _, err := os.Stat(ARCHIVEDDIR); os.IsNotExist(err) {
		_err := os.Mkdir(ARCHIVEDDIR, dir_file_mode)
		return _err
	}

	if _, err := os.Stat(CurrentFilePath(v.Id)); os.IsNotExist(err) {
		_, err := os.Create(CurrentFilePath(v.Id))
		return err
	}
	return nil
}

// LogGame
// Writes games to file
// Retturn Game position or an error
func (f FileLayer) LogGame(playLog model.GamePlayLog) (int, error) {
	path := CurrentFilePath(playLog.GameID)
	game := FindGame(playLog.GameID)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	_, err = file.WriteString(playLog.StringifyGame())
	if err != nil {
		log.Panic(err)
		return 0, err
	}
	err = file.Sync()
	if err != nil {
		return 0, err
	}
	position, err := lineCounter(path)
	log.Println(position)
	if err != nil {
		return 0, err
	}
	if position >= game.Target {
		err := f.ArchiveGame(playLog)
		return 0, err
	}
	return position, err
}

// ArchiveCollection - Create an archive for completed game
func (f FileLayer) ArchiveGame(playLog model.GamePlayLog) error {
	gamePath := CurrentFilePath(playLog.GameID)
	dumpPath := DumpedFilePath(playLog.GameID)
	return os.Rename(gamePath, dumpPath)
}

// lineCounter - Position counter
func lineCounter(r string) (int, error) {
	file, err := os.Open(r)
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
