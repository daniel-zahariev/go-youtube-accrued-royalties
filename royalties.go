package go_youtube_accrued_royalties

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const AccruedRoyaltiesRecordFieldsCount = 15

//type AccruedRoyaltiesRecord struct {
//	CustomId                 string
//	AssetId                  string
//	Title                    string
//	Writers                  string
//	Iswc                     string
//	HfaSongCode              string
//	MatchPolicy              string
//	PublisherName            string
//	SyncOwnershipShare       string
//	SyncOwnershipTerritory   string
//	SyncOwnershipRestriction string
//	RelatedAssetId           string
//	SrIsrc                   string
//	SrArtist                 string
//	RelatedTitle             string
//}

func ReadFile(filePath string) (int, error) {
	linesCount := 0

	royaltiesFile, err := os.Open(filePath)
	if err != nil {
		return linesCount, err
	}

	csvReader := csv.NewReader(royaltiesFile)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return linesCount, err
		} else if length := len(line); length < AccruedRoyaltiesRecordFieldsCount {
			err = fmt.Errorf("line %v is incomplete. It only has %v columns", linesCount, length)
			return linesCount, err
		}
		if linesCount == 0 {
			linesCount += 1
			continue
		}

		//record := &AccruedRoyaltiesRecord{
		//	CustomId:                 line[0],
		//	AssetId:                  line[1],
		//	Title:                    line[2],
		//	Writers:                  line[3],
		//	Iswc:                     line[4],
		//	HfaSongCode:              line[5],
		//	MatchPolicy:              line[6],
		//	PublisherName:            line[7],
		//	SyncOwnershipShare:       line[8],
		//	SyncOwnershipTerritory:   line[9],
		//	SyncOwnershipRestriction: line[10],
		//	RelatedAssetId:           line[11],
		//	SrIsrc:                   line[12],
		//	SrArtist:                 line[13],
		//	RelatedTitle:             line[14],
		//}

		linesCount += 1
	}

	return linesCount, nil
}
