# nomad_Coin
nomad_Coin study

#   노마드코인 스터디 Rough

### 상수 및 변수 ** 
상수선언 const name string = "nico"
변수선언 var name string = "nico"
변수선언 var name:= "nico"


### 함수

* func 함수명(입력 파라미터) (전달 파라미터)
 
* 형식 1
func func_01(name string) (int, string) {
    return len(name), strings.ToUpper(name)
}
 
func func_02(name string) int {
    return len(name)
}

func Fn_Afunc_03(name string) (length int, uppercase string) {
    length = len(name)
    uppercase = strings.ToUpper(name)
    return
}

### 반복문

func for_01(numbers ...int) {
	for number := range numbers {
		fmt.Println(number)
	}
}
 
func for_02(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}
 
func for_03(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}


### IF

func if_01() {
	age := 23
	if age < 23 {
		fmt.Println(age)
	}
}
  
func if_02() {
	if age2 := 23; age2 < 23 {
		fmt.Println("age21 ==========")
		fmt.Println(age2)
	}
	fmt.Println("age22")
    
    // if 문 밖에서 변수 호출 불가능
	fmt.Println(age2)
}


### SWITCH

func swith_01() {
	age := 30
	switch age {

	case 10:
		println(10)
	case 20:
		println(20)
	case 30:
		println(30)
	}
}

func swith_02() {
	age := 30
	switch {
	case age == 10:
		println(10)
	case age == 20:
		println(20)
	case age == 30:
		println(30)
	}
}

### ARRAY

func array_01() {
	names := [5]string{"aaa", "bbb", "ccc"}
	fmt.Println(names)
	// [aaa bbb ccc  ]

	names[3] = "ddd"
	names[4] = "eee"
	// names[5] = "fff" : 오류 발생
	
}

func array_02() {
	names := []string{"aaa", "bbb", "ccc"}
	fmt.Println(names)
	// [aaa bbb ccc  ]
	
	// names[3] = "ddd"     : 오류 발생
	// append(names, "ddd") : 오류 발생

	names = append(names, "ddd")
	fmt.Println(names)
	// [aaa bbb ccc ddd]
}


### POINTER

func pointer_01() {
	a := 2
	b := 3

	fmt.Println(a, b)
	// 2 3

	// 메모리 포인투 주소 출력
	fmt.Println(&a, &b)
	// 0xc000014098 0xc0000140b0

	fmt.Println(a + b)
}

func pointer_02() {
	a := 2
	b := &a

	fmt.Println(a, b)
	// 2 0xc0000140f0

	// 메모리 포인투 주소 출력
	fmt.Println(&a, &b)
	// 0xc000014098 0xc0000140b0
}

func pointer_03() {
	a := 2
	b := &a

	fmt.Println(a, &a, b, &b)
	// 2 0xc0000140f8 0xc0000140f8 0xc000006038

	*b = 20

	fmt.Println(a, &a, b, &b)
	// 20 0xc0000140f8 0xc0000140f8 0xc000006038
	
    // 메모리 값을 볼때는 *
	fmt.Println(a, &a, *b, &b)
	// 20 0xc0000140f8 20 0xc000006038
}


### MAP

func map_01() {
	nico := map[string]string{"name": "nico", "age": "30"}
	fmt.Println(nico)
}


### STRUCTS

func struncts_01() {
	favFood := []string{"김치", "라면"}
	nico := person{"nico", 22, favFood}
	fmt.Println(nico)
	// {nico 22 [김치 라면]}
}

func struncts_02() {
	favFood := []string{"김치", "라면"}
	nico := person{name: "nico", age: 22, favFood: favFood}
	fmt.Println(nico)
	// {nico 22 [김치 라면]}
	fmt.Println(nico.name)
	// nico
}
