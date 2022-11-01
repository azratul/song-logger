package config

var Conf Config

type Config struct {
	LastSong   string `json:"lastsong"`
	IcecastURL string `json:"icecast_url"`
	GDrive     struct {
		ClientEmail string   `json:"client_email"`
		PrivateKey  string   `json:"private_key"`
		Parents     []string `json:"parents"`
	}
	Logs struct {
		Filename string `json:"filename"`
		MimeType string `json:"mime_type"`
	}
}

type Radio struct {
	Icestats struct {
		Admin          string   `json:"admin"`
		Host           string   `json:"host"`
		Location       string   `json:"location"`
		ServerID       string   `json:"server_id"`
		ServerStart    string   `json:"server_start"`
		ServerStartISO string   `json:"server_start_iso8601"`
		Source         []Source `json:"source"`
	} `json:"icestats"`
}

type Source struct {
	Artist            string `json:"artist"`
	AudioBitrate      uint32 `json:"audio_bitrate"`
	AudioChannels     uint8  `json:"audio_channels"`
	AudioInfo         string `json:"audio_info"`
	AudioSamplerate   uint32 `json:"audio_samplerate"`
	Channels          uint8  `json:"channels"`
	Genre             string `json:"genre"`
	IceBitrate        uint8  `json:"ice-bitrate"`
	ListenerPeak      uint32 `json:"listener_peak"`
	Listeners         uint32 `json:"listeners"`
	ListenURL         string `json:"listenurl"`
	Quality           string `json:"quality"`
	Samplerate        uint32 `json:"samplerate"`
	ServerDescription string `json:"server_description"`
	ServerName        string `json:"server_name"`
	ServerType        string `json:"server_type"`
	StreamStart       string `json:"stream_start"`
	StreamStartISO    string `json:"stream_start_iso8601"`
	Subtype           string `json:"subtype"`
	Title             string `json:"title"`
	Dummy             string `json:"dummy"`
}
