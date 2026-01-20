package services
import(
	"fmt"
	"time"
	"NotesApp/utilities"
	"NotesApp/models"
	"errors"

)
func CreateList(name string) (models.List, error){
	lists, err := utilities.LoadLists()
	if err !=nil{
		return models.List{},err
	}
	newid := utilities.NextID(utilities.AsIdentifiable(lists))

	list := models.List{
		ID: newid,
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	
	lists = append(lists, list)

	if err := utilities.SaveLists(lists); err != nil {
		return models.List{},err
	}
	return list,nil
}
func GetLists()([]models.List, error){
	lists, err := utilities.LoadLists()
    if err != nil {
        return nil,err
    }
    return lists, nil
}
func DeleteList(id int)error{
	lists, err := utilities.LoadLists()
	if err !=nil{
		return err
	}
	var newlists []models.List
	found :=false
	for _,list:=range lists{
		if list.ID==id{
			found = true
		}else{
			newlists=append(newlists, list)
		}
	}
	if !found {
        return fmt.Errorf("no list found with ID %d", id)
    }
	if err := utilities.SaveLists(newlists); err != nil {
		return err
	}
	return nil
}
func AddItem(listID int, text string) (models.Item, error) {
	lists, err := utilities.LoadLists()
	if err != nil {
		return models.Item{}, err
	}

	var target *models.List
	for i := range lists {
		if lists[i].ID == listID {
			target = &lists[i]
			break
		}
	}

	if target == nil {
		return models.Item{}, errors.New("list not found")
	}

	nextID := utilities.NextItemID(target.Items)

	item := models.Item{
		ID:        nextID,
		Text:      text,
		Check:     false,
		CreatedAt: time.Now(),
	}

	target.Items = append(target.Items, item)

	now := time.Now()
	target.UpdatedAt = &now

	if err := utilities.SaveLists(lists); err != nil {
		return models.Item{}, err
	}

	return item, nil
}
func RemoveItem(listID int, itemID int) error {
	lists, err := utilities.LoadLists()
	if err != nil {
		return err
	}

	var target *models.List
	for i := range lists {
		if lists[i].ID == listID {
			target = &lists[i]
			break
		}
	}

	if target == nil {
		return errors.New("list not found")
	}

	found := false
	for i, item := range target.Items {
		if item.ID == itemID {
			target.Items = append(target.Items[:i], target.Items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("item not found")
	}

	now := time.Now()
	target.UpdatedAt = &now

	return utilities.SaveLists(lists)
}

