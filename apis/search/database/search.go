package database

import (
	"context"
	"golang_globalsearch_new/apis/search/model"
	"golang_globalsearch_new/initializers"
	"log"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchDB interface {
	CreateObjects() error
	SearchObject(searchWord string) ([]model.Object, error)
	ShowObject(objectID string)(model.Object,error)
}

type searchDB struct {
	Mgr *initializers.Manager
}

func NewSearchDB(Mgr *initializers.Manager) SearchDB {
	return &searchDB{
		Mgr: Mgr,
	}
}

func (s *searchDB) CreateObjects() error {
	objSlice := []model.Object{}
	for i := 0; i < 500; i++ {
		object := GenerateObj()
		objSlice = append(objSlice, object)
	}
	slice := convertToInterfaceSlice(objSlice)
	collection := s.Mgr.Connection.Database("new_global_search").Collection("new_search_objects")
	_, err := collection.InsertMany(context.TODO(), slice)

	return err
}

func (s *searchDB) SearchObject(searchword string) ([]model.Object, error) {
	collection := s.Mgr.Connection.Database("new_global_search").Collection("new_search_objects")
	query := bson.M{"$or": []interface{}{
		bson.M{"visible": bson.M{"$regex": searchword, "$options": "i"}},
		bson.M{"tags": bson.M{"$regex": searchword, "$options": "i"}},
		bson.M{"desctiption": bson.M{"$regex": searchword, "$options": "i"}},
		bson.M{"title": bson.M{"$regex": searchword, "$options": "i"}},
		bson.M{"type": bson.M{"$regex": searchword, "$options": "i"}},
	}}
	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	var objects []model.Object

	if scanErr := cursor.All(context.TODO(), &objects); scanErr != nil {
		log.Fatal(scanErr)
	}
	return objects, nil
}

func(s *searchDB)ShowObject(ObjectID string)(model.Object,error){
	objectId,err:=primitive.ObjectIDFromHex(ObjectID)
	var object model.Object
	if err!=nil{
		log.Fatal(err)
	}
	collection:=s.Mgr.Connection.Database("new_global_search").Collection("new_search_objects")
	query:=bson.M{"_id":objectId}
	resErr:=collection.FindOne(context.TODO(),query).Decode(&object)
	return object,resErr
	
}

func GenerateObj() model.Object {

	objID := primitive.NewObjectID()
	boardID := primitive.NewObjectID()
	date := GenerateRandomDate()
	obj := model.Object{
		ID: objID,
		BoardID: model.Board{
			ID: boardID,
		},
		Visible:     AssignRandomVisibility(),
		Tags:        AssignRandomTags(),
		Description: AssignRandomDescription(boardID.Hex(), date),
		Title:       AssignRandomTitle(date),
		Type:        AssignRandomType(),
	}
	return obj
}

func AssignRandomType() string {
	types := []string{"BOARD", "FILE", "NOTE"}

	rand.Seed(time.Now().Unix())
	return types[rand.Intn(len(types))]
}

func AssignRandomVisibility() string {
	Visibilities := []string{"PUBLIC", "PRIVATE", "CONTACTS"}

	rand.Seed(time.Now().Unix())
	return Visibilities[rand.Intn(len(Visibilities))]
}

func AssignRandomDescription(boardID, date string) string {
	// startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	// endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	// duration := endDate.Sub(startDate)

	// randomDuration := time.Duration(rand.Int63n(int64(duration)))
	// randomDate := startDate.Add(randomDuration)

	// dateStr := randomDate.Format("Jan 02,2006")
	// dateStr = strings.ReplaceAll(dateStr, "2023", "")

	ans := date + ":" + boardID
	return ans

}

func GenerateRandomDate() string {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	duration := endDate.Sub(startDate)

	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	randomDate := startDate.Add(randomDuration)

	dateStr := randomDate.Format("Jan 02,2006")
	dateStr = strings.ReplaceAll(dateStr, "2023", "")
	return dateStr
}

func AssignRandomTitle(date string) string {
	return "Board:" + date
}

func AssignRandomTags() []string {
	ans := []string{}
	tags := []string{"Ocean", "Water", "IsLand", "Vacation", "Enjoyment", "Fun", "Chill", "Fire", "Cool", "Nice"}

	for i := 0; i < rand.Intn(10); i++ {
		ans = append(ans, tags[rand.Intn(10)])
	}
	return ans
}

func convertToInterfaceSlice(objSlice []model.Object) []interface{} {

	ans := []interface{}{}

	for _, val := range objSlice {
		ans = append(ans, val)
	}
	return ans
}
