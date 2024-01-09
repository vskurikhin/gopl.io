// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
// Напишите функцию, которая без выделения дополнительной памяти
// удаляет все смежные дубликаты в срезе []string.

package main

import "github.com/vskurikhin/gotool"

func Deduplicate(strings []string) []string {

	var result = strings[0:min(1, len(strings))]
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[i-1] {
			result = append(result, strings[i])
		}
	}
	return result
}

func main() {
	got := Deduplicate([]string{"a", "a", "b", "b", "c", "c"})
	gotool.AssertEqual(nil, got, []string{"a", "b", "c"})

}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
