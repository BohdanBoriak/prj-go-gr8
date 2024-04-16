package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"prj-go/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsPerQuestion = 50
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі MATHCORE!")
	time.Sleep(1 * time.Second)

	users := getUsers()
	for _, user := range users {
		if user.Id >= id {
			id = user.Id + 1
		}
	}

	for {
		menu()

		point := ""
		fmt.Scan(&point)

		switch point {
		case "1":
			u := play()
			users := getUsers()
			users = append(users, u)
			sortAndSave(users)
		case "2":
			users := getUsers()
			for i, user := range users {
				fmt.Printf(
					"i: %v, id: %v, name: %s, time: %v\n",
					i, user.Id, user.Name, user.Time,
				)
			}
		case "3":
			return
		default:
			fmt.Println("Не коректний вибір 0_о")
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	start := time.Now()
	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if ansInt == res {
				myPoints += pointsPerQuestion
				fmt.Printf(
					"Чудово, ти зібрав(ла) %v!\nЗалишилось зібрати %v!\n",
					myPoints, totalPoints-myPoints,
				)
			} else {
				fmt.Println("Пощастить наступного разу Т_Т")
			}
		}
	}
	finish := time.Now()
	timeSpent := finish.Sub(start)
	fmt.Printf("Вітаємо, ти вправся за %v!\n Введіть своє ім'я: ", timeSpent)

	name := ""
	fmt.Scan(&name)

	// var user domain.User
	// user.Id = id
	// user.Name = name
	// user.Time = timeSpent

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755,
	)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}

func getUsers() []domain.User {
	var users []domain.User

	info, err := os.Stat("users.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("Error: %s", err)
			return nil
		}
		_, err := os.Create("users.json")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}
		return nil
	}

	if info.Size() != 0 {
		file, err := os.Open("users.json")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}
	}

	return users
}
