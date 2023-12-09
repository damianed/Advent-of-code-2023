package main

import (
  "fmt"
  "strings"
  "strconv"

  "adventofcode/helper"
)

const (
  FIVE_OF_KIND int = 7
  FOUR_OF_KIND     = 6
  FULL_HOUSE       = 5
  THREE_OF_KIND    = 4
  TWO_PAIR         = 3
  ONE_PAIR         = 2
  HIGH_CARD        = 1
)

type Hand struct {
  cards string
  bid int
  handType int
}

type Node struct {
  next *Node
  prev *Node
  value Hand
}

type Ranks struct {
  head *Node
  length int
}

type Pair struct {
  letter rune
  count int
}

func main() {
  lines := helper.ReadFile("input")

  hands := []Hand{}
  ranks := Ranks{}
  for _, line := range lines {
    data := strings.Fields(line)
    bid, _ := strconv.Atoi(data[1])
    cards := data[0]
    handType := getType(cards)
    hand := Hand{cards, bid, handType}
    hands = append(hands, hand)

    ranks = updateRanks(ranks, hand)
  }

  node := ranks.head
  i := ranks.length
  sum := 0
  for node != nil {
    sum += node.value.bid * i
    node = node.next
    i--
  }
  fmt.Println(sum)
}

func updateRanks(ranks Ranks, hand Hand) Ranks {
  cardValues := map[string]int {
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "J": 1,
    "Q": 12,
    "K": 13,
    "A": 14,
  }

  if (ranks.head == nil) {
    head := Node{nil, nil, hand}
    ranks.head = &head
    ranks.length++
    return ranks
  }

  node := ranks.head
  for node != nil {
    if node.value.handType < hand.handType {
      newN := Node{node, node.prev, hand}
      if (node.prev != nil) {
        node.prev.next = &newN
      }
      node.prev = &newN

      if (node == ranks.head) {
        ranks.head = &newN
      }
      ranks.length++

      break
    }

    if node.value.handType == hand.handType {
      nFirst := node.value.cards[0:1]
      hFirst := hand.cards[0:1]

      i := 0
      for nFirst == hFirst {
        i++
        nFirst = node.value.cards[i:i+1]
        hFirst = hand.cards[i:i+1]
      }

      if cardValues[nFirst] < cardValues[hFirst] {
        newN := Node{node, node.prev, hand}
        if (node.prev != nil) {
          node.prev.next = &newN
        }

        node.prev = &newN

        if (node == ranks.head) {
          ranks.head = &newN
        }
        ranks.length++

        break
      } else {
        if (node.next != nil && node.next.value.handType == node.value.handType) {
          node = node.next
          continue
        }

        newN := Node{node.next, node, hand}
        if (node.next != nil) {
          node.next.prev = &newN
        }
        node.next = &newN
        ranks.length++
        break
      }
    }

    if (node.next == nil) {
        newN := Node{node.next, node, hand}
        node.next = &newN
        ranks.length++
        break
    }
    node = node.next
  }

  return ranks
}

func getType(hand string) int {
  seen := map[rune]int{}
  jokers := 0
  topCount := Pair{' ', 0}
  for _, c := range hand {
    if string(c) == "J" {
      jokers++
      continue
    }

    if _, ok := seen[c]; !ok {
      seen[c] = 0
    }

    seen[c]++
    if seen[c] > topCount.count {
      topCount.letter = c
      topCount.count = seen[c]
    }
  }

  seen[topCount.letter] += jokers

  if len(seen) == 1 {
    return FIVE_OF_KIND
  }

  numbers := map[int]int{}
  for _, count := range seen {
    _, ok := numbers[count]
    if !ok {
      numbers[count] = 0
    }

    numbers[count]++
  }

  _, has4 := numbers[4]
  if has4 {
    return FOUR_OF_KIND
  }

  _, has3 := numbers[3]
  has2Count, has2 := numbers[2]

  if has3 && has2 {
    return FULL_HOUSE
  }

  if has3 {
    return THREE_OF_KIND
  }

  if has2 && has2Count == 2 {
    return TWO_PAIR
  }

  if has2 {
    return ONE_PAIR
  }

  return HIGH_CARD
}
