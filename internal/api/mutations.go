package api

import (
	"fmt"
	"strings"
)

func hasAllKeys(params map[string]any, keys ...string) bool {
	for _, key := range keys {
		if _, ok := params[key]; !ok {
			return false
		}
	}
	return true
}

func mediaEntryUpdateMutation(params map[string]any) string {
	varDecls := []string{"$id: Int"}
	args := []string{"mediaId: $id"}

	if _, ok := params["progress"]; ok {
		varDecls = append(varDecls, "$progress: Int")
		args = append(args, "progress: $progress")
	}
	if _, ok := params["status"]; ok {
		varDecls = append(varDecls, "$status: MediaListStatus")
		args = append(args, "status: $status")
	}
	if _, ok := params["score"]; ok {
		varDecls = append(varDecls, "$score: Float")
		args = append(args, "score: $score")
	}
	if _, ok := params["notes"]; ok {
		varDecls = append(varDecls, "$notes: String")
		args = append(args, "notes: $notes")
	}

	if hasAllKeys(params, "cDate", "cMonth", "cYear") {
		varDecls = append(varDecls, "$cDate: Int", "$cMonth: Int", "$cYear: Int")
		args = append(args, "completedAt: { day: $cDate, month: $cMonth, year: $cYear }")
	}
	if hasAllKeys(params, "sDate", "sMonth", "sYear") {
		varDecls = append(varDecls, "$sDate: Int", "$sMonth: Int", "$sYear: Int")
		args = append(args, "startedAt: { day: $sDate, month: $sMonth, year: $sYear }")
	}

	return fmt.Sprintf(`mutation(%s) {
    SaveMediaListEntry(%s) {
		media {
			id
			title {
				userPreferred
			}
		}
    }
}`, strings.Join(varDecls, ", "), strings.Join(args, ", "))
}
