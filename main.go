package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/SaidovZohid/websocket-task/database"
	"github.com/SaidovZohid/websocket-task/websocket"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := database.Make()
	if err != nil {
		log.Fatal("Database Make:", err)
	}
	cm := &websocket.ClientManager{
		Clients: make(map[string]*websocket.Client),
		Db:      conn,
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleConnections(cm, conn, w, r)
	})

	go func() {
		rand.New(rand.NewSource(time.Now().Unix()))
		ln := len(quotes)
		for {
			quote := quotes[rand.Intn(ln)]
			message := "Quote: " + quote
			cm.Broadcast(message)
			time.Sleep(5 * time.Second)
		}
	}()

	log.Println("WebSocket server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today. – Franklin D. Roosevelt",
	"Success is not final, failure is not fatal: It is the courage to continue that counts. – Winston Churchill",
	"Do not watch the clock; do what it does. Keep going. – Sam Levenson",
	"You miss 100% of the shots you don't take. – Wayne Gretzky",
	"Believe you can and you're halfway there. – Theodore Roosevelt",
	"Act as if what you do makes a difference. It does. – William James",
	"The future belongs to those who believe in the beauty of their dreams. – Eleanor Roosevelt",
	"It is never too late to be what you might have been. – George Eliot",
	"Everything you’ve ever wanted is on the other side of fear. – George Addair",
	"Success is not how high you have climbed, but how you make a positive difference to the world. – Roy T. Bennett",
	"Hardships often prepare ordinary people for an extraordinary destiny. – C.S. Lewis",
	"Your time is limited, don’t waste it living someone else’s life. – Steve Jobs",
	"Don’t be afraid to give up the good to go for the great. – John D. Rockefeller",
	"Failure is another stepping stone to greatness. – Oprah Winfrey",
	"Success usually comes to those who are too busy to be looking for it. – Henry David Thoreau",
	"Don’t count the days, make the days count. – Muhammad Ali",
	"Opportunities don't happen, you create them. – Chris Grosser",
	"The road to success and the road to failure are almost exactly the same. – Colin R. Davis",
	"Don’t let yesterday take up too much of today. – Will Rogers",
	"Start where you are. Use what you have. Do what you can. – Arthur Ashe",
	"Success is walking from failure to failure with no loss of enthusiasm. – Winston Churchill",
	"I never dreamed about success, I worked for it. – Estée Lauder",
	"Go as far as you can see; when you get there, you’ll be able to see further. – Thomas Carlyle",
	"The harder I work, the luckier I get. – Gary Player",
	"Don’t watch the clock; do what it does. Keep going. – Sam Levenson",
	"Do something today that your future self will thank you for. – Sean Patrick Flanery",
	"It always seems impossible until it’s done. – Nelson Mandela",
	"Success is not the key to happiness. Happiness is the key to success. – Albert Schweitzer",
	"Success doesn’t come to you, you go to it. – Marva Collins",
	"Small daily improvements over time lead to stunning results. – Robin Sharma",
	"Success is getting what you want. Happiness is wanting what you get. – Dale Carnegie",
	"Success is liking yourself, liking what you do, and liking how you do it. – Maya Angelou",
	"Don’t be pushed by your problems. Be led by your dreams. – Ralph Waldo Emerson",
	"Work hard in silence, let your success be your noise. – Frank Ocean",
	"Success is the sum of small efforts, repeated day in and day out. – Robert Collier",
	"Do what you can, with what you have, where you are. – Theodore Roosevelt",
	"Success is not in what you have, but who you are. – Bo Bennett",
	"Dream big and dare to fail. – Norman Vaughan",
	"Your success and happiness lie in you. – Helen Keller",
	"You must expect great things of yourself before you can do them. – Michael Jordan",
	"Success is a state of mind. If you want success, start thinking of yourself as a success. – Joyce Brothers",
	"The way to get started is to quit talking and begin doing. – Walt Disney",
	"Success is the result of preparation, hard work, and learning from failure. – Colin Powell",
	"Failure is the condiment that gives success its flavor. – Truman Capote",
	"You are never too old to set another goal or to dream a new dream. – C.S. Lewis",
	"Success doesn’t just find you. You have to go out and get it. – Anonymous",
	"Believe in yourself and all that you are. Know that there is something inside you that is greater than any obstacle. – Christian D. Larson",
	"Start where you are. Use what you have. Do what you can. – Arthur Ashe",
	"Success is not in what you have, but who you are. – Bo Bennett",
	"Do what you can with all you have, wherever you are. – Theodore Roosevelt",
	"Success is not the absence of failure; it's the persistence through failure. – Aisha Tyler",
	"Action is the foundational key to all success. – Pablo Picasso",
	"Success is how high you bounce when you hit bottom. – George S. Patton",
	"Success seems to be connected with action. Successful people keep moving. – Conrad Hilton",
	"The only place where success comes before work is in the dictionary. – Vidal Sassoon",
	"Success is the progressive realization of a worthy goal or ideal. – Earl Nightingale",
	"It is our choices that show what we truly are, far more than our abilities. – J.K. Rowling",
	"Only put off until tomorrow what you are willing to die having left undone. – Pablo Picasso",
	"The starting point of all achievement is desire. – Napoleon Hill",
	"Success is liking yourself, liking what you do, and liking how you do it. – Maya Angelou",
	"Don’t let what you cannot do interfere with what you can do. – John Wooden",
	"Success is going from failure to failure without losing your enthusiasm. – Winston Churchill",
	"Don’t be afraid to give up the good to go for the great. – John D. Rockefeller",
	"Successful people do what unsuccessful people are not willing to do. – Jim Rohn",
	"Success is not the key to happiness. Happiness is the key to success. – Albert Schweitzer",
	"The secret to success is to know something nobody else knows. – Aristotle Onassis",
	"Success is the sum of small efforts, repeated day in and day out. – Robert Collier",
	"The secret of getting ahead is getting started. – Mark Twain",
	"Success is getting what you want. Happiness is wanting what you get. – Dale Carnegie",
	"You don’t have to be great to start, but you have to start to be great. – Zig Ziglar",
	"I find that the harder I work, the more luck I seem to have. – Thomas Jefferson",
	"Don't wait for opportunity. Create it. – George Bernard Shaw",
	"The harder the struggle, the more glorious the triumph. – Swami Sivananda",
	"Success is achieved by developing our strengths, not by eliminating our weaknesses. – Marilyn vos Savant",
	"Dream it. Believe it. Build it. – Anonymous",
	"Success is sweet and sweeter if long delayed and gotten through many struggles and defeats. – Amos Bronson Alcott",
	"Success is not just about making money. It's about making a difference. – Anonymous",
	"The road to success and the road to failure are almost exactly the same. – Colin R. Davis",
	"Success is a journey, not a destination. The doing is often more important than the outcome. – Arthur Ashe",
	"Success is not measured by what you accomplish, but by the opposition you have encountered. – Booker T. Washington",
	"Success is not how high you have climbed, but how you make a positive difference to the world. – Roy T. Bennett",
	"Success is peace of mind, which is a direct result of self-satisfaction in knowing you did your best. – John Wooden",
	"Success is falling nine times and getting up ten. – Jon Bon Jovi",
	"Success is to be measured not so much by the position that one has reached in life as by the obstacles which he has overcome. – Booker T. Washington",
	"Success is the result of perfection, hard work, learning from failure, loyalty, and persistence. – Colin Powell",
	"Success isn’t about being the best. It’s about always getting better. – Behance",
	"Success is not the key to happiness. Happiness is the key to success. If you love what you are doing, you will be successful. – Albert Schweitzer",
	"The biggest risk is not taking any risk. – Mark Zuckerberg",
	"Success is not in what you have, but who you are. – Bo Bennett",
	"It’s not whether you get knocked down, it’s whether you get up. – Vince Lombardi",
	"Success consists of going from failure to failure without loss of enthusiasm. – Winston Churchill",
	"Success is the progressive realization of a worthy ideal. – Earl Nightingale",
	"The best way to predict your future is to create it. – Peter Drucker",
}
