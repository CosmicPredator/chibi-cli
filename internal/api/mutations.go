package api

const mediaEntryUpdateMutation = `mutation(
    $id: Int, 
    $progress: Int,
    $score: Float,
	$notes: String,
    $cDate: Int,
    $cMonth: Int,
    $cYear: Int,
    $sDate: Int,
    $sMonth: Int,
    $sYear: Int,
    $status: MediaListStatus
) {
    SaveMediaListEntry(
        mediaId: $id, 
        progress: $progress,
        status: $status,
        score: $score,
        notes: $notes,
        completedAt: {
            day: $cDate,
            month: $cMonth,
            year: $cYear
        },
        startedAt: {
            day: $sDate,
            month: $sMonth,
            year: $sYear
        }
    ) {
		media {
			id
			title {
				userPreferred
			}
		}
    }
}`
