# Reddit Frontpage Analyzer (go)

This go application does the following:

- Fetches the top 25 posts from [/r/all][1]
- Stores all data in PostgreSQL
- Analyzes (tags) images with the [Microsoft Computer Vision API][3]

For the NodeJS serverless version see [here][2].

### Installation / Running

Clone the repository
```
λ git clone git@github.com:swordbeta/reddit-frontpage-analyzer-go.git && cd reddit-frontpage-analyzer-go
```

Copy the config file and edit
```
cp config.yaml.default config.yaml && vim config.yaml
```

Build docker image
```
λ docker build -t reddit-frontpage-analyzer-go .
```

Run docker container
```
λ docker run --rm --name reddit-frontpage-analyzer-go reddit-frontpage-analyzer-go
```

### Roadmap

- [X] Fetch reddit frontpage
- [X] Save unique posts to PostgreSQL
- [X] Tag images with [Microsoft Computer Vision API][3]
- [X] Add instructions for running in README
- [X] Gracefully exit current run when hitting rate limits
- [ ] Add tests
- [ ] Add Travis CI support
- [ ] Add code coverage and other badges

[1]: https://reddit.com/r/all
[2]: https://github.com/swordbeta/reddit-frontpage-analyzer-nodejs
[3]: https://www.microsoft.com/cognitive-services/en-us/computer-vision-api