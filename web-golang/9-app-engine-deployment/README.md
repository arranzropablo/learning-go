# deploying to [Google Cloud](https://cloud.google.com/)
- [install google appengine](https://cloud.google.com/appengine/docs/go/download)
- [make sure python is installed VERSION 2.7.x](https://www.python.org/downloads/release/python-2712/)
  - python -V
  - configure environment PATH variables
- google cloud developer console
  - create a project
- gcloud init selecting the app we want

```
runtime: go
api_version: go1

handlers:
  - url: /.*
    script: _go_app
```
- test the deployment with
```
 dev_appserver.py app.yaml
```
- deploy to that project
```
gcloud app deploy
```
- view your project
  - http://YOUR_PROJECT_ID.appspot.com/


example
http://temp-145415.appspot.com/


change DNS info at google domains
point your domain to your appengine project
(need to research if I can do this w strato sub-URL of arranzropablo.com, either way I can buy arranzropablo.pw, its 9€ year)
