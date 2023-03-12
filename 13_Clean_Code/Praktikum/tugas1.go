package main

type user struct {
	id       int
	username int // salah dalam memberi type data untuk username seharusnya string
	password int // salah dalam memberi type data untuk password seharusnya string
}

type userservice struct { // penamaan tidak menggunakan Camel Case seharusnya userService
	t []user //penamaan t kurang bisa dipahami
}

func (u userservice) getallusers() []user { // penamaan tidak menggunakan Camel Case seharusnya getAllUsers
	// dan juga penamaan parameter kurang bisa dipahami
	return u.t
}

func (u userservice) getuserbyid(id int) user { // penamaan tidak menggunakan Camel Case seharusnya getUserById
	// dan juga penamaan parameter kurang bisa dipahami

	for _, r := range u.t { //penamaan r juga kurang bisa dipahami seharusnya bisa menggunakan value atau lainya
		if id == r.id {
			return r
		}
	}

	return user{}
}