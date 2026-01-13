package dndbeyond

type CharacterResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    Character `json:"data"`
}

type Character struct {
	ID                int              `json:"id"`
	UserID            int              `json:"userId"`
	Username          string           `json:"username"`
	ReadonlyURL       string           `json:"readonlyUrl"`
	Name              string           `json:"name"`
	Race              *Race            `json:"race"`
	Classes           []Class          `json:"classes"`
	Background        *Background      `json:"background"`
	Stats             []Stat           `json:"stats"`
	BaseHitPoints     int              `json:"baseHitPoints"`
	BonusHitPoints    int              `json:"bonusHitPoints"`
	RemovedHitPoints  int              `json:"removedHitPoints"`
	TemporaryHitPoints int             `json:"temporaryHitPoints"`
	CurrentXP         int              `json:"currentXp"`
	AlignmentID       int              `json:"alignmentId"`
	Age               any              `json:"age"`
	Hair              any              `json:"hair"`
	Eyes              any              `json:"eyes"`
	Skin              any              `json:"skin"`
	Height            any              `json:"height"`
	Weight            any              `json:"weight"`
	Faith             any              `json:"faith"`
	Gender            any              `json:"gender"`
	Inventory         []InventoryItem  `json:"inventory"`
	Currencies        Currencies       `json:"currencies"`
	Spells            *Spells          `json:"spells"`
	ClassSpells       []ClassSpells    `json:"classSpells"`
	Campaign          *Campaign        `json:"campaign"`
	Notes             *Notes           `json:"notes"`
	Traits            *Traits          `json:"traits"`
	Preferences       *Preferences     `json:"preferences"`
	DeathSaves        *DeathSaves      `json:"deathSaves"`
}

type Race struct {
	FullName            string `json:"fullName"`
	BaseRaceName        string `json:"baseRaceName"`
	SubRaceShortName    string `json:"subRaceShortName"`
	IsHomebrew          bool   `json:"isHomebrew"`
	WeightSpeeds        *WeightSpeeds `json:"weightSpeeds"`
}

type WeightSpeeds struct {
	Normal *SpeedInfo `json:"normal"`
}

type SpeedInfo struct {
	Walk int `json:"walk"`
}

type Class struct {
	ID             int     `json:"id"`
	Level          int     `json:"level"`
	Definition     *ClassDefinition `json:"definition"`
	SubclassDefinition *SubclassDefinition `json:"subclassDefinition"`
}

type ClassDefinition struct {
	Name       string `json:"name"`
	HitDice    int    `json:"hitDice"`
}

type SubclassDefinition struct {
	Name string `json:"name"`
}

type Background struct {
	Definition *BackgroundDefinition `json:"definition"`
}

type BackgroundDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Stat struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

type InventoryItem struct {
	ID         int                 `json:"id"`
	Definition *ItemDefinition     `json:"definition"`
	Equipped   bool                `json:"equipped"`
	Quantity   int                 `json:"quantity"`
}

type ItemDefinition struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	ArmorClass  int    `json:"armorClass"`
	Damage      *Damage `json:"damage"`
	Description string `json:"description"`
}

type Damage struct {
	DiceString string `json:"diceString"`
	DamageType string `json:"damageType"`
}

type Currencies struct {
	CP int `json:"cp"`
	SP int `json:"sp"`
	EP int `json:"ep"`
	GP int `json:"gp"`
	PP int `json:"pp"`
}

type Spells struct {
	Race       []SpellEntry `json:"race"`
	Class      []SpellEntry `json:"class"`
	Item       []SpellEntry `json:"item"`
	Feat       []SpellEntry `json:"feat"`
	Background []SpellEntry `json:"background"`
}

type ClassSpells struct {
	CharacterClassID int          `json:"characterClassId"`
	Spells           []SpellEntry `json:"spells"`
}

type SpellEntry struct {
	ID         int              `json:"id"`
	Definition *SpellDefinition `json:"definition"`
	Prepared   bool             `json:"prepared"`
	AlwaysPrepared bool         `json:"alwaysPrepared"`
}

type SpellDefinition struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	School      string `json:"school"`
	Description string `json:"description"`
	Range       *SpellRange `json:"range"`
	Duration    *SpellDuration `json:"duration"`
	Ritual      bool   `json:"ritual"`
}

type SpellRange struct {
	Origin       string `json:"origin"`
	RangeValue   int    `json:"rangeValue"`
}

type SpellDuration struct {
	DurationType string `json:"durationType"`
	DurationInterval int `json:"durationInterval"`
}

type Campaign struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	PublicNotes string        `json:"publicNotes"`
	DMID        int           `json:"dmId"`
	DMUsername  string        `json:"dmUsername"`
	Characters  []PartyMember `json:"characters"`
}

type PartyMember struct {
	UserID        int    `json:"userId"`
	Username      string `json:"username"`
	CharacterID   int    `json:"characterId"`
	CharacterName string `json:"characterName"`
	CharacterURL  string `json:"characterUrl"`
	AvatarURL     string `json:"avatarUrl"`
	PrivacyType   int    `json:"privacyType"`
	IsAssigned    bool   `json:"isAssigned"`
}

type Notes struct {
	Allies             string `json:"allies"`
	PersonalPossessions string `json:"personalPossessions"`
	OtherHoldings      string `json:"otherHoldings"`
	Organizations      string `json:"organizations"`
	Enemies            string `json:"enemies"`
	Backstory          string `json:"backstory"`
	OtherNotes         string `json:"otherNotes"`
}

type Traits struct {
	PersonalityTraits string `json:"personalityTraits"`
	Ideals            string `json:"ideals"`
	Bonds             string `json:"bonds"`
	Flaws             string `json:"flaws"`
	Appearance        string `json:"appearance"`
}

type Preferences struct {
	UseHomebrewContent bool `json:"useHomebrewContent"`
}

type DeathSaves struct {
	FailCount    int  `json:"failCount"`
	SuccessCount int  `json:"successCount"`
	IsStabilized bool `json:"isStabilized"`
}

func (c *Character) TotalLevel() int {
	total := 0
	for _, class := range c.Classes {
		total += class.Level
	}
	return total
}

func (c *Character) PrimaryClass() string {
	if len(c.Classes) == 0 {
		return "Unknown"
	}
	class := c.Classes[0]
	name := ""
	if class.Definition != nil {
		name = class.Definition.Name
	}
	if class.SubclassDefinition != nil && class.SubclassDefinition.Name != "" {
		name = class.SubclassDefinition.Name + " " + name
	}
	return name
}

func (c *Character) MaxHP() int {
	return c.BaseHitPoints + c.BonusHitPoints
}

func (c *Character) CurrentHP() int {
	return c.MaxHP() - c.RemovedHitPoints
}

func (c *Character) GetStat(statID int) int {
	for _, stat := range c.Stats {
		if stat.ID == statID {
			return stat.Value
		}
	}
	return 10
}

func (c *Character) StatModifier(statID int) int {
	return (c.GetStat(statID) - 10) / 2
}

const (
	StatStrength     = 1
	StatDexterity    = 2
	StatConstitution = 3
	StatIntelligence = 4
	StatWisdom       = 5
	StatCharisma     = 6
)
