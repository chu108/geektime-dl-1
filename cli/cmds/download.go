package cmds

import (
	"encoding/json"
	"fmt"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/downloader"
	"github.com/mmzou/geektime-dl/service"
	"github.com/urfave/cli"
)

//NewDownloadCommand login command
func NewDownloadCommand() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:      "download",
			Usage:     "下载",
			UsageText: appName + " download",
			Action:    downloadAction,
			Before:    authorizationFunc,
		},
	}
}

func downloadAction(c *cli.Context) error {
	course, articles, err := application.CourseWithArticles(301)

	if err != nil {
		return err
	}

	downloadData := extractDownloadData(course, articles)
	// printExtractDownloadData(downloadData)

	downloader.Download(downloadData)

	return nil
}

func extractDownloadData(course *service.Course, articles []*service.Article) downloader.Data {
	downloadData := downloader.Data{
		Title: course.ColumnTitle,
	}
	data := downloader.EmptyData
	if course.IsColumn() {
		key := "default"
		for _, article := range articles {
			if !article.IncludeAudio {
				//	continue
			}
			urls := []downloader.URL{
				{
					URL:  article.AudioDownloadURL,
					Size: article.AudioSize,
					Ext:  "mp3",
				},
			}

			streams := map[string]downloader.Stream{
				key: downloader.Stream{
					URLs:    urls,
					Size:    article.AudioSize,
					Quality: key,
				},
			}

			data = append(data, downloader.Datum{
				ID:      article.ID,
				Title:   article.ArticleTitle,
				IsCanDL: article.IsCanPreview(),
				Streams: streams,
				Type:    "audio",
			})
		}
	}

	downloadData.Data = data

	return downloadData
}

func printExtractDownloadData(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", jsonData)
	}
}
