package postgreSQL

import (
	"context"
	"fmt"
	"gitlab.com/package7225033/go-modules/check"
	"gitlab.com/package7225033/go-modules/conv"
	"gitlab.com/package7225033/go-modules/logs"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type Model struct {
	ID        bson.ObjectID `bson:"_id" json:"id" form:"id"`
	CreatedAt time.Time     `json:"created_at" form:"created_at" bson:"index"`
	UpdatedAt time.Time     `json:"updated_at" form:"updated_at" bson:"index"`

	//Pagination and Filter
	Pagination Pagination `json:"-" form:"-" bson:"-"`
	//Practice
	IsTest bool `json:"-" form:"tutorial" bson:"-"`

	//ILike cond
	ILike   []string `json:"-" bson:"-"`
	Preload []string `json:"-" bson:"-"`

	DISTINCT string `json:"-" form:"-" bson:"-"`

	//Connection
	Config               Config          `json:"-" form:"-" bson:"-"`
	conn                 *mongo.Database `json:"-" form:"-" bson:"-"`
	client               *mongo.Client   `json:"-" form:"-" bson:"-"`
	debug                bool            `json:"-" form:"-" bson:"-"`
	joins                [][]interface{} `json:"-" form:"-" bson:"-"`
	collecNameAdditional string          `json:"-" form:"-" bson:"-"`
}

func (m *Model) CollectionNameAdditional(str string) *Model {
	m.collecNameAdditional = str
	return m
}
func (m *Model) connect() (err error) {
	MONGODB_URI := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", m.Config.USER, m.Config.PASS, m.Config.CLUSTER_URL)

	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Defines the options for the MongoDB client
	opts := options.Client().
		ApplyURI(MONGODB_URI).
		SetServerAPIOptions(serverAPI).
		SetMinPoolSize(uint64(m.Config.MIN_POOL_SIZE)).
		SetMaxPoolSize(uint64(m.Config.MAX_POOL_SIZE)).
		SetTimeout(conv.IntToTime(m.Config.MAX_TIME_OUT_CONNS) * time.Millisecond)

	// Creates a new client and connects to the server
	if m.client, err = mongo.Connect(opts); err != nil {
		logs.ErrorHandler(err)
		return
	}

	m.conn = m.client.Database(m.Config.NAME)

	return
}

func (m *Model) Read(data interface{}, filter bson.M) (err error) {
	err = m.connect()
	if err != nil {
		panic(err)
	}

	if check.SliceStruct(data) {
		var cursor *mongo.Cursor
		cursor, err = m.conn.Collection(conv.NameStructToStringSnakeCase(data, m.collecNameAdditional)).Find(context.TODO(), m.filters(&[]bson.M{filter}, nil))
		if err != nil {
			logs.ErrorHandler(err)
			return
		}

		if err = cursor.All(context.TODO(), data); err != nil {
			logs.ErrorHandler(err)
			return
		}
	} else {
		singleResult := m.conn.Collection(conv.NameStructToStringSnakeCase(data, m.collecNameAdditional)).FindOne(context.TODO(), filter)
		if err = singleResult.Decode(data); err != nil {
			logs.ErrorHandler(err)
			return
		}
	}
	return
}

func (m *Model) Create(data interface{}, filter bson.M) (err error) {
	err = m.connect()
	if err != nil {
		panic(err)
	}

	if check.SliceStruct(data) {
		//var cursor *mongo.InsertManyResult
		_, err = m.conn.Collection(conv.NameStructToStringSnakeCase(data, m.collecNameAdditional)).InsertMany(context.TODO(), data)
		if err != nil {
			logs.ErrorHandler(err)
			return
		}

	} else {
		//var single *mongo.InsertOneResult
		_, err = m.conn.Collection(conv.NameStructToStringSnakeCase(data, m.collecNameAdditional)).InsertOne(context.TODO(), data)
		if err != nil {
			logs.ErrorHandler(err)
			return
		}
	}
	return
}

func (m Model) filters(filter *[]bson.M, iLike *[]string) bson.M {
	m.Pagination.collectDefault()
	var (
		filterData = bson.M{}
		filterAnd  = bson.A{}
	)

	if filter != nil && len(*filter) > 0 {
		for _, elm := range *filter {
			filterAnd = append(filterAnd, elm)
		}

	}

	//filter search
	if m.Pagination.Keyword != "" {
		var (
			filterKeyword  bson.A
			filterKeywordM bson.M
		)
		if iLike != nil && len(*iLike) > 0 {
			for _, elm := range *iLike {
				filterKeyword = append(filterKeyword, bson.M{elm: bson.M{"$regex": m.Pagination.Keyword, "$options": "i"}})
			}

			filterKeywordM = bson.M{"$or": filterKeyword}
			filterAnd = append(filterAnd, filterKeywordM)
		}

	}

	//filter date
	if m.Pagination.StartDate != "" && m.Pagination.EndDate != "" {
		dateFilter := bson.M{}
		timeParseStartDate, _ := time.Parse("2006-01-02 15:04:05", m.Pagination.StartDate+" 00:00:00")
		timeParseEndDate, _ := time.Parse("2006-01-02 15:04:05", m.Pagination.EndDate+" 23:59:59")
		dateFilter["$gte"] = timeParseStartDate
		dateFilter["$lte"] = timeParseEndDate

		if len(filterAnd) > 0 {
			filterAnd = append(filterAnd, bson.M{"created_at": dateFilter})
		} else {
			if len(dateFilter) > 0 {
				filterData["created_at"] = dateFilter
			}
		}
	}

	if len(filterAnd) > 1 {
		filterData = bson.M{"$and": filterAnd}
	} else {
		if len(filterAnd) > 0 && (filter != nil && len(*filter) == 1) {
			newFilter := *filter
			return newFilter[0]
		}
	}

	return filterData
}
