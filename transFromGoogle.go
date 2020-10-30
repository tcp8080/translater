package main

import (
	"fmt"
    "os"
    "bufio"
	"strings"
	"github.com/bregydoc/gtranslate"
)

func transSentence( sentence string) string{
    if len(sentence) == 0 {
      return "----------------"
    }

    translated, err := gtranslate.TranslateWithParams(
		sentence,
		gtranslate.TranslationParams{
			From: "en",
			To:   "zh",
		},
	)
	if err != nil {
		panic(err)
	}
	
	return translated
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func test() {

  //text := "The daughter of a man who died after an encounter with police in New York says she believes there's a cover-up in her father's death."
  text := "On behalf of the staff, this is Shane Bigham in the Chinese capital, hoping you'll join us for the next edition of the Beijing Hour, and open a window to the world together."
  translated := transSentence(text)
  fmt.Printf("en: %s | zh: %s \n", text, translated)
  // en: Hello World | ja: こんにちは世界

    f, err := os.Create("test.md")
    check(err)
    defer f.Close()

    _, err = f.WriteString(text + "\n" + translated)
    check(err) 

}

func translateArticle(filename string, outputfile string) {
/*
	text := "The daughter of a man who died after an encounter with police in New York says she believes there's a cover-up in her father's death."

    translated := transSentence(text)
	fmt.Printf("en: %s | zh: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
*/
	
    file, err := os.Open(filename)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // This is our buffer now
    var lines []string

    for scanner.Scan() {
      lines = append(lines, scanner.Text())
    }

    f, err := os.Create(outputfile)
    check(err)

    defer f.Close()

    readSum := 0
    wholelen := len(lines)
    
    //var buffer bytes.Buffer
    for _, line := range lines {
      fmt.Println("translating...", 100*readSum/wholelen, "%") 
      readSum = readSum +1

      linestr := strings.TrimSpace(line)
      chstr := transSentence(linestr)

      //buffer.WriteString(strings.Join([]string{linestr}, "\n\n"))
      //buffer.WriteString(strings.Join([]string{chstr}, "\n\n"))
	    _, err = f.WriteString(linestr + "\n" + chstr + "\n")
    }

    //_, err = f.WriteString(buffer.String())
    //check(err) 
}


func main(){
   translateArticle("test_en.txt", "test_en_ch_output.txt")
   //test()
}