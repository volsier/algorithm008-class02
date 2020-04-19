package main

import "sort"

/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 *
 * https://leetcode-cn.com/problems/design-twitter/description/
 *
 * algorithms
 * Medium (36.67%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 24K
 * Testcase Example:  '["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]\n[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]'
 *
 *
 * 设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：
 *
 *
 * postTweet(userId, tweetId): 创建一条新的推文
 * getNewsFeed(userId):
 * 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
 * follow(followerId, followeeId): 关注一个用户
 * unfollow(followerId, followeeId): 取消关注一个用户
 *
 *
 * 示例:
 *
 *
 * Twitter twitter = new Twitter();
 *
 * // 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
 * twitter.postTweet(1, 5);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * twitter.getNewsFeed(1);
 *
 * // 用户1关注了用户2.
 * twitter.follow(1, 2);
 *
 * // 用户2发送了一个新推文 (推文id = 6).
 * twitter.postTweet(2, 6);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
 * // 推文id6应当在推文id5之前，因为它是在5之后发送的.
 * twitter.getNewsFeed(1);
 *
 * // 用户1取消关注了用户2.
 * twitter.unfollow(1, 2);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * // 因为用户1已经不再关注用户2.
 * twitter.getNewsFeed(1);
 *
 *
 */

// @lc code=start
type Tweet struct {
	Id int
}
type User struct {
	Id int
}
type Twitter struct {
	Tweets    []Tweet
	UserFeeds map[int][]int        //存储每个用户发布的推文索引
	Follows   map[int]map[int]User //关注关系 map键代表用户id 对应的切片包含所有关注的用户id
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Tweets:    make([]Tweet, 0),
		UserFeeds: make(map[int][]int),
		Follows:   make(map[int]map[int]User),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.UserFeeds[userId] == nil {
		this.UserFeeds[userId] = make([]int, 0)
	}
	this.UserFeeds[userId] = append(this.UserFeeds[userId], len(this.Tweets))
	this.Tweets = append(this.Tweets, Tweet{Id: tweetId})

}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	users, uOk := this.Follows[userId]
	if !uOk {
		users = make(map[int]User)
	}
	users[userId] = User{Id: userId}
	feeds := make([]int, 0)
	for _, u := range users {
		feeds = append(feeds, this.UserFeeds[u.Id]...)
	}
	sort.Ints(feeds)
	l := len(feeds) - 1
	findFeeds := make([]int, 0)
	for i := l; i >= 0; i-- {
		findFeeds = append(findFeeds, this.Tweets[feeds[i]].Id)
		if len(findFeeds) == 10 {
			break
		}
	}

	return findFeeds
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	//判断是否重复关注
	if this.Follows[followerId] == nil {
		this.Follows[followerId] = make(map[int]User)
	}
	_, ok := this.Follows[followerId][followeeId]
	if !ok && followerId != followeeId {
		this.Follows[followerId][followeeId] = User{Id: followeeId}
	}

}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	//判断是否重复关注
	_, ok := this.Follows[followerId][followeeId]
	if ok {
		delete(this.Follows[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end
