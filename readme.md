# Vikunja Emoji
This is simple script to update task names with emojis.

It is being done with ChatGPT, which helps to provide proper emoji for given text. This app works best with english names, but should make do with other languages.

## Usage
Simplest option is to use it with docker

```bash
    docker run \
      -e OPENAI_API_KEY <key> \
      -e VIKUNJA_SERVICE_URL <vikunja url> \
      -e VIKUNJA_SERVICE_TOKEN <api-token> \
      image
```