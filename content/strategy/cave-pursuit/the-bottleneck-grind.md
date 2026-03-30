---
title: "Open Cavern Combat"
date: 2026-03-27
description: "Fighting without the bottleneck - flanking threats, open terrain, and how to adapt"
tags: ["combat", "tactics", "cave", "open-terrain"]
weight: 2
---

The bottleneck is behind you. The cave geometry that made the first fight manageable is gone. In open terrain, enemies can spread out, flank, and attack from multiple directions simultaneously. Kaltor can no longer single-handedly block the entire enemy force. The party has to fight smarter.

## What Changed

In the bottleneck:
- Kaltor blocked everything. One enemy could reach the party at a time
- Action economy massively favored the party
- Spellcasters could fire freely with zero risk

In an open cavern:
- Enemies can attack from the left, right, and front at the same time
- Multiple enemies can reach Barzer and Selena if the line breaks
- Flanking grants Advantage to enemies (if using optional rule) or simply puts squishies in danger

**The core tactical principle shifts**: Instead of "hold the chokepoint", the new principle is **"create local superiority at the point of contact while protecting the back line"**.

## Core Positioning Rules in Open Terrain

1. **Kaltor does NOT stand alone at the front**. In open terrain, Nysera flanks him on one side to prevent enemies from going around him
2. **Barzer stays 15-20ft behind the front line**, not adjacent to Kaltor. His Shield reaction and Magic Missile work from distance
3. **Elaris and Selena stay at max Healing Word range (60ft) from the front line** - they should never be in melee range unless things have gone catastrophically wrong
4. **Theren positions for Longbow fire**: 10-20ft behind Kaltor, clear sightline, mobile enough to step sideways if flanked

## Tactical Combos for Open Terrain

### Situation A: The Missing Ogre Charges From Ahead

Standard engagement - the ogre comes straight at you.

| Round | Priority Action |
|-------|----------------|
| 1 | Kaltor intercepts. Theren fires Longbow (Hunter's Mark + Colossus Slayer). Elaris drops Bless on Kaltor + Theren + Nysera. Nysera casts Moonbeam centered on the ogre |
| 2 | Kaltor attacks. Theren Longbow. Nysera moves Moonbeam (bonus action) to stay on target. Barzer holds Magic Missile for kill-shot when ogre drops below 15 HP |
| 3 | Ogre should be near dead. Barzer fires Magic Missile to guarantee the kill. Selena Healing Word on Theren or Nysera if they took hits |

**Do not let this become a prolonged fight.** One ogre against the full party should end in 2-3 rounds.

---

### Situation B: Ogre Tries to Go Around Kaltor

Open terrain means the ogre might not play along and walk into Kaltor's sword. It may try to reach Barzer or Elaris directly.

- **Nysera shifts position** to cut off the flank - Flaming Sphere placed between the ogre's likely path and the back line forces it to divert or take 2d6 fire
- **Theren uses Ready Action**: "When the ogre moves past Kaltor, I fire." This triggers before it reaches the back line
- **Kaltor uses his reaction for an Opportunity Attack** as it passes - that longsword hits hard and may stop the charge
- **Barzer's Shield reaction**: If the ogre reaches Barzer and attacks, Shield brings AC to 18. Use it. Barzer surviving is worth the slot

---

### Situation C: Multiple Enemies From Multiple Directions

This is the worst case. Two ogres splitting around the party, or ogre plus goblins, or anything that attacks from more than one direction simultaneously.

**Recognition**: If Theren's scouting spots two separate approach vectors, call it immediately. This is not a "stand and fight" situation by default.

**Options if you choose to fight:**

1. **Flaming Sphere as a wall**: Nysera places it to block one approach vector entirely. Enemies coming from that direction must go through 2d6 fire or stop. Kaltor handles the other vector
2. **Elaris drops Hideous Laughter** on the most dangerous enemy among the secondary group - removes it from the fight for as long as concentration holds
3. **Barzer uses Fog Cloud** to cut off the secondary approach - enemies in the fog are blind, they'll move slowly and unpredictably. This buys time to finish the primary group first
4. **Theren and Selena watch the flanks** - Selena uses Sacred Flame at range, Theren fires Longbow at whoever is most likely to reach the back line

**If there are 3+ enemies approaching**: See the retreat document. Open terrain against 3+ foes on a hurt party is not a winnable fight on current resources.

---

### Situation D: Ambush - Enemies Were Waiting for You

The missing ogre has had time to prepare. It may have positioned itself in your blind zone (beyond 60ft) and charged when you entered its range.

**Surprise round**: If the party is surprised, you lose your first round of actions. This is why Theren scouts ahead - a successful Perception check prevents the surprise.

**If surprised anyway:**
- **Kaltor's first non-surprised turn**: Interpose himself between the threat and the back line
- **Barzer's first turn**: Shield reaction if hit, otherwise Magic Missile at the attacker - the auto-hit matters when you're disoriented
- **Selena**: Sanctuary on Barzer immediately (4 HP even after healing is still critical - protect him first)
- **Theren and Nysera**: Strike hard in round 1 to prevent the fight from going to round 3+

---

### Situation E: Finding Human Captives

Stop. Do not advance past them.

1. **Elaris moves to captives** - Healing Word if any are injured
2. **Selena casts Sanctuary** on the most injured captive - protects them from targeting
3. **Kaltor and Theren hold** the direction of deeper threats
4. **Nysera Wild Shape** into a non-threatening form if captives are panicked - a druid-wolf is not reassuring
5. **Barzer covers the rear** with Fog Cloud ready if the party extracts
6. **Do not push further toward the infernal source with captives present** - extract them first

---

## Engagement Rules: Open Terrain Version

The bottleneck rules were generous. Open terrain is less forgiving.

| What You Find | Fight or Withdraw? |
|---------------|-------------------|
| 1 ogre (the missing one) | Fight - full party, manageable |
| 1 ogre + 2-3 goblins | Fight if Theren identifies them first and the party can position |
| 2 fresh ogres | Fight only if there's a defensible position - otherwise withdraw |
| 3+ enemies in open space | Withdraw - resource cost too high |
| Unidentified infernal entity | Identify before engaging. Unknown = withdraw |
| Captives present | Secure captives, then reassess |

---

## The Infernal Source Decision Tree

```
Something infernal is ahead
        |
        ├── Can you identify it? (Barzer: Arcana, Theren: Nature/Religion)
        |         |
        |         ├── YES: Minor fiend / imp / low CR → Fight if party is stable
        |         └── NO: Unknown → Do not engage, retreat and identify
        |
        └── Is the party at fighting strength?
                  |
                  ├── Barzer above 10 HP, Theren above 12 HP, Nysera above 12 HP?
                  |         |
                  |         ├── YES → Proceed carefully with terrain assessment
                  |         └── NO → Retreat for resources first
                  |
                  └── Are captives present?
                            |
                            ├── YES → Extract before engaging the infernal source
                            └── NO → Proceed per resource check above
```

---

## Concentration Management in Open Terrain

Multiple concentration spells competing is the key internal resource conflict.

| Caster | Best Concentration Spell | Notes |
|--------|--------------------------|-------|
| Theren | Hunter's Mark | On the primary target at all times |
| Elaris | Bless | Cast first, maintain all fight - the +1d4 on Kaltor/Theren/Nysera attacks is worth more than any one offensive spell |
| Nysera | Moonbeam (single target) or Flaming Sphere (area control) | Moonbeam vs. a single ogre; Sphere when you need to block a route or cover area |
| Barzer | No concentration needed | Magic Missile is instant, no concentration |

**The critical rule**: Nysera should not cast Moonbeam AND Bless. Elaris handles Bless. Nysera handles area/damage concentration. If Elaris goes down or loses concentration, Nysera can pick up Bless - but that's an emergency pivot, not the plan.

---

## Pros
- Full party action economy
- Universal darkvision means no light management
- Kaltor at full HP is a genuine frontline anchor even in open terrain
- Nysera's Flaming Sphere creates area denial without requiring a chokepoint

## Cons
- No chokepoint means enemies can reach the back line if the front line is bypassed
- Barzer remains fragile - even after healing he can't absorb hits
- Theren and Nysera are hurt - a bad round on either could take them out of the fight
- Open terrain means no guaranteed angle for ranged fire - Theren may have to move each round for clean shots

## Best If
Theren spots the threat before it reaches melee range, the party has time to position, and Kaltor/Nysera establish a front line before enemies close. Speed of positioning is the key variable in open terrain.
