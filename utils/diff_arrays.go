package utils

func DiffArrays(oldArray, newArray []string) (added []string, removed []string) {
	oldMap := make(map[string]struct{})
	for _, item := range oldArray {
		oldMap[item] = struct{}{}
	}

	for _, item := range newArray {
		if _, exists := oldMap[item]; !exists {
			added = append(added, item)
		}
	}

	newMap := make(map[string]struct{})
	for _, item := range newArray {
		newMap[item] = struct{}{}
	}

	for _, item := range oldArray {
		if _, exists := newMap[item]; !exists {
			removed = append(removed, item)
		}
	}
	return
}
