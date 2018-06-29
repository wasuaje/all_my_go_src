
package main



import (
	"os"
	"thesaurus"
	"bufio"
	"log"
	"fmt"
)

func main() {
	//apiKey := os.Getenv("BHT_APIKEY")
	apiKey := "c3799ce3bb4364cd0322c38c1241fee4"
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for  "+word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for " + word )
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
