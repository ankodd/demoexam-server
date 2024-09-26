package models

type Statistics struct {
	CountCompletedOrders int64            `json:"count_completed_orders"`
	AverageTimeInHours   float64          `json:"average_time_in_hours"`
	CountFailuresByTypes map[string]int64 `json:"count_failures_by_types"`
}
