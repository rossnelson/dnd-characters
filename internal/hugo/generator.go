package hugo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ross/dnd/internal/dndbeyond"
)

type Generator struct {
	contentDir string
}

func NewGenerator(contentDir string) *Generator {
	return &Generator{contentDir: contentDir}
}

func (g *Generator) GenerateCharacter(char *dndbeyond.Character) error {
	slug := slugify(char.Name)
	filename := filepath.Join(g.contentDir, "characters", slug+".md")

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	content := g.buildCharacterMarkdown(char)

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	fmt.Printf("Generated: %s\n", filename)
	return nil
}

func (g *Generator) GeneratePartyIndex(chars []*dndbeyond.Character, campaign *dndbeyond.Campaign) error {
	filename := filepath.Join(g.contentDir, "party", "_index.md")

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	content := g.buildPartyMarkdown(chars, campaign)

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	fmt.Printf("Generated: %s\n", filename)
	return nil
}

func (g *Generator) buildCharacterMarkdown(char *dndbeyond.Character) string {
	var sb strings.Builder

	race := "Unknown"
	if char.Race != nil {
		race = char.Race.FullName
	}

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: \"%s - %s %s\"\n", char.Name, race, char.PrimaryClass()))
	sb.WriteString(fmt.Sprintf("date: %s\n", time.Now().Format("2006-01-02")))
	sb.WriteString(fmt.Sprintf("description: \"%s %s - Level %d\"\n", race, char.PrimaryClass(), char.TotalLevel()))
	sb.WriteString(fmt.Sprintf("character_name: \"%s\"\n", char.Name))
	sb.WriteString(fmt.Sprintf("character_race: \"%s\"\n", race))
	sb.WriteString(fmt.Sprintf("character_class: \"%s\"\n", char.PrimaryClass()))
	sb.WriteString(fmt.Sprintf("character_level: %d\n", char.TotalLevel()))
	sb.WriteString(fmt.Sprintf("player: \"%s\"\n", char.Username))
	sb.WriteString(fmt.Sprintf("dndbeyond_id: %d\n", char.ID))
	sb.WriteString(fmt.Sprintf("tags: [\"%s\", \"%s\", \"active\"]\n",
		strings.ToLower(strings.ReplaceAll(race, " ", "-")),
		strings.ToLower(strings.ReplaceAll(char.PrimaryClass(), " ", "-"))))
	sb.WriteString(fmt.Sprintf("hp_current: %d\n", char.CurrentHP()))
	sb.WriteString(fmt.Sprintf("hp_max: %d\n", char.MaxHP()))
	sb.WriteString(fmt.Sprintf("gp: %d\n", char.Currencies.GP))
	sb.WriteString(fmt.Sprintf("sp: %d\n", char.Currencies.SP))
	sb.WriteString(fmt.Sprintf("cp: %d\n", char.Currencies.CP))
	sb.WriteString("---\n\n")

	sb.WriteString("## Ability Scores\n\n")
	sb.WriteString("| STR | DEX | CON | INT | WIS | CHA |\n")
	sb.WriteString("|:---:|:---:|:---:|:---:|:---:|:---:|\n")
	sb.WriteString(fmt.Sprintf("| %d (%s) | %d (%s) | %d (%s) | %d (%s) | %d (%s) | %d (%s) |\n",
		char.GetStat(1), formatMod(char.StatModifier(1)),
		char.GetStat(2), formatMod(char.StatModifier(2)),
		char.GetStat(3), formatMod(char.StatModifier(3)),
		char.GetStat(4), formatMod(char.StatModifier(4)),
		char.GetStat(5), formatMod(char.StatModifier(5)),
		char.GetStat(6), formatMod(char.StatModifier(6))))
	sb.WriteString("\n")

	spells := collectSpells(char)
	if len(spells) > 0 {
		sb.WriteString("## Spells\n\n")
		for _, spell := range spells {
			sb.WriteString(fmt.Sprintf("- %s\n", spell))
		}
		sb.WriteString("\n")
	}

	if char.Campaign != nil {
		sb.WriteString("## Campaign\n\n")
		sb.WriteString(fmt.Sprintf("**%s** (DM: %s)\n", char.Campaign.Name, char.Campaign.DMUsername))
	}

	return sb.String()
}

func (g *Generator) buildPartyMarkdown(chars []*dndbeyond.Character, campaign *dndbeyond.Campaign) string {
	var sb strings.Builder

	campaignName := "Unknown Campaign"
	dmName := "Unknown"
	if campaign != nil {
		campaignName = campaign.Name
		dmName = campaign.DMUsername
	}

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: \"%s\"\n", campaignName))
	sb.WriteString(fmt.Sprintf("date: %s\n", time.Now().Format("2006-01-02")))
	sb.WriteString(fmt.Sprintf("description: \"Party roster for %s\"\n", campaignName))
	sb.WriteString("---\n\n")

	sb.WriteString(fmt.Sprintf("# %s\n\n", campaignName))
	sb.WriteString(fmt.Sprintf("**DM:** %s\n\n", dmName))

	sb.WriteString("## Party Members\n\n")
	sb.WriteString("| Character | Race | Class | Level | Player |\n")
	sb.WriteString("|-----------|------|-------|:-----:|--------|\n")

	for _, char := range chars {
		race := "Unknown"
		if char.Race != nil {
			race = char.Race.FullName
		}
		slug := slugify(char.Name)
		sb.WriteString(fmt.Sprintf("| [%s](../characters/%s/) | %s | %s | %d | %s |\n",
			char.Name, slug, race, char.PrimaryClass(), char.TotalLevel(), char.Username))
	}

	return sb.String()
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "'", "")
	return s
}

func formatMod(mod int) string {
	if mod >= 0 {
		return fmt.Sprintf("+%d", mod)
	}
	return fmt.Sprintf("%d", mod)
}

func collectSpells(char *dndbeyond.Character) []string {
	var result []string
	seen := make(map[string]bool)

	addSpell := func(name string) {
		if name != "" && !seen[name] {
			result = append(result, name)
			seen[name] = true
		}
	}

	// Collect from classSpells (the main spell list for spellcasters)
	for _, cs := range char.ClassSpells {
		for _, entry := range cs.Spells {
			if entry.Definition != nil {
				addSpell(entry.Definition.Name)
			}
		}
	}

	// Collect from spells object (race, class features, feats, etc.)
	if char.Spells != nil {
		for _, entry := range char.Spells.Class {
			if entry.Definition != nil {
				addSpell(entry.Definition.Name)
			}
		}
		for _, entry := range char.Spells.Race {
			if entry.Definition != nil {
				addSpell(entry.Definition.Name)
			}
		}
		for _, entry := range char.Spells.Feat {
			if entry.Definition != nil {
				addSpell(entry.Definition.Name)
			}
		}
		for _, entry := range char.Spells.Background {
			if entry.Definition != nil {
				addSpell(entry.Definition.Name)
			}
		}
		for _, entry := range char.Spells.Item {
			if entry.Definition != nil {
				addSpell(entry.Definition.Name)
			}
		}
	}

	return result
}
