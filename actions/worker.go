package actions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/edudorus/dead-man-switch/models"
)

func init() {
	go StartPingChecker()
}

func StartPingChecker() {
	nextTime := time.Now().Truncate(10 * time.Second)
	nextTime = nextTime.Add(10 * time.Second)
	time.Sleep(time.Until(nextTime))
	CheckPing()
	go StartPingChecker()
}

func CheckPing() {
	files := ReadFilesFromBlockchain()

	for index, file := range files {
		if file.Ping != "" {
			i, err := strconv.ParseInt(file.Ping, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			tm := time.Unix(i, 0)
			duration := time.Since(tm)
			fmt.Println(duration)
			if duration.Seconds() > float64(30) {
				entry := &models.Upload{}
				models.DB.Where("ipfs_hash = ?", file.IpfsHash).First(entry)
				writeKeyToBlockchain(index, entry.Key)
			}
		}
	}
}
