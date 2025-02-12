package quizz

import (
	"math/rand"
	"time"
)

type Quizzer struct {
	list []Quizz
}

func NewQuizzer() Quizzer {
	return Quizzer{
		list: []Quizz{
			{Text: "Is 2 greater than 4?", Response: false},
			{Text: "Is 2 equal to 4?", Response: false},
			{Text: "Is 2 less than 4?", Response: true},
			{Text: "Is 4 greater than or equal to 4?", Response: true},
			{Text: "Does 5 + 3 = 8?", Response: true},
			{Text: "Does 10 - 7 = 4?", Response: false},
			{Text: "Does 6 * 7 = 42?", Response: true},
			{Text: "Does 12 / 4 = 3?", Response: true},
			{Text: "Does 15 % 4 = 3?", Response: true},
			{Text: "Does 9 + 8 = 17?", Response: true},
			{Text: "Does 25 / 5 = 6?", Response: false},
			{Text: "Is the capital of Spain Madrid?", Response: true},
			{Text: "Is the Atlantic Ocean larger than the Pacific Ocean?", Response: false},
			{Text: "Are dogs mammals?", Response: true},
			{Text: "Is the sun a planet?", Response: false},
			{Text: "Are whales fish?", Response: false},
			{Text: "Does a year have 365 days?", Response: true},
			{Text: "Does water boil at 100°C?", Response: true},
			{Text: "Is the sum of the internal angles of a triangle 180°?", Response: true},
			{Text: "Was the Great Wall of China built in the 20th century?", Response: false},
			{Text: "Does the human body have 206 bones?", Response: true},
			{Text: "Is the North Pole a continent?", Response: false},
			{Text: "Does chocolate come from cocoa?", Response: true},
			{Text: "Was the Mona Lisa painted by Leonardo da Vinci?", Response: true},
			{Text: "Is Everest the highest mountain in the world?", Response: true},
			{Text: "Is electricity measured in meters?", Response: false},
			{Text: "Is the largest ocean in the world the Pacific?", Response: true},
			{Text: "Do giraffes have short necks?", Response: false},
			{Text: "Is iron a metal?", Response: true},
			{Text: "Is a tomato a fruit?", Response: true},
			{Text: "Does 2 x 9 = 18?", Response: true},
			{Text: "Does a kilometer have 500 meters?", Response: false},
			{Text: "Is water composed of hydrogen and oxygen?", Response: true},
			{Text: "Is the Eiffel Tower in Rome?", Response: false},
			{Text: "Can fish breathe out of water?", Response: false},
			{Text: "Does a pentagon have 5 sides?", Response: true},
			{Text: "Does DNA contain the genetic information of living beings?", Response: true},
			{Text: "Is the Moon a natural satellite of the Earth?", Response: true},
			{Text: "Does a century have 50 years?", Response: false},
			{Text: "Is the color of the sky blue due to the dispersion of light?", Response: true},
			{Text: "Is air a mixture of gases?", Response: true},
			{Text: "Are dolphins mammals?", Response: true},
			{Text: "Do humans have 3 lungs?", Response: false},
			{Text: "Is oxygen necessary for human respiration?", Response: true},
			{Text: "Does solar energy come from the sun?", Response: true},
			{Text: "Are spiders insects?", Response: false},
			{Text: "Is aluminum a metal?", Response: true},
			{Text: "Does sound travel faster in air than in water?", Response: false},
			{Text: "Do volcanoes eject lava?", Response: true},
			{Text: "Is Antarctica the warmest continent on the planet?", Response: false},
			{Text: "Is sugar a carbohydrate?", Response: true},
			{Text: "Is ice water in a solid state?", Response: true},
			{Text: "Is mercury a liquid metal at room temperature?", Response: true},
			{Text: "Is the Nile the longest river in the world?", Response: false}, // Changed to false
			{Text: "Is the capital of France Paris?", Response: true},
			{Text: "Is the Amazon a rainforest?", Response: true},
			{Text: "Is Mount Kilimanjaro in Africa?", Response: true},
			{Text: "Is the Sahara Desert the largest desert in the world?", Response: true},
			{Text: "Is the capital of Japan Tokyo?", Response: true},
			{Text: "Is Everest in the Andes?", Response: false},
			{Text: "Is the Yangtze River in China?", Response: true},
			{Text: "Is the capital of Australia Sydney?", Response: false}, // Changed to false
			{Text: "Is the Dead Sea the lowest point on Earth?", Response: true},
			{Text: "Does the Amazon River run through Brazil?", Response: true},
			{Text: "Is the Atacama Desert in Chile?", Response: true},
			{Text: "Is the capital of Canada Toronto?", Response: false}, // Changed to false
			{Text: "Is Mount Fuji in Japan?", Response: true},
			{Text: "Does the Thames River pass through London?", Response: true},
			{Text: "Is the capital of Russia Moscow?", Response: true},
			{Text: "Is the Gobi Desert in Mongolia?", Response: true},
			{Text: "Is the capital of Egypt Cairo?", Response: true},
			{Text: "Does the Danube River pass through Vienna?", Response: true},
			{Text: "Is the capital of Argentina Buenos Aires?", Response: true},
			{Text: "Does the Seine River pass through Paris?", Response: true},
			{Text: "Is the capital of Italy Rome?", Response: true},
			{Text: "Is the Volga River in Russia?", Response: true},
			{Text: "Is the capital of Germany Berlin?", Response: true},
			{Text: "Is the Kalahari Desert in Africa?", Response: true},
			{Text: "Is the capital of Mexico Mexico City?", Response: true},
			{Text: "Is the Ganges River in India?", Response: true},
			{Text: "Is the capital of China Beijing?", Response: true},
			{Text: "Is the Sonoran Desert in the United States?", Response: true},
			{Text: "Is the capital of Brazil Rio de Janeiro?", Response: false}, // Changed to false
			{Text: "Is the Mississippi River in the United States?", Response: true},
			{Text: "Is the capital of India New Delhi?", Response: true},
			{Text: "Is the Mojave Desert in the United States?", Response: true},
			{Text: "Is the capital of Peru Lima?", Response: true},
			{Text: "Is the Ebro River in Spain?", Response: true},
			{Text: "Is the capital of Colombia Bogota?", Response: true},
			{Text: "Is the Namib Desert in Africa?", Response: true},
			{Text: "Is the capital of Chile Santiago?", Response: true},
			{Text: "Does the Nile River pass through Egypt?", Response: true},
			{Text: "Is the capital of Venezuela Toronto?", Response: false},
			{Text: "Is the Simpson Desert in Australia?", Response: true},
			{Text: "Is the capital of Uruguay Montevideo?", Response: true},
			{Text: "Is the Tigris River in Iraq?", Response: true},
			{Text: "Is the capital of Perú La Paz?", Response: false},
			{Text: "Is the Thar Desert in India?", Response: true},
			{Text: "Is the capital of Paraguay Asuncion?", Response: true},
			{Text: "Is the Mekong River in Asia?", Response: true},
			{Text: "Is the capital of Ecuador Mallorca?", Response: false},
			{Text: "Is the Karakum Desert in Turkmenistan?", Response: true},
			{Text: "Is the capital of China Tegucigalpa?", Response: false},
			{Text: "Is the Po River in Germany?", Response: false},
			{Text: "Is the capital of Guatemala Guatemala City?", Response: true},
			{Text: "Is the Rub al-Jali Desert in Saudi Arabia?", Response: true},
			{Text: "Is the capital of Nicaragua Managua?", Response: true},
			{Text: "Is the Loire River in France?", Response: true},
		
		},
	}
}

func (q Quizzer) NewQuizz() Quizz {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	return q.list[ra.Intn(len(q.list)-1)]
}

type Quizz struct {
	Text     string
	Response bool
}
