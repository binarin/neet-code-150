// https://leetcode.com/problems/design-twitter/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

type Twitter struct {
}

func Constructor() Twitter {
	return Twitter{}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return nil
}

func (this *Twitter) Follow(followerId int, followeeId int) {
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {
	// Example 1
	twitter := Constructor()
	twitter.PostTweet(1, 5)                   // User 1 posts a new tweet (id = 5)
	fmt.Println(twitter.GetNewsFeed(1))       // Expected: [5]
	twitter.Follow(1, 2)                      // User 1 follows user 2
	twitter.PostTweet(2, 6)                   // User 2 posts a new tweet (id = 6)
	fmt.Println(twitter.GetNewsFeed(1))       // Expected: [6, 5]
	twitter.Unfollow(1, 2)                    // User 1 unfollows user 2
	fmt.Println(twitter.GetNewsFeed(1))       // Expected: [5]
}
