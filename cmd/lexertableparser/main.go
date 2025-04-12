package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type CliConfig struct {
	Input      string
	outputPath string
}

type TokenInfo struct {
	Name     string `json:"Name"`
	Priority int    `json:"Priority"`
}

type Tables struct {
	ClassifierTable map[rune]int      `json:"classifierTable"`
	TransitionTable [][]int           `json:"transitionTable"`
	TokenTypeTable  map[int]TokenInfo `json:"tokenTypeTable"`
}

func parseFlags() *CliConfig {
	var cfg CliConfig
	flag.StringVar(&cfg.outputPath, "o", "", "Output file (don't write this flag for stdout)")
	flag.StringVar(&cfg.outputPath, "output", "-", "Output file (alias)")
	flag.StringVar(&cfg.Input, "i", "tables.json", `"Config path. Default - "tables.json"`)
	flag.StringVar(&cfg.Input, "input", "tables.json", `"Config path. Default - "tables.json"`)
	flag.Parse()

	return &cfg
}

func main() {
	cfg := parseFlags()
	inFile, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer inFile.Close()

	fileData, err := io.ReadAll(inFile)
	if err != nil {
		log.Fatalf("Failed to read from input file: %v", err)
	}

	var tables Tables
	if err := json.Unmarshal(fileData, &tables); err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	var outFile *os.File
	if cfg.outputPath != "-" {
		outFile, err = os.Create(cfg.outputPath)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
	}

	var sb strings.Builder
	sb.WriteString("var ClassifierTable = map[rune]int{\n")
	for k, v := range tables.ClassifierTable {
		sb.WriteString(fmt.Sprintf("\t%v: %v,\n", k, v))
	}
	sb.WriteString("}\n\n")

	sb.WriteString("var TransitionTable = [][]int{\n")
	for _, row := range tables.TransitionTable {
		sb.WriteString("\t[]int{")
		for i, val := range row {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(strconv.Itoa(val))
		}
		sb.WriteString("},\n")
	}
	sb.WriteString("}\n\n")

	sb.WriteString("var TokenTypeTable = map[int]TokenInfo{\n")
	for k, v := range tables.TokenTypeTable {
		sb.WriteString(fmt.Sprintf("\t%v: {Name: %q, Priority: %v},\n", k, v.Name, v.Priority))
	}
	sb.WriteString("}\n")

	if outFile != nil {
		_, err = fmt.Fprint(outFile, sb.String())
		if err != nil {
			log.Fatalf("Failed to write output into file: %v", err)
		}
	} else {
		fmt.Println(sb.String())
	}
}
