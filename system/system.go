package system

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	Name string
}

type System struct {
	Conf           Config
	Counter        int
	CounterCurrent int
	CounterChan    chan int
	CounterMut     sync.Mutex
}

func (s *System) Initialize() {
	s.CounterChan = make(chan int)
	s.setConfig()
}

func (s *System) RunCounter() {
	for {
		count := <-s.CounterChan
		s.CounterMut.Lock()
		if count == 1 {
			s.Counter++
			s.CounterCurrent++
		} else {
			s.CounterCurrent--
		}
		s.CounterMut.Unlock()
	}
}

func (s *System) ReadCurrent() int {
	s.CounterMut.Lock()
	defer s.CounterMut.Unlock()
	return s.CounterCurrent
}

func (s *System) setConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loadimg .env: %s", err)
	}

	s.Conf.Name = os.Getenv("PROJECT_NAME")
	fmt.Println(s.Conf.Name)
}
