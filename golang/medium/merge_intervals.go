/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

package main

import (
	"math"
	"sort"
)

//Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	ret := []Interval{}

	var tmp *Interval
	for i, v := range intervals {
		if tmp == nil {
			tmp = &Interval{
				Start: v.Start,
				End:   v.End,
			}
		}
		if i < len(intervals)-1 {
			if tmp.End >= intervals[i+1].Start {
				tmp.End = int(math.Max(float64(tmp.End), float64(intervals[i+1].End)))
			} else {
				ret = append(ret, *tmp)
				tmp = nil
			}
		} else {
			ret = append(ret, *tmp)
			tmp = nil
		}
	}
	return ret
}
