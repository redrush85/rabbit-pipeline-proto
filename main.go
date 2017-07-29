package main

import "fmt"
import "time"

func main() {
	start := time.Now()

	msg := []byte(`{"session_id": "na:68479082:mid:phase1_1:1501341830", "phase_name": "phase1_1", "role": "mid", "summoner_id": 68479082, "match_list": [{"match_id": 2522862401, "timestamp": 1497306913361}, {"match_id": 2520880707, "timestamp": 1497074539879}, {"match_id": 2537976292, "timestamp": 1498881774123}, {"match_id": 2522865957, "timestamp": 1497308760412}, {"match_id": 2522828741, "timestamp": 1497305127675}, {"match_id": 2521988103, "timestamp": 1497221776269}, {"match_id": 2522219816, "timestamp": 1497232090701}, {"match_id": 2520685223, "timestamp": 1497071659804}, {"match_id": 2538072059, "timestamp": 1498890146972}, {"match_id": 2522224299, "timestamp": 1497234102377}, {"match_id": 2520044493, "timestamp": 1496981503761}, {"match_id": 2522090638, "timestamp": 1497227866761}, {"match_id": 2521563887, "timestamp": 1497155686267}, {"match_id": 2520088978, "timestamp": 1496989494383}, {"match_id": 2521587859, "timestamp": 1497157863967}, {"match_id": 2520660958, "timestamp": 1497069405375}, {"match_id": 2521515575, "timestamp": 1497151925447}, {"match_id": 2520084538, "timestamp": 1496986566276}, {"match_id": 2521418683, "timestamp": 1497144135625}, {"match_id": 2520531646, "timestamp": 1497060076167}], "account_id": 229038635, "channel_id": "d3268b4d-9133-457e-a8d8-38356124eb5f"}`)
	task := Task{concurrency: 10}
	task.Process(msg)

	fmt.Println("Execution time: ", time.Since(start))
}
