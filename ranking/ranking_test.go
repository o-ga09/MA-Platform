package ranking

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestRanking(t *testing.T) {
	entryfile := "entry_user.csv"
	playfile := "play_data.csv"
	status := Ranking(entryfile,playfile)
	expected := 0
	assert.Equal(t,status,expected)
}

func TestReadData(t *testing.T) {
	entryfile := "entry_user.csv"
	playfile := "play_data.csv"
	entry_users, play_logs, err := ReadData(entryfile,playfile)
	
	expectedEntryUsers := []Entry_data{
		{handle_name: "testUser1",player_id: "aaa001"},
		{handle_name: "testUser2",player_id: "aaa002"},
	}
	expectedPlayLogs := []Play_data{
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "aaa001",score: 100},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "aaa002",score: 200},
	}
	for i, data := range *entry_users {
		assert.Equal(t,data.handle_name,expectedEntryUsers[i].handle_name)
		assert.Equal(t,data.player_id,expectedEntryUsers[i].player_id)
	}
	for i, data := range *play_logs {
		assert.Equal(t,data.playedAt,expectedPlayLogs[i].playedAt)
		assert.Equal(t,data.player_id,expectedPlayLogs[i].player_id)
		assert.Equal(t,data.score,expectedPlayLogs[i].score)
	}
	assert.Equal(t,err,nil)
}

func TestSerchEntryUser(t *testing.T) {
	entry_user := []Entry_data{
		{player_id: "testUser1",handle_name: "aaa001"},
		{player_id: "testUser2",handle_name: "aaa002"},
	}
	play_logs := []Play_data{
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser1",score: 100},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser1",score: 300},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser1",score: 200},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser2",score: 500},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser2",score: 300},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser2",score: 200},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser3",score: 300},
		{playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC),player_id: "testUser4",score: 400},
	}
	EntryPlayer_logs,err := SearchEntryUser(&entry_user,&play_logs)

	expectedEntryplayer := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 300,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa002",score: 500,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}
	for i, data := range *EntryPlayer_logs {
		assert.Equal(t,data.handle_name,expectedEntryplayer[i].handle_name)
		assert.Equal(t,data.player_id,expectedEntryplayer[i].player_id)
		assert.Equal(t,data.score,expectedEntryplayer[i].score)
		assert.Equal(t,data.rank,expectedEntryplayer[i].rank)
		assert.Equal(t,data.player_id,expectedEntryplayer[i].player_id)
	}
	assert.Equal(t,len(*EntryPlayer_logs),len(expectedEntryplayer))
	assert.Equal(t,err,nil)
}

func TestSortedByScorePlayer(t *testing.T) {
	entryPlayerLogs := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 1000,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa002",score: 400,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser3",handle_name: "aaa003",score: 600,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser4",handle_name: "aaa004",score: 700,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser5",handle_name: "aaa005",score: 300,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser6",handle_name: "aaa006",score: 200,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}
	SortedByScorePlayer(&entryPlayerLogs)

	expectedRankedPlayer := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 1000,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser4",handle_name: "aaa004",score: 700,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser3",handle_name: "aaa003",score: 600,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa002",score: 400,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser5",handle_name: "aaa005",score: 300,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser6",handle_name: "aaa006",score: 200,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}
	for i, data := range entryPlayerLogs {
		assert.Equal(t,data.player_id,expectedRankedPlayer[i].player_id)
		assert.Equal(t,data.handle_name,expectedRankedPlayer[i].handle_name)
		assert.Equal(t,data.rank,expectedRankedPlayer[i].rank)
		assert.Equal(t,data.score,expectedRankedPlayer[i].score)
		assert.Equal(t,data.playedAt,expectedRankedPlayer[i].playedAt)
	}
}

func TestRankingPlayer(t *testing.T) {
	sortedPlayer := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 1000,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa004",score: 900,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser3",handle_name: "aaa003",score: 800,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser4",handle_name: "aaa002",score: 700,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser5",handle_name: "aaa005",score: 600,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser6",handle_name: "aaa006",score: 500,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser7",handle_name: "aaa004",score: 400,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser8",handle_name: "aaa003",score: 300,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser9",handle_name: "aaa002",score: 200,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser10",handle_name: "aaa005",score: 100,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser11",handle_name: "aaa006",score: 50,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}
	err := RankingPlayer(&sortedPlayer)

	expectedSortedPlayer := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 1000,rank: 1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa004",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser3",handle_name: "aaa003",score: 800,rank: 3,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser4",handle_name: "aaa002",score: 700,rank: 4,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser5",handle_name: "aaa005",score: 600,rank: 5,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser6",handle_name: "aaa006",score: 500,rank: 6,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser7",handle_name: "aaa004",score: 400,rank: 7,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser8",handle_name: "aaa003",score: 300,rank: 8,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser9",handle_name: "aaa002",score: 200,rank: 9,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser10",handle_name: "aaa005",score: 100,rank: 10,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser11",handle_name: "aaa006",score: 50,rank: -1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}
	for i, data := range sortedPlayer {
		assert.Equal(t,data.handle_name,expectedSortedPlayer[i].handle_name)
		assert.Equal(t,data.player_id,expectedSortedPlayer[i].player_id)
		assert.Equal(t,data.rank,expectedSortedPlayer[i].rank)
		assert.Equal(t,data.score,expectedSortedPlayer[i].score)
		assert.Equal(t,data.playedAt,expectedSortedPlayer[i].playedAt)
	}
	assert.Equal(t,err,nil)
}

func TestSortedByIdePlayer(t *testing.T) {
	ranledPlayer := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 1000,rank: 1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser4",handle_name: "aaa004",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa003",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser3",handle_name: "aaa002",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser5",handle_name: "aaa005",score: 600,rank: 3,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser6",handle_name: "aaa006",score: 500,rank: 4,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser7",handle_name: "aaa004",score: 400,rank: 5,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser8",handle_name: "aaa003",score: 300,rank: 6,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser9",handle_name: "aaa002",score: 200,rank: 7,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser10",handle_name: "aaa005",score: 100,rank: 8,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser11",handle_name: "aaa006",score: 50,rank: 9,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}

	SortedByIdPlayer(&ranledPlayer)

	expectedRankedPlayer := []EntryUser_logs{
		{player_id: "testUser1",handle_name: "aaa001",score: 1000,rank: 1,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser2",handle_name: "aaa003",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser3",handle_name: "aaa002",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser4",handle_name: "aaa004",score: 900,rank: 2,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser5",handle_name: "aaa005",score: 600,rank: 3,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser6",handle_name: "aaa006",score: 500,rank: 4,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser7",handle_name: "aaa004",score: 400,rank: 5,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser8",handle_name: "aaa003",score: 300,rank: 6,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser9",handle_name: "aaa002",score: 200,rank: 7,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser10",handle_name: "aaa005",score: 100,rank: 8,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
		{player_id: "testUser11",handle_name: "aaa006",score: 50,rank: 9,playedAt: time.Date(2023, time.August, 5, 12, 30, 0, 0, time.UTC)},
	}

	for i, data := range ranledPlayer {
		assert.Equal(t,data.handle_name,expectedRankedPlayer[i].handle_name)
		assert.Equal(t,data.player_id,expectedRankedPlayer[i].player_id)
		assert.Equal(t,data.score,expectedRankedPlayer[i].score)
		assert.Equal(t,data.rank,expectedRankedPlayer[i].rank)
		assert.Equal(t,data.playedAt,expectedRankedPlayer[i].playedAt)
	}
}