package scrapper

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	BaseURL = "https://comic.naver.com/genre/"
)

var (
	errNotvalidURL = errors.New("This is not allowed URL")
)

type comic struct {
	title string
	writer string
	description string
	rating string
	imageURL string
}

func Scrapper(comic_type string) {
	var comics []comic
	var c = make(chan []comic)
	var url string = BaseURL;
	switch comic_type {
		case "challenge":
			url += "challenge"
		case "bestchallenge":
			url += "bestChallenge"
		default: 
			url += "bestChallenge"
	}
	endPageURL := url + "?&page=" + strconv.Itoa(200)
	totalPage := GetTotalPagesCount(endPageURL);
	for i:=0;i<totalPage;i++ {
		go GetPage(endPageURL, i+1, c)
	}
	for i:=0;i<totalPage;i++ {
		extractedComics := <- c
		comics = append(comics, extractedComics...)
	}
	writeComicList(comics)
	fmt.Println("Webtoon list making done. Total count: ", len(comics))
}

func GetPage(url string, i int, c chan<- []comic) {
	var comics []comic
	c_1 := make(chan comic)
	pageURL := url + "?&page=" + strconv.Itoa(i)
	
	fmt.Println("Get " + pageURL + "...")

	res, err := http.Get(pageURL)
	CheckErr(err)
	CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckErr(err)

	card := doc.Find(".weekchallengeBox td").Not(".challengeListDot")
	card.Each(func(_ int, card *goquery.Selection){
		go extractComic(card, c_1)
	})
	for i:=0;i<card.Length();i++ {
		comic := <- c_1
		comics = append(comics, comic)
	}
	c <- comics
}

func writeComicList(comics []comic) {
	file, mErr := os.Create("comic_list.csv")
	CheckErr(mErr)

	w := csv.NewWriter(file)
	defer w.Flush()
	header := []string{"Title", "Writer", "Description", "rating", "Thumbnail"}
	
	wErr := w.Write(header)
	CheckErr(wErr)

	for _, comic := range comics {
		wErr := w.Write([]string{comic.title, comic.writer, comic.description, comic.rating, comic.imageURL})
		CheckErr(wErr)
	} 
}

func extractComic(card *goquery.Selection, c chan<- comic)  {
	title := CleanString(card.Find(".challengeTitle > a").Text())
	writer := CleanString(card.Find(".user").Text())
	description := CleanString(card.Find(".summary").Text())
	rating := CleanString(card.Find(".rating_type > strong").Text())
	imageURL := CleanString(card.Find(".fl img").AttrOr("src", ""))

	c <- comic{
		title: title,
		writer: writer,
		description: description,
		rating: rating,
		imageURL: imageURL,
	}
}

func CleanString(str string) string {
	return strings.TrimSpace(str)
}

func GetTotalPagesCount(url string) int {
	pages := 0
	res, err := http.Get(url)
	CheckErr(err)
	CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckErr(err)

	endPage, err := strconv.Atoi(doc.Find(".page_wrap .num_page").Last().Text());
	CheckErr(err);

	pages = endPage;

	return pages
}

func CheckErr(err error){
	if(err != nil){
		log.Fatal(err)
	}
}

func CheckCode(res *http.Response){
	if(res.StatusCode != 200){
		log.Fatal("The url is return failed.")
	}
}