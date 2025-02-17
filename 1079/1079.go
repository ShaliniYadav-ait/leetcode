package main

func numTilePossibilities(tiles string) int {
    mapChar := getCharactersMap(tiles)
	
	return getTilePossibilities(mapChar)
}

func getTilePossibilities(mapChar map[byte]int)int{
	count := 0

	for key,val := range mapChar{
		if val > 0{
			count++
			mapChar[key]--
			count += getTilePossibilities(mapChar)
			mapChar[key]++
		}
	}
	return count
}

func getCharactersMap(tiles string)map[byte]int{
	mapChar := make(map[byte]int)
	for i := 0; i < len(tiles); i++{
		mapChar[tiles[i]]++
	}
	return mapChar
}