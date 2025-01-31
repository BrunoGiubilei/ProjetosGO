package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brunogiubilei/go-rest-api/greetings"
	"github.com/gorilla/mux"

	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func GetAPI(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	greeting := greetings.GetGreeting(name)
	response := map[string]string{"message": greeting}
	json.NewEncoder(w).Encode(response)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decodificar o corpo da requisição JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inserir dados na tabela
	insertUserSQL := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err = db.Exec(insertUserSQL, user.Name, user.Age)
	if err != nil {
		log.Fatal(err)
	}

	// Resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully"})

	db.Close()
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Buscar todos os registros da tabela users
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// Verificar erros após a iteração
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Responder com a lista de usuários
	json.NewEncoder(w).Encode(users)

	db.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Obter o ID do usuário a partir dos parâmetros da URL
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Deletar o usuário da tabela
	deleteUserSQL := `DELETE FROM users WHERE id = ?`
	_, err = db.Exec(deleteUserSQL, userID)
	if err != nil {
		log.Fatal(err)
	}

	// Resposta de sucesso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})

	db.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Obter o ID do usuário a partir dos parâmetros da URL
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User
	// Decodificar o corpo da requisição JSON
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Atualizar dados na tabela
	updateUserSQL := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	_, err = db.Exec(updateUserSQL, user.Name, user.Age, userID)
	if err != nil {
		log.Fatal(err)
	}

	// Resposta de sucesso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}
