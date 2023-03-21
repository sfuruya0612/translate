# translate

Simple translation CLI tool using [DeepL API](https://www.deepl.com/ja/docs-api/translate-text/).

## Requirements

You will need an API authorization key and a domain(api-free.deepl.com or api.deepl.com) that matches your pricing plan.  
Please create an account at DeepL API.

## Installation

```bash
go install github.com/sfuruya0612/translate@latest
```

## Usage

```text
❯ translate -h
Usage of translate:
  -authkey string
        (Required) DeepL API auth key (default $DEEPL_AUTH_KEY)
  -debug
        Use debug mode
  -domain string
        (Required) DeepL API domain (default $DEEPL_DOMAIN)
  -source-lang string
        Language of the text to be translated (default "JA")
  -target-lang string
        The language into which the text should be translated (default "EN")
  -text string
        (Required) Text to be translated
```

## License

[MIT License](./LICENSE)
