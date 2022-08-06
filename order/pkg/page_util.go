package pkg

import "order/model/rep"

// ToPage
// @description: page_size,page_index,total_item
// @param nums
// @return page
// @2022-08-06 15:23:05
func ToPage(nums ...int64) (page rep.PageRep) {
	page.PageSize = nums[0]
	page.PageIndex = nums[1]
	page.ItemTotal = nums[2]
	page.PageTotal = nums[2] / nums[0]
	return page
}
