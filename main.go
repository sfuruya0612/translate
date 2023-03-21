package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	domain     = flag.String("domain", os.Getenv("DEEPL_DOMAIN"), "(Required) DeepL API domain (default $DEEPL_DOMAIN)")
	authKey    = flag.String("authkey", os.Getenv("DEEPL_AUTH_KEY"), "(Required) DeepL API auth key (default $DEEPL_AUTH_KEY)")
	sourceLang = flag.String("source-lang", "JA", "Language of the text to be translated")
	targetLang = flag.String("target-lang", "EN", "The language into which the text should be translated")
	text       = flag.String("text", "", "(Required) Text to be translated")
	debug      = flag.Bool("debug", false, "Use debug mode")
)

type TranslationResult struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

func validate() error {
	if len(*domain) == 0 || len(*authKey) == 0 || len(*text) == 0 {
		return fmt.Errorf("Required parameter not specified. Please check help (-h option).")
	}
	return nil
}

func callApi(endpoint, authKey string) (string, error) {
	val := url.Values{}
	val.Add("source_lang", *sourceLang)
	val.Add("target_lang", *targetLang)
	val.Add("text", *text)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(val.Encode()))
	if err != nil {
		return "", fmt.Errorf("NewReq Error: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authKey)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("Do %v", err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	var result TranslationResult
	switch resp.StatusCode {
	case 200:
		if err := json.Unmarshal(b, &result); err != nil {
			return "", fmt.Errorf("Json Unmarshal %v", err)
		}
	default:
		fmt.Printf("Status: %v\n", resp.Status)
		fmt.Printf("Body: %v", string(b))
	}

	return result.Translations[0].Text, nil
}

func run() error {
	flag.Parse()

	if err := validate(); err != nil {
		return fmt.Errorf("Validate Error: %v", err)
	}

	endpoint := fmt.Sprintf("https://%s/v2/translate", *domain)
	authKey := fmt.Sprintf("DeepL-Auth-Key %s", *authKey)

	if *debug {
		fmt.Printf("Domain: %v\n", endpoint)
		fmt.Printf("Auth_Key: %v\n", authKey)
		fmt.Printf("SourceLang: %v\n", *sourceLang)
		fmt.Printf("TargetLang: %v\n", *targetLang)
		fmt.Printf("Text: %v\n", *text)
	}

	resp, err := callApi(endpoint, authKey)
	if err != nil {
		return fmt.Errorf("Call API Error: %v", err)
	}

	fmt.Println(resp)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
