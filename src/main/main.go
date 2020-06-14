package main

import (
	"fmt"
	"github.com/Jammizzle/watchlist-alert/src/data"
	"github.com/Jammizzle/watchlist-alert/src/logging"
	"github.com/Jammizzle/watchlist-alert/src/models"
	"github.com/Jammizzle/watchlist-alert/src/notification"

	//"github.com/Jammizzle/watchlist-alert/src/notification"
	"github.com/Jammizzle/watchlist-alert/src/remote"
	"io/ioutil"
	"sort"
	"strconv"
	"time"
)

func main() {

	db, err := data.CreateConnection()
	if err != nil {
		logging.Fatal(err)
	}

	for {
		shows, err := db.GetShows()
		if err != nil {
			logging.Fatal(err)
		}

		for _, v := range shows {
			v := v
			// If we can't even compile the show pattern then just skip over
			if err = v.CompileRegex(); err != nil {
				logging.Errorf("Unable to compile regex for show %s, err:", v.Name, err)
				continue
			}

			go func() {
				res, err := remote.Get(v.URL)
				if err != nil {
					logging.Errorf("Unable to check episodes for %s, err:", v.Name, err)
					return
				}

				dataInBytes, err := ioutil.ReadAll(res)
				if err != nil {
					logging.Error(err)
					return
				}

				matches := v.Regex.FindAllStringSubmatch(string(dataInBytes), -1)
				if len(matches) == 0 {
					logging.Error("matches returned an empty slice")
					return
				}

				epList := make([]int, len(matches))
				for _, k := range matches {
					ep, err := strconv.Atoi(k[1])
					if err != nil {
						logging.Errorf("Failed to parse int, err: %s", err)
						continue
					}
					epList = append(epList, ep)
				}

				sort.Slice(epList, func(i, j int) bool {
					return epList[i] > epList[j]
				})
				logging.Infof("Latest episode for %s is episode %d", v.Name, epList[0])

				subs, err := db.GetShowSubscribers(v.ID)
				if err != nil {
					logging.Errorf("Failed to get subs, err: %s", err)
					return
				}
				alertSubs(subs, v, epList[0])
			}()
		}
		time.Sleep(time.Hour) // check again in 6 hours
	}
}

func alertSubs(subs models.Subscribers, show models.Show, episodeNumber int) {
	for _, sub := range subs {
		if sub.EpisodeNumber < episodeNumber {
			if sub.ViewerPushoverID != "" {
				if err := show.SendPushoverNotification(sub.ViewerPushoverID, episodeNumber); err != nil {
					logging.Errorf("Failed to alert %s, err: %s", sub.ViewerName, err)
				}
			}
			if sub.ViewerEmail != "" {
				err := notification.Mail{
					Recipient:   sub.ViewerEmail,
					Subject:     fmt.Sprintf("New %s episode [%d]", show.Name, episodeNumber),
					ContentType: "text/html",
				}.RenderAndSend("one_piece.html", sub)
				if err != nil {
					logging.Error(err)
				}
			}
		}
	}
}
