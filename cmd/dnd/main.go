package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ross/dnd/internal/dndbeyond"
	"github.com/ross/dnd/internal/hugo"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "fetch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: dnd fetch <character_id>")
			os.Exit(1)
		}
		runFetch(os.Args[2])
	case "party":
		if len(os.Args) < 3 {
			fmt.Println("Usage: dnd party <character_id>")
			os.Exit(1)
		}
		runParty(os.Args[2])
	case "sync":
		if len(os.Args) < 3 {
			fmt.Println("Usage: dnd sync <character_id>")
			os.Exit(1)
		}
		runSync(os.Args[2])
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("D&D Beyond Sync Tool")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  fetch <character_id>  - Fetch and display a character")
	fmt.Println("  party <character_id>  - List all party members")
	fmt.Println("  sync <character_id>   - Sync all party members to Hugo")
	fmt.Println()
	fmt.Println("Environment variables:")
	fmt.Println("  DND_COBALT_SESSION - CobaltSession cookie (auto-refreshes tokens)")
	fmt.Println("  DND_TOKEN          - Direct cobalt token (expires in 5 min)")
}

func getClient() (*dndbeyond.Client, error) {
	if session := os.Getenv("DND_COBALT_SESSION"); session != "" {
		auth := dndbeyond.NewAuth(session)
		token, err := auth.GetToken()
		if err != nil {
			return nil, fmt.Errorf("refreshing token: %w", err)
		}
		return dndbeyond.NewClient(token), nil
	}

	if token := os.Getenv("DND_TOKEN"); token != "" {
		return dndbeyond.NewClient(token), nil
	}

	return dndbeyond.NewClient(""), nil
}

func runFetch(characterID string) {
	client, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	char, err := client.FetchCharacter(characterID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	printCharacter(char)
}

func runParty(characterID string) {
	client, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	party, err := client.FetchParty(characterID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Party Members:")
	fmt.Println("==============")
	for _, member := range party {
		privacy := "public"
		if member.PrivacyType != 1 {
			privacy = "private"
		}
		fmt.Printf("  %s (Player: %s) - ID: %d [%s]\n",
			member.CharacterName, member.Username, member.CharacterID, privacy)
	}
}

func runSync(characterID string) {
	client, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	characters, err := client.FetchAllPartyCharacters(characterID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	contentDir := os.Getenv("DND_CONTENT_DIR")
	if contentDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
			os.Exit(1)
		}
		contentDir = filepath.Join(cwd, "content")
	}

	gen := hugo.NewGenerator(contentDir)

	fmt.Printf("Syncing %d characters to %s\n\n", len(characters), contentDir)

	for _, char := range characters {
		if err := gen.GenerateCharacter(char); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating %s: %v\n", char.Name, err)
			continue
		}
	}

	var campaign *dndbeyond.Campaign
	if len(characters) > 0 && characters[0].Campaign != nil {
		campaign = characters[0].Campaign
	}

	if err := gen.GeneratePartyIndex(characters, campaign); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating party index: %v\n", err)
	}

	fmt.Println("\nSync complete!")
}

func printCharacter(char *dndbeyond.Character) {
	fmt.Printf("Name: %s\n", char.Name)
	fmt.Printf("Player: %s\n", char.Username)

	if char.Race != nil {
		fmt.Printf("Race: %s\n", char.Race.FullName)
	}

	fmt.Printf("Class: %s (Level %d)\n", char.PrimaryClass(), char.TotalLevel())
	fmt.Printf("HP: %d/%d\n", char.CurrentHP(), char.MaxHP())

	fmt.Println("\nAbility Scores:")
	statNames := []string{"STR", "DEX", "CON", "INT", "WIS", "CHA"}
	for i, name := range statNames {
		val := char.GetStat(i + 1)
		mod := char.StatModifier(i + 1)
		sign := "+"
		if mod < 0 {
			sign = ""
		}
		fmt.Printf("  %s: %d (%s%d)\n", name, val, sign, mod)
	}

	fmt.Printf("\nGold: %d gp, %d sp, %d cp\n",
		char.Currencies.GP, char.Currencies.SP, char.Currencies.CP)

	if char.Campaign != nil {
		fmt.Printf("\nCampaign: %s (DM: %s)\n", char.Campaign.Name, char.Campaign.DMUsername)
	}
}
