package viewmodel

import (
	"context"
	"strings"

	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleMediaInfo(id int) error {
	var mediaInfo *responses.MediaInfo
	var err error
	
	err = ui.ActionSpinner("Fetching media info...", func(ctx context.Context) error {
		mediaInfo, err = api.GetMediaInfo(id)
		return err
	})
	if err != nil {
		return err
	}
	
	isAnime := mediaInfo.Data.Media.Type == "ANIME"
	chapEp := func() int {
		if isAnime { return mediaInfo.Data.Media.Episodes } else { return mediaInfo.Data.Media.Chapters }
	}()
	durVol := func() int {
		if isAnime{ return mediaInfo.Data.Media.Duration } else { return mediaInfo.Data.Media.Volumes }
	}()
	tags := func() string {
		tagsList := make([]string, len(mediaInfo.Data.Media.Tags))
		for i, tag := range mediaInfo.Data.Media.Tags {
			tagsList[i] = tag.Name
		}
		return strings.Join(tagsList, ", ")
	}()
	studios := func() string {
		studioList := make([]string, len(mediaInfo.Data.Media.Studios.Nodes))
		for i, studio := range mediaInfo.Data.Media.Studios.Nodes {
			studioList[i] = studio.Name
		}
		return strings.Join(studioList, ", ")
	}()
	
	mediaInfoUI := &ui.MediaInfoUI{
		Id: mediaInfo.Data.Media.ID,
		MalId: mediaInfo.Data.Media.IDMal,
		Score: mediaInfo.Data.Media.MeanScore,
		EnglishTitle: mediaInfo.Data.Media.Title.English,
		RomajiTitle: mediaInfo.Data.Media.Title.Romaji,
		CoverImage: mediaInfo.Data.Media.CoverImage.ExtraLarge,
		NativeTitle: mediaInfo.Data.Media.Title.Native,
		IsAnime: isAnime,
		ChapterEpisode: chapEp,
		VolumeDuration: durVol,
		Genres: strings.Join(mediaInfo.Data.Media.Genres, ", "),
		Tags: tags,
		Studios: studios,
		Format: mediaInfo.Data.Media.Format,
		Description: mediaInfo.Data.Media.Description,
	}
	
	return mediaInfoUI.Render()
}