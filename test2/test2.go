// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("test2.go")
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()

// 	// getting size of file
// 	stat, err := file.Stat()
// 	if err != nil {
// 		return
// 	}

// 	// reading file
// 	bs := make([]byte, stat.Size())
// 	_, err = file.Read(bs)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println(string(bs))
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	dir, err := os.Open(".")
// 	if err != nil {
// 		return
// 	}
// 	defer dir.Close()
// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		return
// 	}
// 	for _, fi := range fileInfos {
// 		fmt.Println(fi.Name())
// 	}

// }

// package main

// import (
// 	"os"
// )

// func main() {
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()
// 	file.WriteString("Привет мир!")

// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	filepath.Walk(".\\..", func(path string, info os.FileInfo, err error) error {
// 		fmt.Println(path)
// 		return nil
// 	})
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	argsWithProg := os.Args
// 	argsWithoutProg := os.Args[1:]

// 	arg := os.Args[0]

// 	fmt.Println(argsWithProg)
// 	fmt.Println(argsWithoutProg)
// 	fmt.Println(arg)
// }

// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	strPtr := flag.String("str", "hello", "a string")
// 	numPtr := flag.Int("num", 42, "an int")
// 	boolPtr := flag.Bool("fork", false, "a bool")

// 	var strVar string
// 	flag.StringVar(&strVar, "strVar", "world", "a string var")

// 	flag.Parse()

// 	fmt.Println("str:", *strPtr)
// 	fmt.Println("num:", *numPtr)
// 	fmt.Println("fork:", *boolPtr)
// 	fmt.Println("strVar:", strVar)
// 	fmt.Println("tail:", flag.Args())
// }

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.OpenFile("file.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	log.SetOutput(f)

// 	log.Print("save string to the log file")
// }

//

package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	note := &Notes{To: "Шеф",
		From:    "Меня",
		Heading: "Извинения",
		Body:    "Ну очень большие сиськи! Класс!",
	}

	file, err := xml.MarshalIndent(note, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("note1.xml", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
