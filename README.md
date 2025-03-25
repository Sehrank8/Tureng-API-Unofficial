# [Tureng](https://tureng.com/tr/turkce-ingilizce) Translation API English to Turkish.
#### **Inspired by [tureng-api](https://github.com/gokhanamal/tureng-api), (Because I couldn't get it to work for some reason).**

A simple REST API that scrapes **Tureng** to fetch English-to-Turkish translations.  

Built with [**Go (Golang)**](https://go.dev/) and [**colly**](https://github.com/gocolly/colly) for web scraping.

There is also a branch with swagger support if you want to try it.

---

## Features
* Fetch English-to-Turkish translations from [**Tureng**](https://tureng.com/tr/turkce-ingilizce)  
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
### **3️⃣ Run the API**
```sh
go run main.go
```
The server will start on http://localhost:8080.

---

## Usage
### **GET /translate/{word}**
Fetch translations for a given English word.

**Example Request:**
```sh
curl http://localhost:8080/translate?word=book
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
---


