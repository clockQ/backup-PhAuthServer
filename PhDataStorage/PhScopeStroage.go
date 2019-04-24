package PhDataStorage

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/PhAuthServer/PhModel"
)

// PhScopeStroage stores all of the tasty modelleaf, needs to be injected into
// Account and Account Resource. In the real world, you would use a database for that.
type PhScopeStroage struct {
	db *BmMongodb.BmMongodb
}

func (s PhScopeStroage) NewAccountStorage(args []BmDaemons.BmDaemon) *PhScopeStroage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &PhScopeStroage{db: mdb}
}

// GetAll of the modelleaf
func (s PhScopeStroage) GetAll(r api2go.Request, skip int, take int) []PhModel.Account {
	in := PhModel.Account{}
	var out []PhModel.Account
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s PhScopeStroage) GetOne(id string) (PhModel.Account, error) {
	in := PhModel.Account{ID: id}
	out := PhModel.Account{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Account for id %s not found", id)
	return PhModel.Account{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *PhScopeStroage) Insert(c PhModel.Account) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *PhScopeStroage) Delete(id string) error {
	in := PhModel.Account{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Account with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *PhScopeStroage) Update(c PhModel.Account) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Account with id does not exist")
	}

	return nil
}