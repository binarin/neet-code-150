package main

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	twitter := Constructor()

	// User 1 posts a new tweet (id = 5)
	twitter.PostTweet(1, 5)

	// User 1's news feed should return a list with 1 tweet id -> [5]
	got := twitter.GetNewsFeed(1)
	want := []int{5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed(1) after PostTweet(1,5) = %v, want %v", got, want)
	}

	// User 1 follows user 2
	twitter.Follow(1, 2)

	// User 2 posts a new tweet (id = 6)
	twitter.PostTweet(2, 6)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5]
	got = twitter.GetNewsFeed(1)
	want = []int{6, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed(1) after follow and PostTweet(2,6) = %v, want %v", got, want)
	}

	// User 1 unfollows user 2
	twitter.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5]
	got = twitter.GetNewsFeed(1)
	want = []int{5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed(1) after unfollow = %v, want %v", got, want)
	}
}
