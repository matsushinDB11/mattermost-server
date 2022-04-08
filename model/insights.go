// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"net/http"
	"time"
)

const (
	TimeRange1Day  string = "1_day"
	TimeRange7Day  string = "7_day"
	TimeRange28Day string = "28_day"
)

type InsightsOpts struct {
	StartUnixMilli int64
	Page           int
	PerPage        int
}

type InsightsListData struct {
	HasNext bool `json:"has_next"`
}

type InsightsData struct {
	Rank int `json:"rank"`
}

// Top Reactions
type TopReactionList struct {
	InsightsListData
	Items []*TopReaction `json:"items"`
}

type TopReaction struct {
	InsightsData
	EmojiName string `json:"emoji_name"`
	Count     int64  `json:"count"`
}

// Top Channels
type TopChannelList struct {
	InsightsListData
	Items []*TopChannel `json:"items"`
}

type TopChannel struct {
	InsightsData
	ID           string      `json:"id"`
	Type         ChannelType `json:"type"`
	DisplayName  string      `json:"display_name"`
	Name         string      `json:"name"`
	MessageCount int64       `json:"message_count"`
}

type TopChannelByTimeList struct {
	InsightsListData
	Items []*TopChannelByTime `json:"items"`
}

type TopChannelByTime struct {
	Date  string        `json:"time"`
	Items []*TopChannel `json:"items"`
}

// GetStartUnixMilliForTimeRange gets the unix start time in milliseconds from the given time range.
// Time range can be one of: "1_day", "7_day", or "28_day".
func GetStartUnixMilliForTimeRange(timeRange string) (int64, *AppError) {
	switch timeRange {
	case TimeRange1Day:
		return GetMillisForTime(time.Now().Add(time.Hour * time.Duration(-24))), nil
	case TimeRange7Day:
		return GetMillisForTime(time.Now().Add(time.Hour * time.Duration(-168))), nil
	case TimeRange28Day:
		return GetMillisForTime(time.Now().Add(time.Hour * time.Duration(-672))), nil
	}

	return GetMillisForTime(time.Now()), NewAppError("Insights.IsValidRequest", "model.insights.time_range.app_error", nil, "", http.StatusBadRequest)
}

// GetTopReactionListWithRankAndPagination adds a rank to each item in the given list of TopReaction and checks if there is
// another page that can be fetched based on the given limit and offset. The given list of TopReaction is assumed to be
// sorted by Count. Returns a TopReactionList.
func GetTopReactionListWithRankAndPagination(reactions []*TopReaction, limit int, offset int) *TopReactionList {
	// Add pagination support
	var hasNext bool
	if (limit != 0) && (len(reactions) == limit+1) {
		hasNext = true
		reactions = reactions[:len(reactions)-1]
	}

	// Assign rank to each reaction
	for i, reaction := range reactions {
		reaction.Rank = offset + i + 1
	}

	return &TopReactionList{InsightsListData: InsightsListData{HasNext: hasNext}, Items: reactions}
}

// GetTopChannelListWithRankAndPagination adds a rank to each item in the given list of TopChannel and checks if there is
// another page that can be fetched based on the given limit and offset. The given list of TopChannel is assumed to be
// sorted by Score. Returns a TopChannelList.
func GetTopChannelListWithRankAndPagination(channels []*TopChannel, limit int, offset int) *TopChannelList {
	// Add pagination support
	var hasNext bool
	if (limit != 0) && (len(channels) == limit+1) {
		hasNext = true
		channels = channels[:len(channels)-1]
	}

	// Assign rank to each reaction
	for i, channel := range channels {
		channel.Rank = offset + i + 1
	}

	return &TopChannelList{InsightsListData: InsightsListData{HasNext: hasNext}, Items: channels}
}
