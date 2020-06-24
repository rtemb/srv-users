package permissions

import srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"

var Permissions = map[srvUsers.Role][]string{
	srvUsers.Role_USER: {
		"USER_READ",
	},
	srvUsers.Role_USER_ADMIN: {
		"USER_CREATE",
		"USER_READ",
	},
}
