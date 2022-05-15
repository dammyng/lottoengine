package filelayer

import (
	"lottoengine/model"
	"os"
	"testing"
	"time"
)

var tester FileLayer

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}

func setup() {
	station :=   os.Getenv("WORKSTATION")
	tester.Initialize(station)
}

func TestFileLayer_Write(t *testing.T) {
	type fields struct {
		File os.File
	}
	type args struct {
		game model.GamePlayLog
	}
	play1 := model.GamePlayLog{
		Id:        "genId()",
		GameID:    "1",
		UserId:    "user-Id",
		UserEntry: "XX 56",
		Amount:    10,
		PlayTime:  time.Now().Local().String(),
	}
	play3 := model.GamePlayLog{
		Id:        "genId()",
		GameID:    "3",
		UserId:    "user-Id",
		UserEntry: "XX 56",
		Amount:    10,
		PlayTime:  time.Now().Local().String(),
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test 1",
			fields:  fields{File: *GameFile(play1)},
			args:    args{play1},
			want:    0,
			wantErr: false,
		},
		{
			name:    "test 1",
			fields:  fields{File: *GameFile(play3)},
			args:    args{play3},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FileLayer{
				File: tt.fields.File,
			}
			got, err := f.LogGame(tt.args.game)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileLayer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileLayer.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
