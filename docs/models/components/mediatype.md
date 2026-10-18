# MediaType

The type of media to retrieve or filter by.

1 = movie
2 = show
3 = season
4 = episode
5 = trailer
7 = person
8 = artist
9 = album
10 = track
12 = clip
13 = photo
14 = photoalbum
15 = playlist
16 = playlistfolder
18 = collection

E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries

## Example Usage

```go
import (
	"github.com/LukeHagar/plexgo/models/components"
)

value := components.MediaTypeMovie
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `MediaTypeMovie`          | 1                         |
| `MediaTypeTvShow`         | 2                         |
| `MediaTypeSeason`         | 3                         |
| `MediaTypeEpisode`        | 4                         |
| `MediaTypeTrailer`        | 5                         |
| `MediaTypePerson`         | 7                         |
| `MediaTypeArtist`         | 8                         |
| `MediaTypeAlbum`          | 9                         |
| `MediaTypeTrack`          | 10                        |
| `MediaTypeClip`           | 12                        |
| `MediaTypePhoto`          | 13                        |
| `MediaTypePhotoAlbum`     | 14                        |
| `MediaTypePlaylist`       | 15                        |
| `MediaTypePlaylistFolder` | 16                        |
| `MediaTypeCollection`     | 18                        |