//
// A collection of Rdio API object types:
// http://www.rdio.com/developers/docs/web-service/types/
//

package rdio

type Album struct {
	Name        string   // the name of the album
	Type        string   // the type of the object, always "a"
	Icon        string   // the URL to the cover art for the album
	BaseIcon    string   // the URL to the cover art for the album
	Url         string   // the URL of the album on the Rdio site
	Artist      string   // the name of the artist that released the album
	ArtistUrl   string   // the URL of the artist that released the album on the Rdio site
	IsExplicit  bool     // is the album explicit?
	IsClean     bool     // is the album clean?
	Length      int      // number of tracks on the album
	ArtistKey   string   // the key of the artist that released the album
	TrackKeys   []string // the keys of the tracks on the album
	Price       string   // the price of the album in the requesting user's currency, if available for download
	CanStream   bool     // the album can be streamed
	CanSample   bool     // the album can be previewed
	CanTether   bool     // the album can be sync to mobile devices
	ShortUrl    string   // a short URL for the album
	EmbedUrl    string   // the URL of a SWF to embed the album
	DisplayDate string   // the release date of the album, human readable
	Key         string   // the key of the album
	ReleaseDate string   // the release date of the album
	Duration    int      // the duration of the album in seconds
}

type Artist struct {
	Name        string // the name of the artist
	Key         string // the artist's key
	Type        string // the object type, always "r"
	Url         string // the URL of the artist on the Rdio web site
	Length      int    // the number of tracks that the artist has on Rdio
	Icon        string // an image for the artist
	BaseIcon    string // an image for the artist, partial URL
	HasRadio    bool   // is a station available for the artist?
	ShortUrl    string // a short URL for the artist page
	RadioKey    string // the key of the station for artist recommendations
	TopSongsKey string // the key of the station for the artist's top songs
}

type Label struct {
	Name     string // the name of the label
	Key      string // the key of the label
	Type     string // the object type, always "l"
	Url      string // the url of the label on teh Rdio web site
	ShortUrl string // a short URL for the label page
	HasRadio bool   // is a station available for the label
	RadioKey string // the key of the station for label recommendations
}

type Track struct {
	Name                 string // the name of the track
	Artist               string // the name of the artist who performed the track
	Album                string // the name of the album that the track appears on
	AlbumKey             string // the key of the album that the track appears on
	AlbumUrl             string // the URL of the album that the track appears on, on the Rdio web site
	ArtistKey            string // the key of the track's artist
	ArtistUrl            string // the URL of the track's artist on the Rdio web site
	Type                 string // the object type, always "t"
	Length               int    // the number of tracks in the track, ie: 1
	Duration             int    // the duration of the track in seconds
	IsExplicit           bool   // is the track explicit?
	IsClean              bool   // is the track clean?
	Url                  string // the URL of the track on the Rdio web site
	BaseIcon             string // the partial URL of the album-art for the track
	AlbumArtist          string // the name of the artist whose album the track appears on
	AlbumArtistKey       string // the key of the artist whose album the track appears on
	CanDownload          bool   // the track can be downloaded
	CanDownloadAlbumOnly bool   // the track can only be downloaded as part of an album download
	CanStream            bool   // the track can be streamed
	CanTether            bool   // the track can be synced to mobile devices
	CanSample            bool   // the track can be previewed
	Price                string // the price of the album in the requesting user's currency, if available for download
	ShortUrl             string // a short URL for the track
	EmbedUrl             string // the URL of a SWF to embed the track
	Key                  string // the object key of the track
	Icon                 string // the URL of the album-art for the track
	RrackNum             string // the order within its album that this track appears
}

type Playlist struct {
	Name        string  // the name of the playlist
	Length      int     // the number of tracks in the playlist
	Type        string  // the object type, always "p"
	Url         string  // the URL of the playlist on the Rdio site
	Icon        string  // the URL of an icon for the playlist
	BaseIcon    string  // the URL of an icon for the playlist
	Owner       string  // the name of the user who created the playlist
	OwnerUrl    string  // the URL on the Rdio site of the user who created the playlist
	OwnerKey    string  // the key of the user who created the playlist
	OwnerIcon   string  // the icon of the user who created the playlist
	LastUpdated float32 // when the playlist was last modified
	ShortUrl    string  // a short URL for the playlist
	EmbedUrl    string  // the URL of a SWF to embed the playlist
	Key         string  // the key of the playlist
}

type User struct {
	Key            string // the object key of the user
	FirstName      string // the first name of the user
	LastName       string // the last name of the user
	Icon           string // the URL of an image of the user
	BaseIcon       string // the URL of an image of the user
	LibraryVersion int    // the library version of the user, used to determine if a user's collection has changed
	Url            string // the URL of the user on the Rdio site
	Gender         string // "m" or "f"
	Type           string // the object type, always "s"
}

type CollectionAlbum struct {
	Name          string   // the name of the album
	Type          string   // the object type of this object, always "al"
	Icon          string   // the URL to the cover art for the album
	BaseIcon      string   // the URL to the cover art for the album
	Url           string   // the URL of the album on the Rdio site
	Artist        string   // the name of the artist that released the album
	ArtistUrl     string   // the URL of the artist that released the album on the Rdio site
	IsExplicit    bool     // is the album explicit?
	IsClean       bool     // is the album clean?
	Length        int      // number of tracks on the album
	ArtistKey     string   // the key of the artist that released the album
	TrackKeys     []string // the keys of the tracks on the album
	Price         string   // the price of the album in the requesting user's currency, if available for download
	CanStream     bool     // the album can be streamed
	CanSample     bool     // the album can be previewed
	CanTether     bool     // the album can be sync to mobile devices
	ShortUrl      string   // a short URL for the album
	EmbedUrl      string   // the URL of a SWF to embed the album
	DisplayDate   string   // the release date of the album, human readable
	Key           string   // the key of the album
	ReleaseDate   string   // the release date of the album
	Duration      int      // the duration of the album in seconds
	UserKey       string   // the key of the user whose collection this album is in
	UserName      string   // the username of the user whose collection this album is in
	AlbumKey      string   // the key of the album
	AlbumUrl      string   // the url of the album
	CollectionUrl string   // the url to the collection
	ItemTrackKeys []string // track keys for all tracks on the album
}

type CollectionArtist struct {
	Name          string // the name of the artist
	Key           string // the artist's key
	Type          string // the object type of this object, always "rl"
	Url           string // the URL of the artist on the Rdio web site
	Length        int    // the number of tracks that the artist has on Rdio
	Icon          string // an image for the artist
	BaseIcon      string // an image for the artist, partial URL
	HasRadio      bool   // is a station available for the artist?
	ShortUrl      string // a short URL for the artist page
	RadioKey      string // the key of the station for artist recommendations
	TopSongsKey   string // the key of the station for the artist's top songs
	UserKey       string // the key of the user whose collection this artist is in
	UserName      string // the username of the user whose collection this artist is in
	ArtistKey     string // the key of the artist
	ArtistUrl     string // the url for the artist
	CollectionUrl string // the url to the collection
}

type LabelStation struct {
	Count          string   // the number of tracks in the station
	LabelName      string   // the name of the label
	Name           string   // the name of the label station
	HasRadio       bool     // is a station available for the label
	Tracks         []string // the tracks for the station
	LabelUrl       string   // the URL of the label on the Rdio site
	ShortUrl       string   // a short URL for the label page
	Length         int      // the number of tracks in the station
	Url            string   // the url of the label on teh Rdio web site
	Key            string   // the key of the label
	RadioKey       string   // the key of the station for label recommendations
	ReloadOnRepeat bool     // the station should be reloaded when it completes playing and repeat is enabled
	Type           string   // the object type, always "lr"
}

type ArtistStation struct {
	RadioKey       string   // the key of the station for artist recommendations
	TopSongsKey    string   // the key of the station for the artist's top songs
	BaseIcon       string   // an image for the artist, partial URL
	Tracks         []string // the tracks for the station
	ArtistUrl      string   // the URL of the artist on the Rdio site
	Key            string   // the key of the station
	ReloadOnRepeat bool     // the station should be reloaded when it completes playing and repeat is enabled
	Icon           string   // an image for the artist
	Count          int      // the number of tracks in the station
	Name           string   // the name of the station
	HasRadio       bool     // is a station available for the artist?
	Url            string   // the URL of the artist on the Rdio web site
	ArtistName     string   // the name of the artist
	ShortUrl       string   // a short URL for the artist page
	Length         int      // the number of tracks in the station
	Type           string   // the object type, always "rr"
}

type HeavyRotationStation struct {
	Type           string   // the object type, always "h"
	Key            string   // the key of the station
	Length         int      // the number of tracks in the station
	Tracks         []string // the tracks for the station
	ReloadOnRepeat bool     // the station should be reloaded when it completes playing and repeat is enabled
	Count          int      // the number of tracks in the station
	User           string   // the user
	BaseIcon       string   // the icon of the user
	Icon           string   // the icon of the user
	Name           string   //the name of the station
}

type HeavyRotationUserStation struct {
	Type           string   // the object type, always "e"
	Key            string   // the key of the station
	Length         int      // the number of tracks in the station
	Tracks         []string // the tracks for the station
	ReloadOnRepeat bool     // the station should be reloaded when it completes playing and repeat is enabled
	Count          string   // the number of tracks in the station
	User           string   // the user
	BaseIcon       string   // the icon of the user
	Icon           string   // the icon of the user
	Name           string   //the name of the station
}

type ArtistTopSongsStation struct {
	RadioKey       string   // the key of the station for artist recommendations
	TopSongsKey    string   // the key of the station for the artist's top songs
	BaseIcon       string   // an image for the artist, partial URL
	Tracks         []string // the tracks for the station
	ArtistUrl      string   // the URL of the artist on the Rdio site
	Key            string   // the key of the station
	ReloadOnRepeat bool     // the station should be reloaded when it completes playing and repeat is enabled
	Icon           string   // an image for the artist
	Count          int      // the number of tracks in the station
	Name           string   // the name of the station
	HasRadio       bool     // is a station available for the artist?
	Url            string   // the URL of the artist on the Rdio web site
	ArtistName     string   // the name of the artist
	ShortUrl       string   // a short URL for the artist page
	Length         int      // the number of tracks in the station
	Type           string   // the object type, always "tr"
}

type UserCollectionStation struct {
	Type           string   // the object type, always "c"
	Key            string   // the key of the station
	Length         int      // the number of tracks in the station
	Tracks         []string // the tracks for the station
	ReloadOnRepeat bool     // the station should be reloaded when it completes playing and repeat is enabled
	Count          int      // the number of tracks in the station
	User           string   // the user
	BaseIcon       string   // the icon of the user
	Icon           string   // the icon of the user
	Name           string   // the name of the station
	Url            string   // the URL of the collection
}

type UserPlaylists struct {
	Owned      []Playlist
	Collab     []Playlist
	Subscribed []Playlist
}
