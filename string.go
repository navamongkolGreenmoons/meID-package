package main

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (g *General) HashPassword(pwd string, saltKey string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+saltKey), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (g *General) ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	return err == nil
}

func (g *General) RandSaltKey() string {
	charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "~!@#$%^&*()_"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 20)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (g *General) InStringList(value string, list []string) bool {
	for _, str := range list {
		if str == value {
			return true
		}
	}
	return false
}
