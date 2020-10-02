package worker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var result = make(map[string]int)
//Worker make a list of Urls and  send it to the 
func Worker (phrase string,urls string, file string) {
	guard := make(chan struct{}, 1)
	var wg sync.WaitGroup
	if file != "" {
		dat, err := ioutil.ReadFile(file)
		if err != nil{
			log.Fatal("Mistake with read file")
			return
		}
		for _, url := range strings.Split(string(dat),","){
			if !strings.Contains(url, "http"){
			url = strings.Join([]string{"https://", url},"")
			result[url] = 0;
			}
		}
	}

	if urls !="" {
		for _, url := range strings.Split(urls,","){
			if !strings.Contains(url, "http"){
			url = strings.Join([]string{"https://", url},"")
			result[url]=0;
			}
		}
	}

	fmt.Println(result)

	for url := range result {
		wg.Add(1)
		guard<-struct{}{}
		go func(wg *sync.WaitGroup){
			defer wg.Done()
			res, err := http.Get(url)
				if err != nil {
					result[url] = -1
				}
			robots, err := ioutil.ReadAll(res.Body)
			if err!=nil{
				log.Fatal("Не могу прочитать страничку")
				return
			}
			result[url]= strings.Count(strings.ToLower(string(robots)), phrase)
			<- guard
		}(&wg)
	}
	wg.Wait()
	fmt.Println(result)
}
