package concurrency

//mport "time"

type WebsiteChecker func(string ) bool


func CheckWebsites( wc WebsiteChecker , urls [] string) map[string]bool {


	results := make(map[string]bool)

	channel := make(chan result)

	for _ , url := range urls {
		 go func( url_ string) {
		 	
			channel <- result{url:url_ , flag:wc(url_)}
		 }(url)

		//results[url] = wc(url)
	}

	for index := 0 ; index < len(urls); index++ {
		r := <- channel
		results[r.url] = r.flag
	}
	return results
}

type result struct {
	url string
	flag bool
}

