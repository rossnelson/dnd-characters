---
name: dnd-character-import
description: Import D&D character data from D&D Beyond API and generate Hugo content pages. Use when the user wants to sync characters, import character data, or update character information from D&D Beyond.
---

# D&D Character Import

Import character data from D&D Beyond and generate Hugo markdown pages.

## Prerequisites

- Go installed (for building the CLI tool)
- CobaltSession token from D&D Beyond (for private characters)

## Getting the CobaltSession Token

1. Log into D&D Beyond in your browser
2. Open Developer Tools (F12)
3. Go to Application > Cookies > dndbeyond.com
4. Copy the value of `CobaltSession`

The token expires periodically and needs to be refreshed.

## Running the Import

```bash
# From the project root
go run ./cmd/dnd sync --cobalt="YOUR_COBALT_SESSION_TOKEN"
```

This will:
1. Fetch character data from the D&D Beyond API
2. Generate Hugo markdown files in `content/characters/`

## Character IDs

Characters are configured in `cmd/dnd/main.go`. The current party:
- Theren Vale (ID: 137344936)
- Barzer (ID: 138486583)
- Kaltor Reinhardt (ID: 137427042)
- Nysera (ID: 137436498)
- Elaris Moonveil (ID: 137533498)
- Selena (ID: 137615566)

## D&D Beyond API

The unofficial API endpoint:
```
https://character-service.dndbeyond.com/character/v5/character/{CHARACTER_ID}
```

Authentication requires the CobaltSession cookie in headers.

## Generated Content Structure

Each character page includes:
- **Frontmatter**: name, race, class, level, HP, wealth, abilities, skills, proficiencies
- **Body**: Background, personality traits, ideals, bonds, flaws, features, spells

## Key Files

- `cmd/dnd/main.go` - CLI entry point with character IDs
- `internal/dndbeyond/client.go` - API client
- `internal/dndbeyond/types.go` - Data structures
- `internal/hugo/generator.go` - Markdown generation

## Troubleshooting

**401 Unauthorized**: CobaltSession token expired. Get a fresh token from browser cookies.

**Missing spells**: Spells come from both `classSpells` (main spell list) and `spells` (race/feat spells) in the API response.
