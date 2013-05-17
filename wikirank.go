package main

import (
  "log"
  "os"
  "github.com/cosbynator/wikirank/ranklib"
  "runtime"
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU()) // Mostly I/O bound, but why not

  if len(os.Args) <= 1 {
    log.Fatal("Not enough arguments")
    return
  }

  switch cmd := os.Args[1]; cmd {
  case "pagerank":
    if len(os.Args) <= 3 {
      log.Fatal("PageRank: Rrequired argument 'input.gob' / 'ranked_output.gob' missing")
    }
    inputName := os.Args[2]
    outputName := os.Args[3]
    log.Printf("Page ranking from '%s' into '%s", inputName, outputName)
    err := ranklib.RankAndWrite(inputName, outputName)
    if err != nil { panic(err) }

  case "extract_graph":
    if len(os.Args) <= 3 {
      log.Fatal("Extract: Required argument 'input.xml' / 'output.gob' missing")
      return
    }
    filename := os.Args[2]
    outputName := os.Args[3]

    log.Printf("Extracting from '%s' into '%s'", filename, outputName)
    err := ranklib.ReadFrom(filename, outputName)
    if err != nil { panic(err) }
  default:
    log.Fatalf("Unknown command '%s'", cmd)
    return
  }
  log.Printf("All done!")
}
