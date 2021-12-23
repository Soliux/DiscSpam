package gateway

import (
	"Raid-Client/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

func parseMemberData(js string, guildID string) {

	parsed := guildMemberListUpdate(js)

	if parsed.GuildID == guildID && belongsToStrSlice(parsed.Types, "SYNC") || belongsToStrSlice(parsed.Types, "UPDATE") {
		//endFetching := false
		for k, v := range parsed.Types {
			if v == "SYNC" {
				p := parsed.Updates.Get(fmt.Sprint(k)).String()
				if len(p) == 0 {
					finishedFetching <- true
				}

				// Parse members
				var m []Member
				members := parsed.Updates.Get("#.member").String()
				err := json.Unmarshal([]byte(members), &m)
				if err != nil {
					log.Println(err)
				}
				scrapedMemb = append(scrapedMemb, m...)
				utils.Logger(scrapedMemb)

			}
		}
	}

}
func guildMemberListUpdate(response string) MemberData {
	var types []string
	var locations gjson.Result
	var updates gjson.Result
	d := gjson.Get(response, "d")

	for _, chunk := range d.Get("ops").Array() {
		types = append(types, chunk.Get("op").String())

		switch chunk.Get("op").String() {
		case "SYNC", "INVALIDATE":
			locations = chunk.Get("range")

			if chunk.Get("op").String() == "SYNC" {
				updates = chunk.Get("items")
			}
		case "INSERT", "UPDATE", "DELETE":
			locations = chunk.Get("index")
			if !(chunk.Get("op").String() == "DELETE") {
				updates = chunk.Get("item")
			}
		}
	}
	memberData := MemberData{
		OnlineCount:  d.Get("online_count").String(),
		MemberCount:  d.Get("member_count").String(),
		ID:           d.Get("id").String(),
		GuildID:      d.Get("guild_id").String(),
		HoistedRoles: d.Get("groups"),
		Types:        types,
		Locations:    locations,
		Updates:      updates,
	}

	return memberData

}

// Helper function to retrieve member count
func getGuildsData(d string) {
	guilds := gjson.Get(d, "d.guilds")

	for _, guild := range guilds.Array() {
		memberCount := guild.Get("member_count")
		guildID := guild.Get("id")
		ScrapedGuilds = append(ScrapedGuilds, Guild{ID: guildID.String(), MemberCount: memberCount.String()})
	}
}
func getRanges(index, multiplier, memberCount int) [][]int { // https://github.com/Merubokkusu/Discord-S.C.U.M/blob/77daf74354415cb5d9411f886899c9817d0bc5b9/discum/gateway/guild/combo.py#L48
	initalNum := index * multiplier
	rangesList := [][]int{{initalNum, initalNum + 99}}

	if memberCount > initalNum+99 {
		rangesList = append(rangesList, []int{initalNum + 100, initalNum + 199})
	}
	if !belongToIntSlice(rangesList, []int{0, 99}) {
		rangesList = append(rangesList, []int{})
		insert(rangesList, []int{0, 99}, 0)
	}
	return rangesList
}

func getMemberCount(guildID string) string {
	var mc string
	for _, guild := range ScrapedGuilds {
		if guild.ID == guildID {
			mc = guild.MemberCount
		}
	}
	return mc
}

func belongsToStrSlice(input []string, lookup string) bool { // https://stackoverflow.com/a/52710077
	for _, val := range input {
		if val == lookup {
			return true
		}
	}
	return false
}

func belongToIntSlice(input [][]int, lookup []int) bool { // https://stackoverflow.com/a/52710077
	for _, val := range input {
		if Equal(val, lookup) {
			return true
		}
	}
	return false
}

func insert(a [][]int, c []int, i int) [][]int { //https://github.com/golang/go/wiki/SliceTricks#insert
	return append(a[:i], append([][]int{c}, a[i:]...)...)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
