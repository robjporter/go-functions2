package main

import (
	"fmt"

	"../security/rbac"
)

const (
	SHOW     = 0
	CREATE   = 1
	UPDATE   = 2
	DELETE   = 3
	IMPORT   = 4
	DOWNLOAD = 5
	ADMIN    = 6
)

const (
	USER = 0
	BLOG = 1
)

const (
	ADMINROLE = 0
	USERROLE  = 1
)

func main() {
	rbac.AddUser("roporter", "roporter@cisco.com")
	rbac.AddUser("roporter2", "roporter2@cisco.com")

	rbac.AddRole("admin")
	rbac.AddRole("user")

	rbac.AddResource("user")
	rbac.AddResource("blog")

	rbac.AddOperation("show")
	rbac.AddOperation("create")
	rbac.AddOperation("update")
	rbac.AddOperation("delete")
	rbac.AddOperation("import")
	rbac.AddOperation("download")
	rbac.AddOperation("admin")

	fmt.Println("USER roporter ID:        >", rbac.GetUserIDByUsername("roporter"))
	fmt.Println("USER roporter2  ID:      >", rbac.GetUserIDByUsername("roporter2"))
	fmt.Println("USER roporter ID:        >", rbac.GetUserIDByEmail("roporter@cisco.com"))
	fmt.Println("USER roporter2  ID:      >", rbac.GetUserIDByEmail("roporter2@cisco.com"))

	fmt.Println("")

	fmt.Println("USER roporter POS:       >", rbac.GetUserPositionByUsername("roporter"))
	fmt.Println("USER roporter2  POS:     >", rbac.GetUserPositionByUsername("roporter2"))
	fmt.Println("USER roporter POS:       >", rbac.GetUserPositionByEmail("roporter@cisco.com"))
	fmt.Println("USER roporter2  POS:     >", rbac.GetUserPositionByEmail("roporter2@cisco.com"))

	fmt.Println("")

	fmt.Println("ROLE ADMIN ID:           >", rbac.GetRoleID("admin"))
	fmt.Println("ROLE USER ID:            >", rbac.GetRoleID("user"))

	fmt.Println("")

	fmt.Println("ROLE ADMIN POS:          >", rbac.GetRolePosition("admin"))
	fmt.Println("ROLE USER POS:           >", rbac.GetRolePosition("user"))

	fmt.Println("")

	fmt.Println("RESOURCE USER ID:        >", rbac.GetResourceID("user"))
	fmt.Println("RESOURCE BLOG ID:        >", rbac.GetResourceID("blog"))

	fmt.Println("")

	fmt.Println("RESOURCE USER POS:       >", rbac.GetResourcePosition("user"))
	fmt.Println("RESOURCE BLOG POS:       >", rbac.GetResourcePosition("blog"))

	fmt.Println("")

	fmt.Println("OPERATION SHOW ID:       >", rbac.GetOperationID("show"))
	fmt.Println("OPERATION CREATE ID:     >", rbac.GetOperationID("create"))
	fmt.Println("OPERATION UPDATE ID:     >", rbac.GetOperationID("update"))
	fmt.Println("OPERATION DELETE ID:     >", rbac.GetOperationID("delete"))
	fmt.Println("OPERATION IMPORT ID:     >", rbac.GetOperationID("import"))
	fmt.Println("OPERATION DOWNLOAD ID:   >", rbac.GetOperationID("download"))
	fmt.Println("OPERATION ADMIN ID:      >", rbac.GetOperationID("admin"))

	fmt.Println("")

	fmt.Println("OPERATION SHOW POS:      >", rbac.GetOperationPosition("show"))
	fmt.Println("OPERATION CREATE POS:    >", rbac.GetOperationPosition("create"))
	fmt.Println("OPERATION UPDATE POS:    >", rbac.GetOperationPosition("update"))
	fmt.Println("OPERATION DELETE POS:    >", rbac.GetOperationPosition("delete"))
	fmt.Println("OPERATION IMPORT POS:    >", rbac.GetOperationPosition("import"))
	fmt.Println("OPERATION DOWNLOAD POS:  >", rbac.GetOperationPosition("download"))
	fmt.Println("OPERATION ADMIN POS:     >", rbac.GetOperationPosition("admin"))

	fmt.Println("")

	rbac.AddResourceOperation(rbac.GetResourcePosition("user"), SHOW, CREATE, UPDATE, DELETE, IMPORT)
	//rbac.AddResourceOperation(rbac.GetResourcePosition("blog"), SHOW, CREATE, UPDATE, DELETE)
	rbac.AddResourceOperationByName("blog", "SHOW", "create", "update", "DELETE")
	rbac.AddResourceOperationByName("user", "download", "admin")

	rbac.AddRoleResource(ADMINROLE, USER, SHOW, CREATE, UPDATE, DELETE, IMPORT, DOWNLOAD, ADMIN)
	rbac.AddRoleResource(USERROLE, USER, SHOW, UPDATE)

	rbac.AddUserRoleByUsername("roporter", ADMINROLE)
	//rbac.AddUserRole(0, ADMINROLE)
	//rbac.AddUserRoleByEmail("roporter@cisco.com", ADMINROLE)
	rbac.AddUserRole(1, USERROLE)

	fmt.Println("ADMIN (USER) SHOW:       >", rbac.HasPermission(0, USER, SHOW))
	fmt.Println("ADMIN (USER) CREATE:     >", rbac.HasPermission(0, USER, CREATE))
	fmt.Println("ADMIN (USER) UDPATE:     >", rbac.HasPermission(0, USER, UPDATE))
	fmt.Println("ADMIN (USER) DELETE:     >", rbac.HasPermission(0, USER, DELETE))
	fmt.Println("ADMIN (USER) IMPORT:     >", rbac.HasPermission(0, USER, IMPORT))
	fmt.Println("ADMIN (USER) DOWNLOAD:   >", rbac.HasPermission(0, USER, DOWNLOAD))
	fmt.Println("ADMIN (USER) ADMIN:      >", rbac.HasPermission(0, USER, ADMIN))

	fmt.Println("")

	fmt.Println("ADMIN (BLOG) SHOW:       >", rbac.HasPermission(0, BLOG, SHOW))
	fmt.Println("ADMIN (BLOG) CREATE:     >", rbac.HasPermission(0, BLOG, CREATE))
	fmt.Println("ADMIN (BLOG) UDPATE:     >", rbac.HasPermission(0, BLOG, UPDATE))
	fmt.Println("ADMIN (BLOG) DELETE:     >", rbac.HasPermission(0, BLOG, DELETE))
	fmt.Println("ADMIN (BLOG) IMPORT:     >", rbac.HasPermission(0, BLOG, IMPORT))
	fmt.Println("ADMIN (BLOG) DOWNLOAD:   >", rbac.HasPermission(0, BLOG, DOWNLOAD))
	fmt.Println("ADMIN (BLOG) ADMIN:      >", rbac.HasPermission(0, BLOG, ADMIN))

	fmt.Println("")

	fmt.Println("USER (USER) SHOW:        >", rbac.HasPermission(1, USER, SHOW))
	fmt.Println("USER (USER) CREATE:      >", rbac.HasPermission(1, USER, CREATE))
	fmt.Println("USER (USER) UDPATE:      >", rbac.HasPermission(1, USER, UPDATE))
	fmt.Println("USER (USER) DELETE:      >", rbac.HasPermission(1, USER, DELETE))
	fmt.Println("USER (USER) IMPORT:      >", rbac.HasPermission(1, USER, IMPORT))
	fmt.Println("USER (USER) DOWNLOAD:    >", rbac.HasPermission(1, USER, DOWNLOAD))
	fmt.Println("USER (USER) ADMIN:       >", rbac.HasPermission(1, USER, ADMIN))

	fmt.Println("")

	fmt.Println("USER (BLOG) SHOW:        >", rbac.HasPermission(1, BLOG, SHOW))
	fmt.Println("USER (BLOG) CREATE:      >", rbac.HasPermission(1, BLOG, CREATE))
	fmt.Println("USER (BLOG) UDPATE:      >", rbac.HasPermission(1, BLOG, UPDATE))
	fmt.Println("USER (BLOG) DELETE:      >", rbac.HasPermission(1, BLOG, DELETE))
	fmt.Println("USER (BLOG) IMPORT:      >", rbac.HasPermission(1, BLOG, IMPORT))
	fmt.Println("USER (BLOG) DOWNLOAD:    >", rbac.HasPermission(1, BLOG, DOWNLOAD))
	fmt.Println("USER (BLOG) ADMIN:       >", rbac.HasPermission(1, BLOG, ADMIN))

	fmt.Println("")

	fmt.Println("OPERATION NAMES:         >", rbac.GetOperationNames())
	fmt.Println("ALL ADMIN USERS:         >", rbac.GetAllUsersWithRole("admin"))
	fmt.Println("ALL NON ADMIN USERS:     >", rbac.GetAllUsersWithRole("user"))

	fmt.Println("")

	fmt.Println("USER RESOURCE OPERATIONS:>", rbac.GetResourceOperations(USER))
	fmt.Println("BLOG RESOURCE OPERATIONS:>", rbac.GetResourceOperations(BLOG))

	fmt.Println("")

	fmt.Println("RESOURCE HAS CREATE:     >", rbac.GetResourceHasOperation(USER, CREATE))
	fmt.Println("RESOURCE HAS ADMIN:      >", rbac.GetResourceHasOperation(BLOG, ADMIN))

	fmt.Println("")

	fmt.Println("ADMIN HAS ADMIN USER:    >", rbac.GetUserHasOperation(ADMINROLE, USER, ADMIN))
	fmt.Println("ADMIN HAS ADMIN BLOG:    >", rbac.GetUserHasOperation(ADMINROLE, BLOG, ADMIN))
	fmt.Println("ADMIN HAS CREATE USER:   >", rbac.GetUserHasOperation(ADMINROLE, USER, CREATE))
	fmt.Println("ADMIN HAS CREATE BLOG:   >", rbac.GetUserHasOperation(ADMINROLE, BLOG, CREATE))
	fmt.Println("USER HAS ADMIN USER:     >", rbac.GetUserHasOperation(USERROLE, USER, ADMIN))
	fmt.Println("USER HAS ADMIN BLOG:     >", rbac.GetUserHasOperation(USERROLE, BLOG, ADMIN))
	fmt.Println("USER HAS CREATE USER:    >", rbac.GetUserHasOperation(USERROLE, USER, CREATE))
	fmt.Println("USER HAS CREATE BLOG:    >", rbac.GetUserHasOperation(USERROLE, BLOG, CREATE))

}
