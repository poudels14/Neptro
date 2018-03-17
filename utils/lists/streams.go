package lists

import "strings"

type Predicate func(string) bool

// Filters everything in the given list that matches the predicate
func FilterStringList(list []string, predicate Predicate) []string {
	var filtered []string
	for _, v := range list {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Filters out all empty strings from the given list
func RemoveEmptyStrings(list []string) []string {
	predicate := func(s string) bool { return strings.TrimSpace(s) != "" }
	return FilterStringList(list, predicate)
}

type Mapper func(string) string

// Maps the given list using the mapper
func MapStringList(list []string, mapper Mapper) []string {
	var mapped []string
	for _, v := range list {
		mapped = append(mapped, mapper(v))
	}
	return mapped
}

type Fold func(string, string) string

func FoldStringList(list []string, acc string, fold Fold) string {
	for _, v := range list {
		acc = fold(acc, v)
	}
	return acc
}
