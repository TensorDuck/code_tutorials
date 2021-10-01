package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type data1 struct {
	Page   int
	Fruits []string
}

type slayer struct {
	Relics            string
	Master_deck       []string
	Is_ascension_mode bool
	Floor_reached     int
}

type slayer_event struct {
	Event struct {
		Gold_per_floor    []int    `json:"gold_per_floor"`
		Floor_reached     int      `json:"floor_reached"`
		Items_purged      []string `json:"items_purged"`
		Master_deck       []string `json:"master_deck"`
		Relics            string   `json:"relics"`
		Is_ascension_mode bool     `json:"is_ascension_mode"`
		Is_prod           bool     `json:"is_prod"`
		Is_daily          bool     `json:"is_daily"`
		Is_endless        bool     `json:"is_endless"`
		Victory           bool     `json:"victory"`
		Ascension_level   int      `json:"ascension_level"`
	}
}

func simple_test() {
	// test basic parsing
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// test parsing of a specific JSON structure
	str := `{"Page": 200, "Fruits": ["apple", "pineapple"], "Test": {"value": 7}}`
	res := data1{}
	json.Unmarshal([]byte(str), &res)
	for i := 0; i < 2; i++ {
		fmt.Println(string(res.Fruits[i]))
	}
}

func simple_slayer_test() {
	new_str := `{"event": {"gold_per_floor": [99, 99], "floor_reached": 0, "playtime": 5, "items_purged": [], "score": 0, "play_id": "2eebda8a-6486-4fed-b32e-306c66ce5b52", "local_time": "20181024222406", "is_ascension_mode": true, "campfire_choices": [], "neow_cost": "", "seed_source_timestamp": 715104038123476, "circlet_count": 0, "master_deck": ["Strike_R", "Strike_R", "Strike_R", "Strike_R", "Strike_R", "Defend_R", "Defend_R", "Defend_R", "Defend_R", "Bash", "AscendersBane"], "relics": ["Burning Blood"], "potions_floor_usage": [], "damage_taken": [], "seed_played": "-4518276804806975519", "potions_obtained": [], "is_trial": false, "path_per_floor": [], "character_chosen": "IRONCLAD", "items_purchased": [], "campfire_rested": 0, "item_purchase_floors": [], "current_hp_per_floor": [68, 68], "gold": 99, "neow_bonus": "", "is_prod": false, "is_daily": false, "chose_seed": false, "campfire_upgraded": 0, "win_rate": 0, "timestamp": 1540434246, "path_taken": [], "build_version": "2018-10-23", "purchased_purges": 0, "victory": false, "max_hp_per_floor": [75, 75], "card_choices": [], "player_experience": 1360121, "relics_obtained": [], "event_choices": [], "is_beta": true, "boss_relics": [], "items_purged_floors": [], "is_endless": false, "potions_floor_spawned": [], "ascension_level": 20}}`

	new_res := slayer_event{}
	json.Unmarshal([]byte(new_str), &new_res)
	fmt.Println(new_res.Event.Master_deck)
	fmt.Println(new_res)

}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/2018-10-25-02-34#1352.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byte_value, _ := ioutil.ReadAll(jsonFile)

	var result []slayer_event
	json.Unmarshal([]byte(byte_value), &result)
	fmt.Println(result[0].Event.Master_deck)
	fmt.Println(result[0])
}
