package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	stepLen   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// actions int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(actions int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	// ваш код здесь
	switch trainingType {
	case "Бег":
		distance := distance(actions)                               // вызовите здесь необходимую функцию
		speed := meanSpeed(actions, duration)                       // вызовите здесь необходимую функцию
		calories := RunningSpentCalories(actions, weight, duration) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case "Ходьба":
		distance := distance(actions)                                       // вызовите здесь необходимую функцию
		speed := meanSpeed(actions, duration)                               // вызовите здесь необходимую функцию
		calories := WalkingSpentCalories(actions, duration, weight, height) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case "Плавание":
		distance := distance(actions)                                              // вызовите здесь необходимую функцию
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)                // вызовите здесь необходимую функцию
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// actions int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(actions int) float64 {
	return float64(actions) * stepLen / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// actions int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func meanSpeed(actions int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance(actions) / duration
}

// RunningSpentCalories возвращает количество потраченных калорий при беге.
//
// Параметры:
//
// actions int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunningSpentCalories(actions int, weight, duration float64) float64 {
	// Константы для расчета калорий, расходуемых при беге.
	const (
		runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
		runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
	)

	// ваш код здесь
	meanSpeed := meanSpeed(actions, duration)
	spentCalories := runningCaloriesMeanSpeedMultiplier * meanSpeed * runningCaloriesMeanSpeedShift * weight /
		mInKm * duration * minInH
	return spentCalories
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// actions int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(actions int, duration, weight, height float64) float64 {
	// Константы для расчета калорий, расходуемых при ходьбе.
	const (
		walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
		walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
	)

	// ваш код здесь
	meanSpeed := meanSpeed(actions, duration) * kmhInMsec
	heightInMeters := height / cmInM
	spentCalories := (walkingCaloriesWeightMultiplier*weight + (math.Pow(meanSpeed, 2)/heightInMeters)*
		walkingSpeedHeightMultiplier*weight) * duration * minInH
	return spentCalories
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	// Константы для расчета калорий, расходуемых при плавании.
	const (
		swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых калорий при плавании относительно скорости.
		swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
	)

	// ваш код здесь
	swimmingMeanSpeed := swimmingMeanSpeed(lengthPool, countPool, duration)
	spentCalories := (swimmingMeanSpeed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier *
		weight * duration
	return spentCalories
}
