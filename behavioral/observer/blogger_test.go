package observer_test

import (
	"patterns/behavioral/observer"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Subscriber(t *testing.T) {
	blogs := observer.NewBlogs()

	alice := observer.NewSubscriber("Alice")
	bob := observer.NewSubscriber("Bob")

	blogs.Subscribe(alice)
	blogs.Subscribe(bob)

	blogs.AddNews("New post about Go patterns!")

	require.Len(t, alice.GetNews(), 1, "Alice should have received one news item")
	require.Len(t, bob.GetNews(), 1, "Bob should have received one news item")
}
