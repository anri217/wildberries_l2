package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в
	собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.


	ПРИМЕНИМОСТЬ

	Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Стратегия позволяет варьировать поведение объекта во время выполнения программы, подставляя в него различные
	объекты-поведения (например, отличающиеся балансом скорости и потребления ресурсов).

	Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
	Стратегия позволяет вынести отличающееся поведение в отдельную иерархию классов, а затем свести первоначальные классы к одному,
	сделав поведение этого класса настраиваемым.


	ПРЕИМУЩЕСТВА

	Горячая замена алгоритмов на лету.
	Изолирует код и данные алгоритмов от остальных классов.
	Уход от наследования к делегированию.
	Реализует принцип открытости/закрытости.


	НЕДОСТАТКИ

	Усложняет программу за счёт дополнительных классов.
	Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
