package sensortag

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"module/httpcmd"

	"github.com/PuerkitoBio/goquery"
)

type Data struct {
	endpoint      string
	read_interval float64
	leftDown      chan bool
	rightDown     chan bool
}

func (sd *Data) LeftDownEvent() <-chan bool {
	return sd.leftDown
}

func (sd *Data) RightDownEvent() <-chan bool {
	return sd.rightDown
}

func (sd *Data) setLeftDown() {
	sd.leftDown <- true
}

func (sd *Data) setRightDown() {
	sd.rightDown <- true
}

func light(mode string) {
	cmd := "./cmd/remocon"
	arg := ""

	switch mode {
	case "on":
		arg = fmt.Sprintf("-d %s", httpcmd.ReadSeq("./cmd/light_all.txt"))
	case "off":
		arg = fmt.Sprintf("-d %s", httpcmd.ReadSeq("./cmd/light_off.txt"))
	case "fav":
		arg = fmt.Sprintf("-d %s", httpcmd.ReadSeq("./cmd/light_fav.txt"))
	}

	if arg != "" {
		httpcmd.RunCommand(cmd, arg)
	} else {
		log.Println("light: argument error")
	}
}

func (sd *Data) getPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("p").Each(func(_ int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		if id == "key" {
			text := s.Text()
			switch text {
			case "1":
				log.Println("left button pushed")
				light("fav")
			case "2":
				log.Println("right button pushed")
				light("off")
			case "3":
				log.Println("both button pushed")
				light("on")
			}
		}
	})
}

//func (sd *Data) getPage(url string) {
//	doc, _ := goquery.NewDocument(url)
//	doc.Find("p").Each(func(_ int, s *goquery.Selection) {
//		id, _ := s.Attr("id")
//		log.Println(id)
//		text := s.Text()
//		log.Println(text)
//	})
//	//doc.Find("a").Each(func(_ int, s *goquery.Selection) {
//	//	url, _ := s.Attr("href")
//	//	fmt.Println(url)
//	//})
//}

//type Result struct {
//	Url string
//}
//func (sd *Data) parseItem(r io.Reader) []Result {
//	results := []Result{}
//	doc, err := html.Parse(r)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println("a!")
//
//	//var result Result
//	var f func(*html.Node)
//	f = func(n *html.Node) {
//		log.Println("b! ", n)
//		// n.Typeでノードの型をチェックできる、ElementNodeでHTMLタグのNode。
//		// n.Dataでノートの値をチェックする、aタグをチェックしている
//		if n.Type == html.ElementNode && n.Data == "p" {
//			// n.Attrで属性を一覧する
//			// ここでもう少し頑張るとparseできる
//			for _, p := range n.Attr {
//				///if p.Key == "href" {
//				//if p.Id == "href" {
//				//	result.Url = a.Val
//				//	results = append(results, result)
//				//}
//				log.Println(p)
//			}
//			log.Println(n)
//		}
//		for c := n.FirstChild; c != nil; c = c.NextSibling {
//			f(c)
//		}
//	}
//	f(doc)
//	return results
//}

func (sd *Data) getSensorData() (string, error) {
	ret := ""
	resp, err := http.Get(sd.endpoint)
	defer resp.Body.Close()
	if err == nil {
		byteArray, read_err := ioutil.ReadAll(resp.Body)
		if read_err == nil {
			//sd.parseItem(resp.Body)
			sd.getPage(sd.endpoint)
			ret = string(byteArray)
			err = nil
		} else {
			log.Println("getSensorData: GET error ", read_err)
			err = read_err
		}
	} else {
		log.Println(err)
	}
	return ret, err
}

func (sd *Data) ReadButtonThread() {
	t := time.NewTicker(time.Duration(sd.read_interval) * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			s, err := sd.getSensorData()
			if err == nil {
				log.Println("ReadButtonThread: ", s)
			} else {
				log.Println("ReadButtonThread: read error occurred... perhaps.")
			}
		}
	}

}

func NewData(url string, interval float64) (*Data, error) {
	sd := &Data{
		endpoint:      url,
		read_interval: interval,
		leftDown:      make(chan bool),
		rightDown:     make(chan bool),
	}
	return sd, nil
}
