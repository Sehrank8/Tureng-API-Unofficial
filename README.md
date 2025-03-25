# Tureng Translation API English to Turkish.

A simple REST API that scrapes **Tureng** to fetch English-to-Turkish translations.  
Built with **Go (Golang)** and `colly` for web scraping.

---

## Features
* Fetch Turkish-to-English translations from **Tureng**  
* Provides **word type (isim, fiil, sıfat, etc.)**  
* Returns **category & source** (e.g., technical, general, economics)  
* JSON response with **word, count, and translations**  

---

##  Installation
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/your-username/tureng-api.git
cd tureng-api
```

### **2️⃣ Install Dependencies**
```sh
go mod tidy
```
### **2️⃣ Install Dependencies**
```sh
go run main.go
```
The server will start on http://localhost:8080.


## Usage
### **GET /translate/{word}**
Fetch translations for a given English word.

**Example Request:**
```sh
curl http://localhost:8080/translate/kitap
```

**Example Response:**
```json
{
  "word": "book",
  "count": 2,
  "translations": [
    {
      "category": "Genel",
      "source": "book",
      "word_type": "isim",
      "translation": "kitap"
    },
    {
      "category": "Teknik",
      "source": "book",
      "word_type": "isim",
      "translation": "kurallar"
    }
  ]
}
```



