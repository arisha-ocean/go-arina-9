package main

import "testing"

// TestGenerateRandomElementsPositiveSize проверяет генерацию слайса заданного размера.
func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	size := 10
	data := generateRandomElements(size)

	if len(data) != size {
		t.Errorf("expected length %d, got %d", size, len(data))
	}

	for _, value := range data {
		if value < 1 || value > 1_000_000 {
			t.Errorf("value %d is out of range", value)
		}
	}
}

// TestGenerateRandomElementsZeroSize проверяет генерацию пустого слайса при размере 0.
func TestGenerateRandomElementsZeroSize(t *testing.T) {
	data := generateRandomElements(0)

	if len(data) != 0 {
		t.Errorf("expected empty slice, got length %d", len(data))
	}
}

// TestGenerateRandomElementsNegativeSize проверяет генерацию пустого слайса при отрицательном размере.
func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	data := generateRandomElements(-5)

	if len(data) != 0 {
		t.Errorf("expected empty slice, got length %d", len(data))
	}
}

// TestMaximumEmptySlice проверяет работу функции с пустым слайсом.
func TestMaximumEmptySlice(t *testing.T) {
	data := []int{}

	result := maximum(data)

	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

// TestMaximumSingleElement проверяет работу функции со слайсом из одного элемента.
func TestMaximumSingleElement(t *testing.T) {
	data := []int{42}

	result := maximum(data)

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

// TestMaximumMultipleElements проверяет поиск максимума среди нескольких элементов.
func TestMaximumMultipleElements(t *testing.T) {
	data := []int{5, 12, 3, 45, 8}

	result := maximum(data)

	if result != 45 {
		t.Errorf("expected 45, got %d", result)
	}
}

// TestMaximumNegativeNumbers проверяет поиск максимума среди отрицательных чисел.
func TestMaximumNegativeNumbers(t *testing.T) {
	data := []int{-10, -3, -25, -7}

	result := maximum(data)

	if result != -3 {
		t.Errorf("expected -3, got %d", result)
	}
}

// TestMaxChunksEmptySlice проверяет работу функции с пустым слайсом.
func TestMaxChunksEmptySlice(t *testing.T) {
	data := []int{}

	result := maxChunks(data)

	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

// TestMaxChunksSingleElement проверяет работу функции с одним элементом.
func TestMaxChunksSingleElement(t *testing.T) {
	data := []int{42}

	result := maxChunks(data)

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

// TestMaxChunksMultipleElements проверяет поиск максимума в большом слайсе.
func TestMaxChunksMultipleElements(t *testing.T) {
	data := []int{5, 12, 3, 45, 8, 100, 2, 17, 9, 30}

	result := maxChunks(data)

	if result != 100 {
		t.Errorf("expected 100, got %d", result)
	}
}

// TestMaxChunksNegativeNumbers проверяет поиск максимума среди отрицательных чисел.
func TestMaxChunksNegativeNumbers(t *testing.T) {
	data := []int{-10, -3, -25, -7, -15, -2, -30, -8}

	result := maxChunks(data)

	if result != -2 {
		t.Errorf("expected -2, got %d", result)
	}
}
