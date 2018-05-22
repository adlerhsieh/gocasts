### Screencast Assets

- thumbnail: 325x164
- video: tag class "video", width 1280, and height 720

### Remote Golang Installation & Update

https://github.com/golang/go/wiki/Ubuntu

### Deployment

https://github.com/pressly/sup

`Supfile`

```
sup production deploy
```

### Boot Reomte Server

The server is boot with `nohup` for now. Remotely run:

```
gocasts && server
```

The alias are written in `~/.bashrc`.

### Cleanup Log

Logs live on the same dir with the app. Remotely do:

```
gocasts && rm gocasts.log
```
