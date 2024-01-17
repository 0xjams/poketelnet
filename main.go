package main

import (
	"log"
	"math/rand"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type Collection []Emoji
type PokemonCollection []Pokemon
type Letter struct {
	emoji  Emoji
	letter byte
}

type Emoji struct {
	Codes    string `json:"codes"`
	Char     string `json:"char"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Group    string `json:"group"`
	Subgroup string `json:"subgroup"`
}

type PokemonName struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}

type Pokemon struct {
	Names PokemonName `json:"name"`
}

var StartTime = time.Now()

func sayHi() string {
	return `
  
	██╗███╗░░██╗░██████╗███████╗░█████╗░██╗░░░██╗██████╗░██╗████████╗██╗░░░██╗
	██║████╗░██║██╔════╝██╔════╝██╔══██╗██║░░░██║██╔══██╗██║╚══██╔══╝╚██╗░██╔╝
	██║██╔██╗██║╚█████╗░█████╗░░██║░░╚═╝██║░░░██║██████╔╝██║░░░██║░░░░╚████╔╝░
	██║██║╚████║░╚═══██╗██╔══╝░░██║░░██╗██║░░░██║██╔══██╗██║░░░██║░░░░░╚██╔╝░░
	██║██║░╚███║██████╔╝███████╗╚█████╔╝╚██████╔╝██║░░██║██║░░░██║░░░░░░██║░░░
  
	I'm sorry, my responses are limited. You must ask the right question.
	Prove you can speak my language fluently and quickly. 
	Do it and you shall receive a flag.
  ` + "\r\n"
}

func generateRandomNumber(start int, end int, count int) []int {
	// scope check
	if end < start || (end-start) < count {
		return nil
	}
	// slice for storing results
	nums := make([]int, 0)
	// random number generator, add time stamp to ensure that the random number generated each time is different
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		// generate random number

		num := r.Intn(end-start) + start
		// duplicate check
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

var collection Collection
var pokemonCollection PokemonCollection

func main() {

	jsonFile, err := os.Open("emoji.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Emoji file not available")
	}

	defer jsonFile.Close()

	pokemonFile, err := os.Open("pokedex.json")
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Pokemon file not available")
	}
	defer pokemonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &collection)
	byteValue, _ = ioutil.ReadAll(pokemonFile)
	json.Unmarshal(byteValue, &pokemonCollection)
	var myHandler telnet.Handler = internalEchoHandler{}
	err = telnet.ListenAndServe(":5555", myHandler)

	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}

type internalEchoHandler struct{}

func buildAlphabet(alphabet *[]Letter) {
	var randomNumbers []int
	randomNumbers = generateRandomNumber(0, len(collection)-1, 26)
	var lowercaseStart = 97
	for i := 0; i < len(randomNumbers); i++ {
		
		var l Letter
		l.emoji = collection[randomNumbers[i]]
		l.letter = (byte)(lowercaseStart + i)
		*alphabet = append(*alphabet, l)
	}
}

func getRandomPokemon() string {
	var pokemon string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(pokemonCollection) - 1)
	pokemon = pokemonCollection[num].Names.English
	re, _ := regexp.Compile(`[^\w]`)
	pokemon = re.ReplaceAllString(pokemon, "")
	if strings.Contains(pokemon, "nido") {
		return getRandomPokemon()
	}
	return strings.ToLower(pokemon)
}

func printEmojiAlphabet(w telnet.Writer, alphabet *[]Letter) {
	for i := 0; i < len(*alphabet); i++ {
		li := fmt.Sprintf("%c = %s\n", (*alphabet)[i].letter, (*alphabet)[i].emoji.Char)
		oi.LongWriteString(w, li)
	}
}

func printWordAsEmoji(word string, alphabet *[]Letter, w telnet.Writer) {
	var abc = *alphabet
	var lowercaseStart = 97
	var fullWord = ""
	for _, chr := range word {
		var letterPosition = int(chr) - lowercaseStart
		fmt.Println(letterPosition)
		fullWord = fullWord + abc[letterPosition].emoji.Char
	}
	fmt.Println(fullWord)
	oi.LongWriteString(w, fullWord+" => ")
}

func (handler internalEchoHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	var alphabet []Letter
	buildAlphabet(&alphabet)
	var buffer [1]byte 
	p := buffer[:]
	s := ""
	var j int
	j = 0
	var pokemon string
	oi.LongWriteString(w, sayHi())
	printEmojiAlphabet(w, &alphabet)
	pokemon = getRandomPokemon()
	fmt.Println(pokemon)
	now := time.Now().UnixNano() / 1000000
	start := time.Now().UnixNano() / 1000000
	printWordAsEmoji(pokemon, &alphabet, w)
	for {

		n, err := r.Read(p)

		if n > 0 {
			
			sTmp := string(p[:n])
			
			s = s + sTmp
			fmt.Println(s)
			if s == "Hi\r\n" {
				s = ""
				
				oi.LongWriteString(w, "Hello to you too, bye\r\n")
				return
			}
			if s == "pokemon\r\n" {
				s = ""
				oi.LongWriteString(w, "Ah, I see you're a person of culture as well. That is not the answer, though. Bye\r\n")
				return
			}
			if strings.Contains(s, "\r\n") {

				if strings.Compare(s, pokemon+"\r\n") == 0 {
					j = j + 1
					now = time.Now().UnixNano() / 1000000
					fmt.Printf("Difference in time %d", (now - start))
					if now-start > 3000 {
						oi.LongWriteString(w, "You're too slow\r\n")
						return
					}
					oi.LongWriteString(w, "Correct.\r\n")
					if j < 150 {
						oi.LongWriteString(w, "Next word\r\n")
						start = time.Now().UnixNano() / 1000000
						pokemon = getRandomPokemon()
						fmt.Println(pokemon)
						printWordAsEmoji(pokemon, &alphabet, w)
					} else {
						var flag = "FLAG{U_SP34K_3M0JI_F4S7}"
						fmt.Println(flag)
						oi.LongWriteString(w, flag+"\r\n")
						return
					}

				} else {
					oi.LongWriteString(w, "Wrong answer, bye\r\n")
					return
				}
			}
			if sTmp == "\n" {
				
				s = ""
			}
		}

		if nil != err {
			print(err)
			break
		}
	}
}
