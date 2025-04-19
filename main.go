package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// <<<<<<<<<<<<<<<<<<<<<<< Utilizando Algoritmo TRIE
type Node struct {
	Children map[string]*Node
}

func NewNode() *Node {
	node := &Node{
		Children: make(map[string]*Node),
	}
	return node
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	root := NewNode()
	return &Trie{RootNode: root}
}

func strippedWord(word string) string {
	return strings.ToLower(strings.ReplaceAll(word, " ", ""))
}

func (t *Trie) Insert(word string) error {
	current := t.RootNode
	for _, v := range strippedWord(word) {
		str := string(v)
		node, ok := current.Children[str]
		if !ok {
			current.Children[str] = NewNode()
			current = current.Children[str]
			continue
		}

		current = node
	}
	return nil
}

func (t *Trie) SearchWord(word string) bool {
	current := t.RootNode
	for _, v := range strippedWord(word) {
		str := string(v)
		if node, ok := current.Children[str]; ok {
			current = node
			continue
		}
		return false
	}
	return true
}

// >>>>>>>>>>>>>>>>>>>>>>>> Fin de algoritmo TRIE

func insertarConSlice(palabra string, palabras *[]string) {
	*palabras = append(*palabras, palabra)
}

func buscarConSlice(palabra string, palabras []string) bool {
	for _, p := range palabras {
		if p == palabra {
			return true
		}
	}
	return false
}

// ////////////////////////////////////////////////
func loadWords(filename string) ([]string, error) {
	var words []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func main() {
	// Crear estructuras
	trie := NewTrie()
	var palabras []string

	// Cargar el diccionario
	words, err := loadWords("words.txt") // Asegúrate de tener el archivo words.txt
	if err != nil {
		fmt.Println("Error al cargar las palabras:", err)
		return
	}

	// Prueba de inserción y búsqueda con Trie
	start := time.Now()
	for _, word := range words {
		trie.Insert(word)
	}
	durationTrieInsert := time.Since(start)

	start = time.Now()
	for _, word := range words {
		trie.SearchWord(word)
	}
	durationTrieSearch := time.Since(start)

	// Prueba de inserción y búsqueda con Slice ([]string)
	start = time.Now()
	for _, word := range words {
		insertarConSlice(word, &palabras)
	}
	durationSliceInsert := time.Since(start)

	start = time.Now()
	for _, word := range words {
		buscarConSlice(word, palabras)
	}
	durationSliceSearch := time.Since(start)

	// Imprimir resultados
	fmt.Printf("Tiempo de inserción con Trie: %v\n", durationTrieInsert)
	fmt.Printf("Tiempo de búsqueda con Trie: %v\n", durationTrieSearch)
	fmt.Printf("Tiempo de inserción con []string: %v\n", durationSliceInsert)
	fmt.Printf("Tiempo de búsqueda con []string: %v\n", durationSliceSearch)

}
