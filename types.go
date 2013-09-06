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
