package linq

import (
	"testing"
)

func Test_linqForm_First(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Where(func(item int) bool {
		return item > 5
	}).First()
	if item != 6 {
		t.Error()
	}
}

func Test_linqForm_ToArray(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Where(func(item int) bool {
		return item > 5
	}).ToArray()
	if len(item) != 2 {
		t.Error()
	}
	if item[0] != 6 || item[1] != 7 {
		t.Error()
	}
}

func Test_linqForm_RemoveAll(t *testing.T) {
	lstYaml := []int{1, 2, 3}
	remove2 := From(lstYaml).RemoveAll(func(item int) bool {
		return item >= 2
	})
	if len(remove2) != 1 {
		t.Error()
	}
	if remove2[0] != 1 {
		t.Error()
	}
}

func Test_linqForm_RemoveItem(t *testing.T) {
	lstYaml := []int{1, 2, 3}
	remove2 := From(lstYaml).RemoveItem(2)
	if len(remove2) != 2 {
		t.Error()
	}
	for _, item := range remove2 {
		if item == 2 {
			t.Error()
		}
	}
}

func Test_linqForm_Count(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	count := From(lst).Where(func(item int) bool {
		return item > 5
	}).Count()

	if count != 2 {
		t.Error()
	}
}

func Test_linqForm_Any(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	isExists := From(lst).Where(func(item int) bool {
		return item > 5
	}).Any()

	if !isExists {
		t.Error()
	}
}

func Test_linqForm_All(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	isAll := From(lst).All(func(item int) bool {
		return item > 0
	})

	if !isAll {
		t.Error()
	}
}

func Test_linqForm_ToPageList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).ToPageList(3, 2)
	if item.RecordCount != int64(len(lst)) {
		t.Error()
	}
	if item.List[0] != 4 || item.List[1] != 5 || item.List[2] != 6 {
		t.Error()
	}
}

func Test_linqForm_Take(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Take(3)

	if len(item) != 3 {
		t.Error()
	}
	if item[0] != 1 || item[1] != 2 || item[2] != 3 {
		t.Error()
	}
}

func Test_linqForm_ContainsItem(t *testing.T) {
	lstYaml := []int{1, 2, 3}
	isContains := From(lstYaml).ContainsItem(4)
	if isContains {
		t.Error()
	}

	isContains = From(lstYaml).ContainsItem(3)
	if !isContains {
		t.Error()
	}
}

func Test_linqForm_Select(t *testing.T) {
	lstYaml := []string{"1", "", "2"}
	var lst []string
	From(lstYaml).Select(&lst, func(item string) any {
		return "go:" + item
	})

	if len(lst) != len(lstYaml) {
		t.Error("数量不致")
	}

	for index, yaml := range lstYaml {
		if lst[index] != "go:"+yaml {
			t.Error()
		}
	}
}

func Test_linqForm_SelectMany(t *testing.T) {
	lstYaml := [][]string{{"1", "2"}, {"3", "4"}}
	var lst []string
	From(lstYaml).SelectMany(&lst, func(item []string) any {
		return item
	})

	if len(lst) != 4 {
		t.Error("数量不致")
	}

	if lst[0] != "1" && lst[1] != "2" && lst[2] != "3" && lst[3] != "4" {
		t.Error("数据不正确")
	}
}

func Test_linqForm_SelectManyItem(t *testing.T) {
	lstYaml := [][]string{{"1", "2"}, {"3", "4"}}
	var lst []string
	From(lstYaml).SelectManyItem(&lst)

	if len(lst) != 4 {
		t.Error("数量不致")
	}

	if lst[0] != "1" && lst[1] != "2" && lst[2] != "3" && lst[3] != "4" {
		t.Error("数据不正确")
	}
}

func Test_linqForm_GroupBy(t *testing.T) {
	type testItem struct {
		name string
		age  int
	}
	lst := []testItem{{name: "steden", age: 36}, {name: "steden", age: 18}, {name: "steden2", age: 40}}
	var lstMap map[string][]testItem
	From(lst).GroupBy(&lstMap, func(item testItem) any {
		return item.name
	})

	if len(lstMap) != 2 {
		t.Error()
	}

	if len(lstMap["steden"]) != 2 {
		t.Error()
	}

	if len(lstMap["steden2"]) != 1 {
		t.Error()
	}

	lstMap = map[string][]testItem{
		"steden": {
			{name: "steden", age: 36},
			{name: "steden", age: 18},
		},
		"steden2": {
			{name: "steden2", age: 40},
		},
	}
}

func Test_linqForm_OrderBy(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := From(lst).OrderBy(func(item int) any {
		return item
	}).ToArray()
	if item[0] != 1 || item[1] != 2 || item[2] != 3 || item[3] != 4 || item[4] != 5 || item[5] != 6 || item[6] != 7 || item[7] != 8 {
		t.Error()
	}
}

func Test_linqForm_OrderByDescending(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := From(lst).OrderByDescending(func(item int) any {
		return item
	}).ToArray()
	if item[0] != 8 || item[1] != 7 || item[2] != 6 || item[3] != 5 || item[4] != 4 || item[5] != 3 || item[6] != 2 || item[7] != 1 {
		t.Error()
	}
}

func Test_linqForm_Min(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := From(lst).Min(func(item int) any {
		return item
	})
	if item.(int) != 1 {
		t.Error()
	}
}

func Test_linqForm_MinItem(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := From(lst).MinItem()
	if item != 1 {
		t.Error()
	}
}

func Test_linqForm_Max(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := From(lst).Max(func(item int) any {
		return item
	})
	if item.(int) != 8 {
		t.Error()
	}
}

func Test_linqForm_MaxItem(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := From(lst).MaxItem()
	if item != 8 {
		t.Error()
	}
}

func Test_linqForm_Sum(t *testing.T) {
	lst := []int{1, 2, 3}
	item := From(lst).Sum(func(item int) any {
		return item
	})
	if item.(int) != 6 {
		t.Error()
	}
}

func Test_linqForm_SumItem(t *testing.T) {
	lst := []int{1, 2, 3}
	item := From(lst).SumItem()
	if item != 6 {
		t.Error()
	}
}

func Test_linqForm_Avg(t *testing.T) {
	lst := []int{1, 2, 3}
	item := From(lst).Avg(func(item int) any {
		return item
	})
	if item != 2 {
		t.Error()
	}
}

func Test_linqForm_AvgItem(t *testing.T) {
	lst := []int{1, 2, 3}
	item := From(lst).AvgItem()
	if item != 2 {
		t.Error()
	}
}
