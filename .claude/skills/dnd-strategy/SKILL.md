---
name: dnd-strategy
description: Create tactical strategy guides for D&D encounters with consistent formatting. Use when the user wants to plan combat strategies, create tactical options, or document encounter approaches.
---

# D&D Strategy Guide Creation

Create tactical strategy guides that help players think through combat and encounter options.

## Strategy Structure

Each strategy is a folder containing:
```
content/strategy/{strategy-name}/
├── _index.md           # Situation overview + questions
├── option-a.md         # First tactical option
├── option-b.md         # Second tactical option
└── option-c.md         # Third tactical option (optional)
```

## Parent Page Format (_index.md)

```markdown
---
title: "{Encounter Name}"
date: {YYYY-MM-DD}
description: "{Brief one-line description}"
tags: ["combat", "tactics", "{relevant-tags}"]
---

## What You Know
- {Fact about the situation}
- {Fact about enemy positions}
- {Fact about environment}
- {Fact about allies/resources}

## Party Assets
- **{Character} ({Class})**: {Key abilities relevant to this encounter}
- **{Character}**: {Key abilities}
- **{Resource}**: {How it helps}

<!--more-->

## Questions to Discuss

1. {Strategic question for party discussion}
2. {Another question}
3. {Character-specific consideration}
4. {Risk/reward question}
5. {Priority question}
```

The `<!--more-->` marker separates the situation briefing from the discussion questions. The template displays option cards between these sections.

## Option Page Format

```markdown
---
title: "{Option Name}"
date: {YYYY-MM-DD}
description: "{Tagline in quotes - 3-6 words}"
tags: ["{tactic-type}", "tactics"]
weight: {1|2|3}
---

{One paragraph describing the core approach}

## Tactics

1. **{Phase 1}**: {Description}
   - {Sub-detail if needed}
   - {Sub-detail}
2. **{Phase 2}**: {Description}
3. **{Phase 3}**: {Description}

## Pros
- {Advantage}
- {Advantage}
- {Advantage}

## Cons
- {Disadvantage}
- {Disadvantage}
- {Disadvantage}

## Best If
{One sentence describing when this option is ideal}
```

## Option Naming Conventions

Use evocative names that suggest the approach:
- "The Moat Ambush" (defensive)
- "The City Assault" (offensive)
- "The Ultimatum" (diplomatic)
- "The Night Raid" (stealth)
- "The Feint" (deceptive)

## Weight Values

Control display order with `weight` in frontmatter:
- `weight: 1` - Recommended/safest option
- `weight: 2` - Alternative option
- `weight: 3` - Riskier/situational option

## Description Guidelines

The `description` field appears as a tagline on option cards:
- Keep it 3-6 words
- Use imperative or evocative phrasing
- Examples: "Make them come to us", "Take the fight to them", "Surrender or die"

## Tags

Common tags:
- Approach: `combat`, `stealth`, `diplomacy`, `defensive`, `offensive`
- Type: `tactics`, `ambush`, `siege`, `negotiation`
- Context: `bandits`, `giants`, `undead`, `{enemy-type}`

## Character-Specific Tips

If relevant, add a section for character-specific advice:

```markdown
## Ranger Tips for {Name}

- **{Ability}**: {How to use it in this scenario}
- **{Spell/Feature}**: {Tactical advice}
```

## Example

See `content/strategy/confronting-the-bandits/` for a complete example with three options.
