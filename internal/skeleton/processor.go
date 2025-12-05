package skeleton

import (
	"path/filepath"
)

// Process визначає стратегію обробки на основі розширення файлу
func Process(filename string, content []byte) ([]byte, error) {
	ext := filepath.Ext(filename)

	switch ext {
	case ".go":
		return reduceGo(content)
	// Тут згодом додаси .ts, .py, .java
	default:
		// Якщо парсера немає — повертаємо оригінал (або можна додати generic regex reducer)
		return content, nil
	}
}
