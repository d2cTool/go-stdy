package lesson1

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Deduplicate[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	write := 0
	for i := 0; i < len(s); i++ {
		if _, ok := seen[s[i]]; ok {
			continue
		}

		seen[s[i]] = struct{}{}
		s[write] = s[i]
		write++
	}
	return s[:write]
}

func Chunk[T any](s []T, size int) [][]T {
	res := make([][]T, 0, (len(s)+size-1)/size)
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		res = append(res, s[i:end])
	}
	return res
}

// func MaxSumWindow(s []int, k int) (sum int, start int) {

// }

/*

// Разбить слайс на части фиксированного размера. Последний чанк может быть меньше.
func Chunk[T any](s []T, size int) [][]T

// Chunk([]int{1,2,3,4,5}, 2) → [[1,2], [3,4], [5]]

// Скользящее окно (sliding window) Найти подпоследовательность заданной длины с максимальной суммой.
func MaxSumWindow(s []int, k int) (sum int, start int)

// Merge двух отсортированных слайсов. Объединить два отсортированных слайса в один отсортированный за O(n).
func MergeSorted(a, b []int) []int

// Сдвинуть элементы на k позиций вправо (или влево) без дополнительной памяти.
func Rotate[T any](s []T, k int)

// Rotate([]int{1,2,3,4,5}, 2) → [4,5,1,2,3]

// Проверить, является ли один слайс подпоследовательностью другого (порядок сохраняется).
func IsSubsequence(sub, full []int) bool

// IsSubsequence([]int{1,3,5}, []int{1,2,3,4,5}) → true

// Flatten вложенного слайса Превратить [][]T в []T.
func Flatten[T any](s [][]T) []T

// Сгруппировать элементы по значению функции-ключа.
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T

// Написать код, где два слайса делят один backing array, и показать, как append в одном влияет на другой.
// Продемонстрировать и объяснить поведение
//a := []int{1, 2, 3, 4, 5}
//b := a[2:4]
//b = append(b, 99)
// Что теперь в a? А если cap(a) был бы больше?

// Реализовать сортировку по нескольким полям (например, сначала по имени, потом по возрасту).
type Person struct {
	Name string
	Age  int
}

func SortByMultiple(people []Person)

// Разделить слайс так, чтобы все элементы, удовлетворяющие предикату, шли первыми.
func Partition[T any](s []T, pred func(T) bool) int

// Возвращает индекс первого элемента, не удовлетворяющего предикату
*/
