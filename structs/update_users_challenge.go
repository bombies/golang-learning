package structs

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	return User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: func() int {
				if membershipType == "premium" {
					return 1000
				} else {
					return 100
				}
			}(),
		},
	}
}

func TestNewUser() {
	user1 := newUser("Ajani", "premium")
	user2 := newUser("Sally", "basic")
	println(user1.Name, user1.Membership.Type, user1.Membership.MessageCharLimit)
	println(user2.Name, user2.Membership.Type, user2.Membership.MessageCharLimit)
}
