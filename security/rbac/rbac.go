package rbac

import (
	"errors"
	"fmt"
	"strings"

	"github.com/robjporter/go-functions/as"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type RoleUser struct {
	Role Role
	User User
}

type Role struct {
	Id   int
	Name string
}

type RoleResource struct {
	Role     Role
	Resource Resource
	Value    int
}

type Resource struct {
	Id   int
	Name string
}

type ResourceOperation struct {
	Resource  Resource
	Operation Operation
}

type Operation struct {
	Id    int
	Name  string
	Value int
}

var (
	users              []User
	roles              []Role
	resources          []Resource
	operations         []Operation
	roleUsers          []RoleUser
	resourceOperations []ResourceOperation
	roleResources      []RoleResource
)

func newUserId() int {
	return len(users) + 1
}

func newRoleId() int {
	return len(roles) + 1
}

func newResourceId() int {
	return len(resources) + 1
}

func newOperationId() int {
	return len(operations) + 1
}

//INTERNAL///////////////////////////////////////////////////////////////////
func checkDistinctUsernameKey(data string) bool {
	toReturn := false
	for i := 0; i < len(users); i++ {
		if users[i].Name == strings.ToLower(data) {
			return true
		}
	}
	return toReturn
}

func checkDistinctRoleKey(data string) bool {
	toReturn := false
	for i := 0; i < len(roles); i++ {
		if roles[i].Name == strings.ToLower(data) {
			return true
		}
	}
	return toReturn
}

func checkDistinctResourceKey(data string) bool {
	toReturn := false
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == strings.ToLower(data) {
			return true
		}
	}
	return toReturn
}

func checkDistinctOperationKey(data string) bool {
	toReturn := false
	for i := 0; i < len(operations); i++ {
		if operations[i].Name == strings.ToLower(data) {
			return true
		}
	}
	return toReturn
}

func addResourceOperation(res Resource, ops Operation) ResourceOperation {
	return ResourceOperation{res, ops}
}

func addRoleResource(role Role, res Resource, ops []Operation) RoleResource {
	num := 0
	for _, op := range ops {
		num += op.Value
	}
	return RoleResource{
		Role:     role,
		Resource: res,
		Value:    num,
	}
}

func assignRole(user User, role Role) RoleUser {
	return RoleUser{
		User: user,
		Role: role,
	}
}

//GET POSITION///////////////////////////////////////////////////////////////////
func GetOperationPosition(name string) int {
	name = strings.ToLower(name)
	for i := 0; i < len(operations); i++ {
		if operations[i].Name == name {
			return operations[i].Id - 1
		}
	}
	return -1
}

func GetRolePosition(name string) int {
	name = strings.ToLower(name)
	for i := 0; i < len(roles); i++ {
		if roles[i].Name == name {
			return roles[i].Id - 1
		}
	}
	return -1
}

func GetResourcePosition(name string) int {
	name = strings.ToLower(name)
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == name {
			return resources[i].Id - 1
		}
	}
	return -1
}

func GetUserPositionByUsername(name string) int {
	name = strings.ToLower(name)
	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			return users[i].Id - 1
		}
	}
	return -1
}

func GetUserPositionByEmail(email string) int {
	email = strings.ToLower(email)
	for i := 0; i < len(operations); i++ {
		if users[i].Email == email {
			return operations[i].Id - 1
		}
	}
	return -1
}

//GET ID///////////////////////////////////////////////////////////////////
func GetOperationID(name string) int {
	for i := 0; i < len(operations); i++ {
		if operations[i].Name == name {
			return operations[i].Id
		}
	}
	return -1
}

func GetRoleID(name string) int {
	for i := 0; i < len(roles); i++ {
		if roles[i].Name == name {
			return roles[i].Id
		}
	}
	return -1
}

func GetResourceID(name string) int {
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == name {
			return resources[i].Id
		}
	}
	return -1
}

func GetUserIDByUsername(name string) int {
	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			return users[i].Id
		}
	}
	return -1
}

func GetUserIDByEmail(email string) int {
	for i := 0; i < len(operations); i++ {
		if users[i].Email == email {
			return operations[i].Id
		}
	}
	return -1
}

//GET OBJECT///////////////////////////////////////////////////////////////////
func getUserByID(userNumber int) User {
	if userNumber <= len(users) {
		return users[userNumber]
	}
	return User{}
}

func getUserByName(username string) User {
	username = strings.ToLower(username)
	for i := 0; i < len(users); i++ {
		if users[i].Name == username {
			return users[i]
		}
	}
	return User{}
}

func getUserByEmail(email string) User {
	email = strings.ToLower(email)
	for i := 0; i < len(users); i++ {
		if users[i].Email == email {
			return users[i]
		}
	}
	return User{}
}

func getResourceByID(resourceNumber int) Resource {
	if resourceNumber <= len(resources) {
		return resources[resourceNumber]
	}
	return Resource{}
}

func getResourceByName(resourceName string) Resource {
	resourceName = strings.ToLower(resourceName)
	if resourceName != "" {
		res := GetResourceID(resourceName)
		if res > -1 {
			return resources[res]
		}
	}
	return Resource{}
}

func getOperationByID(operationNumber int) Operation {
	if operationNumber <= len(operations) {
		return operations[operationNumber]
	}
	return Operation{}
}

func getOperationByName(operationName string) Operation {
	operationName = strings.ToLower(operationName)
	if operationName != "" {
		ops := GetOperationID(operationName)
		if ops > -1 {
			return operations[ops]
		}
	}
	return Operation{}
}

//ADD///////////////////////////////////////////////////////////////////
func AddUser(name string, email string) (error, bool) {
	if !checkDistinctUsernameKey(name) {
		users = append(users, User{Id: newUserId(), Name: strings.ToLower(name), Email: strings.ToLower(email)})
		return nil, true
	}
	return errors.New(name + " User already exists"), false
}

func AddRole(name string) (error, bool) {
	if !checkDistinctRoleKey(name) {
		roles = append(roles, Role{Id: newRoleId(), Name: name})
		return nil, true
	}
	return errors.New(name + " Role already exists"), false
}

func AddResource(name string) (error, bool) {
	if !checkDistinctResourceKey(name) {
		resources = append(resources, Resource{Id: newResourceId(), Name: name})
		return nil, true
	}
	return errors.New(name + " Resource already exists"), false
}

func AddOperation(name string) (error, bool) {
	if !checkDistinctOperationKey(name) {
		id := newOperationId()
		if id == 1 {
			operations = append(operations, Operation{Id: id, Name: name, Value: 1})
		} else {
			operations = append(operations, Operation{Id: id, Name: name, Value: 1 << uint(id)})
		}
		return nil, true
	}
	return errors.New(name + " Operation already exists"), false
}

func AddResourceOperationByName(resource string, operationnames ...string) (error, bool) {
	res := GetResourcePosition(resource)
	if res > -1 {
		for i := 0; i < len(operationnames); i++ {
			ops := GetOperationPosition(operationnames[i])
			if ops > -1 {
				resourceOperations = append(resourceOperations, addResourceOperation(resources[res], operations[ops]))
			}
		}
	} else {
		return errors.New("Uknown resource name."), false
	}
	return errors.New(""), false
}

func AddResourceOperation(resourceNumber int, operationids ...int) (error, bool) {
	if resourceNumber <= len(resources) {
		tmp := []Operation{}
		for i := 0; i < len(operationids); i++ {
			tmp = append(tmp, operations[i])
			resourceOperations = append(resourceOperations, addResourceOperation(resources[resourceNumber], operations[i]))
		}
		return nil, true
	} else {
		return errors.New("The resource " + as.ToString(resourceNumber) + " does not exist."), false
	}
	return errors.New("An unknown error has occured and the resource operation was not added."), false
}

func AddRoleResource(roleNumber int, resourceNumber int, operationids ...int) (error, bool) {
	//fmt.Println("ROLENUMBER: >", roleNumber, "RESOURCENUMBER: >", resourceNumber, "OPERATIONS: >", operationids)
	if roleNumber <= len(roles) && resourceNumber <= len(resources) {
		tmp := []Operation{}
		for i := 0; i < len(operationids); i++ {
			tmp = append(tmp, operations[operationids[i]])
		}
		roleResources = append(roleResources, addRoleResource(roles[roleNumber], resources[resourceNumber], tmp))
		return nil, true
	} else {
		return errors.New("The resource " + as.ToString(resourceNumber) + " does not exist."), false
	}
	return errors.New("An unknown error has occured and the role resource was not added."), false
}

func AddUserRole(userNumber int, roleNumber int) (error, bool) {
	if userNumber <= len(users) && roleNumber <= len(roles) {
		roleUsers = append(roleUsers, assignRole(users[userNumber], roles[roleNumber]))
		return nil, true
	} else {
		return errors.New("The role " + as.ToString(roleNumber) + " does not exist."), false
	}
	return errors.New("An unknown error has occured and the user number was not added."), false
}

func AddUserRoleByUsername(username string, roleNumber int) (error, bool) {
	if username != "" && roleNumber <= len(roles) {
		u := getUserByName(username)
		pos := u.Id - 1
		return AddUserRole(pos, roleNumber)
	}
	return nil, true
}

func AddUserRoleByEmail(email string, roleNumber int) (error, bool) {
	if email != "" && roleNumber <= len(roles) {
		u := getUserByEmail(email)
		pos := u.Id - 1
		return AddUserRole(pos, roleNumber)
	}
	return nil, true
}

//PERMISSIONS///////////////////////////////////////////////////////////////////
func HasPermission(user2 int, res2 int, op2 int) bool {
	if user2 <= len(users) && res2 <= len(resources) && op2 <= len(operations) {
		user := getUserByID(user2)
		res := getResourceByID(res2)
		op := getOperationByID(op2)
		role := findRole(user)
		roleResource := findRoleResource(role, res)

		if op.Value&roleResource.Value == 0 {
			return false
		}
		return true
	}
	return false
}

func findRole(user User) Role {
	for _, ru := range roleUsers {
		if ru.User.Id == user.Id {
			return ru.Role
		}
	}
	return Role{}
}

func findRoleResource(role Role, res Resource) RoleResource {
	for _, rr := range roleResources {
		if rr.Role.Id == role.Id && rr.Resource.Id == res.Id {
			return rr
		}
	}
	return RoleResource{}
}

//OUTPUT///////////////////////////////////////////////////////////////////
func GetOperationNames() []string {
	tmp := []string{}
	for i := 0; i < len(operations); i++ {
		tmp = append(tmp, operations[i].Name)
	}
	return tmp
}

func GetAllUsersWithRole(role string) []string {
	tmp := []string{}
	for i := 0; i < len(roleUsers); i++ {
		if roleUsers[i].Role.Name == role {
			tmp = append(tmp, roleUsers[i].User.Name)
		}
	}
	return tmp
}

func GetResourceOperations(re int) []string {
	ops := []string{}
	if re <= len(resources) {
		res := getResourceByID(re)
		for i := 0; i < len(resourceOperations); i++ {
			if res == resourceOperations[i].Resource {
				ops = append(ops, resourceOperations[i].Operation.Name)
			}
		}
	}
	return ops
}

func GetResourceHasOperation(re int, op int) bool {
	if re <= len(resources) && op <= len(operations) {
		res := getResourceByID(re)
		ops := getOperationByID(op)

		for i := 0; i < len(resourceOperations); i++ {
			if res == resourceOperations[i].Resource {
				if ops == resourceOperations[i].Operation {
					return true
				}
			}
		}
	}
	return false
}

func GetUserHasOperation(us int, resource int, op int) bool {
	if us <= len(users) && op <= len(operations) && resource <= len(resources) {
		user := getUserByID(us)
		ops := getOperationByID(op)
		role := findRole(user)
		reso := getResourceByID(resource)
		res := findRoleResource(role, reso)

		if ops.Value&res.Value == 0 {
			return false
		} else {
			return true
		}
	}
	return false
}

//DEBUG///////////////////////////////////////////////////////////////////
func DebugUsers() {
	fmt.Println("DEBUG*USERS***********************************************")
	for i := 0; i < len(users); i++ {
		fmt.Println("USER:               >", users[i])
	}
}

func DebugRoles() {
	fmt.Println("DEBUG*ROLE*************************************************")
	for i := 0; i < len(roles); i++ {
		fmt.Println("ROLE:               >", roles[i])
	}
}

func DebugResources() {
	fmt.Println("DEBUG*RESOURCE*********************************************")
	for i := 0; i < len(resources); i++ {
		fmt.Println("RESOURCE:           >", resources[i])
	}
}

func DebugOperations() {
	fmt.Println("DEBUG*OPERATIONS*******************************************")
	for i := 0; i < len(operations); i++ {
		fmt.Println("OPERATIONS          >", operations[i])
	}
}

func DebugResourceOperations() {
	fmt.Println("DEBUG*RESOURCEOPERATIONS***********************************")
	for i := 0; i < len(resourceOperations); i++ {
		fmt.Println("RESOURCEOPERATIONS  >", resourceOperations[i])
	}
}

func DebugRoleResources() {
	fmt.Println("DEBUG*ROLERESOURCES****************************************")
	for i := 0; i < len(roleResources); i++ {
		fmt.Println("ROLERESOURCES       >", roleResources[i])
	}
}

func DebugRoleUsers() {
	fmt.Println("DEBUG*ROLEUSERS********************************************")
	for i := 0; i < len(roleUsers); i++ {
		fmt.Println("ROLEUSERS           >", roleUsers[i])
	}
}

func Debug() {
	fmt.Println("USERS:              >", users)
	fmt.Println("ROLES:              >", roles)
	fmt.Println("RESOURCES:          >", resources)
	fmt.Println("OPERATIONS:         >", operations)
	fmt.Println("RESOURCEOPERATIONS: >", resourceOperations)
	fmt.Println("ROlERESOURCES:      >", roleResources)
	fmt.Println("ROLEUSERS:          >", roleUsers)
}
