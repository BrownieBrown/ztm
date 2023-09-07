package challenges

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	hashMap := make(map[*ListNode]bool)
	current := head

	for current != nil {
		if _, found := hashMap[current]; found {
			return current // Return the node where the cycle starts
		}

		hashMap[current] = true
		current = current.Next
	}

	return nil // No cycle detected
}

//Now, for Floyd's Cycle Detection Algorithm (also known as "Tortoise and the Hare" approach):
//
//Theory:
//Imagine a circular racetrack. If you have two runners, one running twice as fast as the other (the hare runs twice as fast as the tortoise), they will eventually meet if they started from the same point, right? The same concept applies to our linked list with a cycle.
//
//Steps:
//
//Start two pointers, slow (tortoise) and fast (hare), at the head of the linked list.
//Move slow one step at a time and fast two steps at a time.
//If there's a cycle, the fast pointer will eventually catch up to the slow pointer, and they'll meet at some node in the cycle.
//Once they meet, reset one of the pointers to the head of the list and then move both pointers one step at a time.
//The point where they next meet is the start of the cycle.

func detectCycleFloyd(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// Step 1 & 2: Move slow one step and fast two steps
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// Step 3: Check if they meet
		if slow == fast {
			// Step 4: Reset one of the pointers
			fast = head

			// Step 5: Find the start of the cycle
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			return slow // Return the node where the cycle starts
		}
	}

	return nil // No cycle detected
}
