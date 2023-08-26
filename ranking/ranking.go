package ranking

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

type Entry_data struct {
	player_id string
	handle_name string
}

type Play_data struct {
	playedAt time.Time
	player_id string
	score int
}

type EntryUser_logs struct {
	rank int
	player_id string
	score int
	handle_name string
	playedAt time.Time
}

type ByScore []EntryUser_logs

func(l *ByScore) Len() int {return len(*l)}
func(l *ByScore) Swap(i,j int) {(*l)[i], (*l)[j] = (*l)[j], (*l)[i]}
func(l *ByScore) Less(i,j int) bool {return (*l)[i].score > (*l)[j].score}

type ById []EntryUser_logs

func(l *ById) Len() int {return len(*l)}
func(l *ById) Swap(i,j int) {(*l)[i], (*l)[j] = (*l)[j], (*l)[i]}
func(l *ById) Less(i,j int) bool {return (*l)[i].player_id < (*l)[j].player_id}

func Ranking(entryfile string, playLogfile string) int {
	entry_data, play_log, err := ReadData(entryfile,playLogfile)
	if err != nil {
		fmt.Fprintf(os.Stderr,"can not read file: %v",err)
		return 1
	}

	entryUserLog,err := SearchEntryUser(entry_data,play_log)

	SortedByScorePlayer(entryUserLog)

	err = RankingPlayer(entryUserLog)
	if err != nil {
		fmt.Fprintf(os.Stderr,"can not rank Player %v",err)
		return 2
	}

	SortedByIdPlayer(entryUserLog)

	fmt.Printf("rank, player_id, handle_name, score\n")
	for _, record := range *entryUserLog {
		fmt.Printf("%d, %s, %s, %d\n",record.rank,record.player_id,record.handle_name,record.score)
	}

	return 0
}

func ReadData(entry_data string, play_data string) (*[]Entry_data, *[]Play_data, error) {
	f1, err := os.Open(entry_data)
	f2, err := os.Open(play_data)
	if err != nil {
		return nil,nil,err
	}
	defer f1.Close()
	defer f2.Close()

	entry_reader := csv.NewReader(f1)
	play_reader := csv.NewReader(f2)

	// ヘッダー行を飛ばす
	entry_reader.Read()
	play_reader.Read()
	
	entry_users := []Entry_data{}
	play_logs := []Play_data{}
	for {
		entry_user := Entry_data{}
		entry_record, err := entry_reader.Read()
		if err == io.EOF  {
			break
		}
		
		entry_user.player_id = entry_record[0]
		entry_user.handle_name = entry_record[1]

		entry_users = append(entry_users, entry_user)
	}

	for {
		play_log := Play_data{}
		play_record, err := play_reader.Read()
		if err == io.EOF  {
			break
		}
		s, _ := strconv.Atoi(play_record[2])
		t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", play_record[0])
		play_log.player_id = play_record[1]
		play_log.score = s
		play_log.playedAt = t

		play_logs = append(play_logs, play_log)
	}
	
	return &entry_users,&play_logs,nil
}

func SearchEntryUser(entryUsers *[]Entry_data, playLogs *[]Play_data) (*[]EntryUser_logs, error) {
	entryUserlogs := []EntryUser_logs{}
	for _, entryUser := range *entryUsers {
		maxScore := -1
		entried := false
		playedAt := time.Time{}
		for _, playLog := range *playLogs {
			if entryUser.player_id == playLog.player_id && playLog.score > maxScore {
				entried = true		
				maxScore = playLog.score
				playedAt = playLog.playedAt
			}
		}

		if entried && maxScore != -1 {
			entryUserLog := EntryUser_logs{
								rank: -1,
								player_id: entryUser.player_id,
								handle_name: entryUser.handle_name,
								score: maxScore,
								playedAt:playedAt,
							}
			
			entryUserlogs = append(entryUserlogs, entryUserLog)
		}
	}

	return &entryUserlogs,nil
}

func SortedByScorePlayer(entryPlayerLogs *[]EntryUser_logs) {
	sort.Slice(*entryPlayerLogs,func(i,j int) bool {
		if (*entryPlayerLogs)[i].score > (*entryPlayerLogs)[j].score {
			return true
		}
		return false
	})
}

func RankingPlayer(sortedPlayer *[]EntryUser_logs) error {
	rank := 1
	for i, player := range *sortedPlayer {
		if rank > 10 {
			break
		}

		if i == len(*sortedPlayer) - 1 {
			(*sortedPlayer)[i].rank = rank
			break
		}

		(*sortedPlayer)[i].rank = rank


		if player.score == (*sortedPlayer)[i+1].score {
			continue
		}
		rank++
	}

	return nil
}

func SortedByIdPlayer(rankedPlayer *[]EntryUser_logs) {
	sort.Slice(*rankedPlayer,func(i,j int) bool {
		if (*rankedPlayer)[i].rank == (*rankedPlayer)[j].rank {
			if (*rankedPlayer)[i].player_id < (*rankedPlayer)[j].player_id {
				return true
			}
		}
		return false
	})
}