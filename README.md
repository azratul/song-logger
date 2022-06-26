# Song Logger
Small app to record songs from Icecast stream

In my use case, as a systemd timer with env vars to set the Icecast URL and Google Drive credentials.

In the [service](https://github.com/azratul/song-logger/tree/main/service) folder you can find my systemd configuration files as an example. One is the track logger and the other is for uploading the file to Google Drive.

Recently I also added a cron just to send some Icecast's stats to Google Analytics
