# imgdler
imgdler is a command for downloading tweet's images and read them.

## Install
```bash
go install github.com/kazdevl/imgdler/cmd/imgdler@latest
```

## How To Use
### Help
imgdler have three main commands.
- start: Periodically downloads images of tweets that match the specified criteria.
- list: Get a list of authors whose names are available for reading
- open: Open browser for reading the image list of a given author.
### Start
start is a command that periodically downloads images of tweets that match the specified criteria.

above command downloads specified tweet's images every night at 9:00 p.m.
- "author name" specifies author name of tweets that user of imgdler want.
    - this is twitter user name withour "@".
- "keyword" specifies keywords that the tweet should contain.
- "token" specifies twitter access token.
- "max" specifies the number of tweets to retrieve.

```bash
imgdler start -a [author name] -k [keyword] -t [token] -m [max]
```

### list
list is a command for getting a list of authors whose names are available for reading
```bash
$ imgdler list
# result if you have _kz_dev's images
The list of author names that you can read
[0]: _kz_dev
You cna read with `imgdler open [author name]`

$ imgdler list
# result if you don't have any auhtor's images
The list of author names that you can read
no author names
You cna read with `imgdler open [author name]`
```

### open
open is a command for opening browser for reading the image list of a given author.
```bash
$ imgdler open [author name]
# if you have _kz_dev's images
$ imgdler open _kz_dev
```
**result of open**
