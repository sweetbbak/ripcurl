package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	urlx "net/url"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var usage = `Usage:
	-h --help
	-u --url
	-s --stdin
Operations:
	ripcurl --url <url>
Examples:
	ripcurl --url <url> > out.txt
	ripcurl --url <url> | bat
	curl -fsSl <url> | ripcurl | bat
	`

func print_help() {
	fmt.Println(usage)
}

func is_stdin_open() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// fmt.Println("data is being piped to stdin")
		return true
	} else {
		// fmt.Println("stdin is from a terminal")
		return false
	}
}

func get_printable(text string) string {
	text = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, text)
	return text
}

func rm_symbols(s string) string {
	// re := regexp.MustCompile("[[:^ascii:]]")
	// re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	re := regexp.MustCompile(`[^a-zA-Z0-9 \!\?\,\.\(\)]+`)
	t := re.ReplaceAllLiteralString(s, "")
	t = strings.Replace(t, "'", "", -1)
	t = strings.Replace(t, "<", "", -1)
	t = strings.Replace(t, ">", "", -1)
	t = strings.Replace(t, "  ", " ", -1)
	t = strings.Replace(t, "\n\n", "\n", -1)
	return t
}

func clean_text(text []string) []string {
	// process each line of text
	for x := range text {
		text[x] = get_printable(text[x])
		text[x] = rm_symbols(text[x])
	}
	return text
}

func create_doc(r io.Reader) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println(err)
	}
	return doc
}

func parse_doc_ptag(doc *goquery.Document) []string {
	// takes go doc and gets text in paragraph tags
	var text []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		text = append(text, t)
	})
	return text
}

// potentially allow for PUP like selector options
func parse_doc_ptag_custom(doc *goquery.Document, tag ...string) []string {
	// takes go doc and gets text in paragraph tags
	var text []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		text = append(text, t)
	})
	return text
}

func request(url string, ua ...string) (*http.Response, error) {
	// create cloudflare resistent client
	cl := tls_client()
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)

	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return resp, err
}

func parse_stdin() {
	a := bufio.NewReader(os.Stdin)
	doc := create_doc(a)
	lines := parse_doc_ptag(doc)
	lines = clean_text(lines)

	for x := range lines {
		fmt.Println(lines[x])
	}
}

func process_url(url string) {
	u, err := urlx.ParseRequestURI(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	resp, err := request(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	doc := create_doc(resp.Body)
	lines := parse_doc_ptag(doc)
	lines = clean_text(lines)

	for x := range lines {
		fmt.Println(lines[x])
	}
}

var user_agentx = `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`

var (
	url        string
	user_agent string
	stdin_bool bool
	pstdout    bool
	helpBool   bool
)

func init() {
	flag.StringVar(&url, "url", "", "Url to request")
	flag.StringVar(&url, "u", "", "Url to request")

	var user_agentx = `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`
	flag.StringVar(&user_agent, "user-agent", user_agentx, "user agent to use")
	flag.StringVar(&user_agent, "U", user_agentx, "user agent to use")

	flag.BoolVar(&stdin_bool, "stdin", false, "Read HTML from Stdin ie curl xyz.io | bin")
	flag.BoolVar(&stdin_bool, "s", false, "Read HTML from Stdin ie curl xyz.io | bin")

	flag.BoolVar(&pstdout, "out", false, "Print to stdout")
	flag.BoolVar(&pstdout, "o", false, "Print to stdout")

	flag.BoolVar(&helpBool, "help", false, "Print help")
	flag.BoolVar(&helpBool, "h", false, "Print help")
}

// TODO change url variable to something else to avoid net/url conflict
func main() {
	flag.Parse()

	if helpBool == true {
		print_help()
		os.Exit(0)
	}

	stdin_open := is_stdin_open()

	if url == "" && stdin_bool == false && stdin_open == false {
		print_help()
		os.Exit(0)
	}

	// explicitly check flag and stdin or
	// try to infer if the action was to pipe data in
	if stdin_bool == true && stdin_open == true {
		parse_stdin()
	}
	if stdin_open == true && url == "" {
		parse_stdin()
	}

	if url != "" {
		process_url(url)
		os.Exit(0)
	} else if url == "" && stdin_open == false && stdin_bool == false {
		fmt.Println("Input url or pipe in HTML")
		print_help()
		os.Exit(1)
	}

}
