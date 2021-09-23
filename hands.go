// Copyright (C) 2021  Antonio Lassandro

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.

// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"math/rand"
	"strings"
)

var Lights = []string{
	// Authentic
	"Cherry Blossom",
	"Sun",
	"Moon",
	"Rain",
	"Willow",
	"Dock",
	"Curtain",
	"Rainman",
	"Phoenix",

	// Other
	"Spring",
	"Summer",
	"Winter",
	"Autumn",
	"Valley",
	"River",
	"Ocean",
	"Sea",
	"Heaven",
	"Mountain",
	"Plum Blossom",
	"Rose",
	"Chrysanthemum",
	"Pine",
	"Oak",
}

var Seeds = []string{
	// Authentic
	"Boar",
	"Deer",
	"Butterfly",
	"Crane",
	"Songbird",
	"Cuckoo",
	"Geese",
	"Cup",
	"Swallow",
	"Lightning",

	// Other
	"Elephant",
	"Tiger",
	"Lion",
	"Monkey",
	"Bat",
	"Fish",
	"Spider",
	"Centipede",
	"Caterpillar",
}

var Slips = []string{
	// Authentic
	"Poetry Slips",

	// Other
	"Scrolls",
	"Prose Slips",
	"Letters",
	"Tablets",
	"Books",
	"Novellas",
}

var Colors = []string{
	"Transparent Black",
	"Transparent White",
	"Alice Blue",
	"Antique White",
	"Aqua",
	"Aquamarine",
	"Azure",
	"Beige",
	"Bisque",
	"Black",
	"Blanched Almond",
	"Blue",
	"Blue Violet",
	"Brown",
	"Burlywood",
	"Cadet Blue",
	"Chartreuse",
	"Chocolate",
	"Coral",
	"Cornflower Blue",
	"Cornsilk",
	"Crimson",
	"Cyan",
	"Dark Blue",
	"Dark Cyan",
	"Dark Goldenrod",
	"Dark Grey",
	"Dark Green",
	"Dark Khaki",
	"Dark Magenta",
	"Dark Olive Green",
	"Dark Orange",
	"Dark Orchid",
	"Dark Red",
	"Dark Salmon",
	"Dark Sea Green",
	"Dark Slate Blue",
	"Dark Slate Grey",
	"Dark Turquoise",
	"Dark Violet",
	"Deep Pink",
	"Deep Sky Blue",
	"Dim Grey",
	"Dodger Blue",
	"Firebrick",
	"Floral White",
	"Forest Green",
	"Fuchsia",
	"Gainsboro",
	"Ghost White",
	"Gold",
	"Goldenrod",
	"Grey",
	"Green",
	"Green Yellow",
	"Honeydew",
	"Hotpink",
	"Indian Red",
	"Indigo",
	"Ivory",
	"Khaki",
	"Lavender",
	"Lavender Blush",
	"Lawn Green",
	"Lemonchiffon",
	"Light Blue",
	"Light Coral",
	"Light Cyan",
	"Light Goldenrod Yellow",
	"Light Green",
	"Light Grey",
	"Light Pink",
	"Light Salmon",
	"Light Sea Green",
	"Light Sky Blue",
	"Light Slate Grey",
	"Light Steel Blue",
	"Light Yellow",
	"Lime",
	"Lime Green",
	"Linen",
	"Magenta",
	"Maroon",
	"Medium Aquamarine",
	"Medium Blue",
	"Medium Orchid",
	"Medium Purple",
	"Medium Sea Green",
	"Medium Slate Blue",
	"Medium Spring Green",
	"Medium Turquoise",
	"Medium Violet Red",
	"Midnight Blue",
	"Mint Cream",
	"Misty Rose",
	"Moccasin",
	"Navajo White",
	"Navy",
	"Old Lace",
	"Olive",
	"Olive Drab",
	"Orange",
	"Orange Red",
	"Orchid",
	"Pale Goldenrod",
	"Pale Green",
	"Pale Turquoise",
	"Pale Violet Red",
	"Papaya Whip",
	"Peach Puff",
	"Peru",
	"Pink",
	"Plum",
	"Powder Blue",
	"Purple",
	"Rebecca Purple",
	"Red",
	"Rosy Brown",
	"Royal Blue",
	"Saddle Brown",
	"Salmon",
	"Sandy Brown",
	"Sea Green",
	"Seashell",
	"Sienna",
	"Silver",
	"Sky Blue",
	"Slate Blue",
	"Slate Grey",
	"Snow",
	"Spring Green",
	"Steel Blue",
	"Tan",
	"Teal",
	"Thistle",
	"Tomato",
	"Turquoise",
	"Violet",
	"Wheat",
	"White",
	"White Smoke",
	"Yellow",
	"Yellow Green",
}

func randchance(chance int) bool {
	return rand.Intn(chance) == chance-1
}

func randbool() bool {
	return rand.Intn(2) == 1
}

func randchoice(options []string) string {
	return options[rand.Intn(len(options)-1)]
}

func LightHand() []string {
	var result []string

	/* Viewing Hand */
	if randbool() {
		result = append(result, randchoice(Lights))
		result = append(result, "Viewing")
		/* Numeric Hand */
	} else {
		maxNum := 10

		/* Absurd numeric */
		if randchance(100) {
			maxNum = 1000000000000000
			/* High Numeric */
		} else if randchance(10) {
			maxNum = 1000
		}

		numLights := rand.Intn(maxNum)

		result = append(result, NumToEnglish(numLights))

		if numLights == 1 {
			result = append(result, "Light")
		} else {
			result = append(result, "Lights")
		}
	}

	return result
}

func SeedHand() []string {
	var result []string

	/* Just Seeds */
	if randbool() {
		result = append(result, "Seeds")
	} else {
		for i := 0; i < rand.Intn(3)+1; i++ {
			result = append(result, randchoice(Seeds))
		}
	}

	return result
}

func SlipHand() []string {
	var result []string

	/* Colored Slip */
	if randbool() {
		result = append(result, randchoice(Colors))
	}

	result = append(result, randchoice(Slips))

	return result
}

func ChaffHand() []string {
	return []string{"Chaff"}
}

func MakeHand() string {
	type HandFunc func() []string

	hands := []HandFunc{
		LightHand,
		SlipHand,
		SeedHand,
		ChaffHand,
	}

	result := hands[rand.Intn(len(hands)-1)]()

	/* 'With' Hand */
	if randbool() {
		result = append(result, "With")
		result = append(result, randchoice(Lights))
	}

	return strings.Join(result, " ")
}
