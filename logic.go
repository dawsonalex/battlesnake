package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |______  /(____  /__|  |__| |____/\___  >____  >___|  (____  /__|_ \\___  >
//         \/      \/                     \/     \/     \/     \/     \/    \/
// 
// This file can be a nice home for your Battlesnake game logic.
// To get started we've included code to help your Battlesnake avoid moving backwards.
// For more info see docs.battlesnake.com

import (
	"log"
	"math/rand"
)

// info() is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance and author permissions
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose a color
		Head:       "default", // TODO: Choose a head
		Tail:       "default", // TODO: Choose a tail
	}
}

// start() is called when your Battlesnake begins a new game
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// end() is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// move() is called on every turn of the game and returns your next move
// Valid moves are "up", "down", "left", or "right".
func move(state GameState) BattlesnakeMoveResponse {
    
	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// Step 0 - Never move backwards
	myHead := state.You.Body[0]     // Coordinates of your head
	myNeck := state.You.Body[1]     // Coordinates of your "neck"
    
	if myNeck.X < myHead.X {        // Neck is left of head, can't move left 
		isMoveSafe["left"] = false
        
	} else if myNeck.X > myHead.X { // Neck is right of head, can't move right
		isMoveSafe["right"] = false
        
	} else if myNeck.Y < myHead.Y { // Neck is below head, can't move down
		isMoveSafe["down"] = false
        
	} else if myNeck.Y > myHead.Y { // Neck is above head, can't move up
		isMoveSafe["up"] = false
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	// boardWidth := state.Board.Width
	// boardHeight := state.Board.Height

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	// mybody := state.You.Body

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
    // opponents := state.Board.Snakes

	// TODO: Step 4 - Move towards food to regain health and survive
	// food := state.Board.Food

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Printf("%s MOVE %d: No safe moves detected! Moving down\n", state.Game.ID, state.Turn)
        return BattlesnakeMoveResponse{Move: "down"}
    }

    // Choose a random move from the safe ones
    // TODO: Step 5 - Build a better strategy for selecting our next move
    nextMove := safeMoves[rand.Intn(len(safeMoves))]
    log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
    
    return BattlesnakeMoveResponse{Move: nextMove}
}
