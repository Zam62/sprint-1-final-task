package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sprint-1-final-task/service"
	"time"
)

type reqBody struct {
	Expression string
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// Выражение
	var expr reqBody
	switch r.Method {
	case "POST":
		_, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request")
			// Логируем информацию о запросе
			log.Printf("Запрос: %s %s", r.Method, r.URL.Path)
			return
		}
		log.Printf("Запрос: %s %s", expr.Expression, r.URL.Path)

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Only POST requests are allowed",
		})

	}
}

func calculateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if r.Method == "POST" {
			var expr reqBody

			err := json.NewDecoder(r.Body).Decode(&expr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			}
			defer r.Body.Close()

			// Логируем информацию о запросе
			log.Printf("Запрос: %s %s %s", r.Method, r.URL.Path, expr.Expression)
			// Считаем выражение
			result, err := service.Calc(expr.Expression)
			// log.Printf("Запрос: %s", result)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "Expression is not valid",
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"result": int(result),
			})
		}

		// Передаём управление следующему обработчику
		next.ServeHTTP(w, r)

		// Вычисляем время выполнения запроса
		duration := time.Since(start)
		log.Printf("Время выполнения запроса: %s", duration)
	})
}

func main() {
	//*
	mux := http.NewServeMux()

	// Создаём обработчик для маршрута "/"
	calculate := http.HandlerFunc(RequestHandler)

	// Применяем middleware к обработчику "/"
	mux.Handle("/api/v1/calculate", calculateMiddleware(calculate))

	// Запускаем сервер на порту 3030
	fmt.Printf("Starting server for web-service...\n")
	if err := http.ListenAndServe(":3030", mux); err != nil {
		log.Fatal(err)
	}

}
