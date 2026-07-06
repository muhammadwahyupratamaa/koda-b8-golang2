package main

import "fmt"

// soal 1
type We struct{
	are Are
}
type Are struct{
	the The	
}
type The struct{
	best string 
}

// Soal 2

type Hello struct{
	world string
}



type Obj struct {
	str [][][]Man 
}
type Man struct {
	man [] Tech
}
type Academy struct {
	academy string
}
type Tech struct {
	tech Academy
}





// Soal 5
type Number struct {
	first [] int
	second [] int
}




func main() {
// Soal 1 
we := We{
	are: Are{
		the: The{
			best: "KODA",
		},
	} ,
}

// Soal 2

hello := Hello{
	world : "Hello World",
}

// Soal 3
obj := Obj{
	str: [][][]Man{
		{},{},{},{
			{}, {
				{},{},{
					man:[]Tech{
						{
							tech: Academy{
								academy: "Tech Academy",
							},
						},
					},
				},
			},
		},
	},
}


// Soal 4




// Soal 5
num := Number{
	first: []int{
		10,15,30,
	},
	second: []int{
		10,17,17,

	},
}

fmt.Println(we.are.the.best)
fmt.Println(hello.world)
fmt.Println(obj.str[3][1][2].man[0].tech.academy)
fmt.Println(num.first[1] + num.second[2])
}