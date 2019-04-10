package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/doug-martin/goqu.v5"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/mysql"
)

func main() {
	sqlDb, err := sql.Open("mysql", "fengbo:FengBo12.@tcp(127.0.0.1:3306)/Hello")
	if err != nil {
		panic(err.Error())
	}
	db := goqu.New("mysql", sqlDb)

	// DataSet
	// As
	ds := db.From("test").As("t")
	sql, _, _ := db.From(ds).ToSql()
	fmt.Println(sql)
	// ClearLimit
	ds = db.From("test").Limit(10)
	sql, _, _ = ds.ClearLimit().ToSql()
	fmt.Println(sql)
	// ClearOffset
	ds = db.From("test").Offset(2)
	sql, _, _ = ds.ClearOffset().ToSql()
	fmt.Println(sql)
	// ClearOrder
	ds = db.From("test").Order(goqu.I("a").Asc())
	sql, _, _ = ds.ClearOrder().ToSql()
	fmt.Println(sql)
	// ClearSelect
	ds = db.From("test").Select("a", "b")
	sql, _, _ = ds.ClearSelect().ToSql()
	fmt.Println(sql)
	ds = db.From("test").SelectDistinct("a", "b")
	sql, _, _ = ds.ClearSelect().ToSql()
	fmt.Println(sql)
	// ClearWhere
	ds = db.From("test").Where(
		goqu.Or(
			goqu.I("a").Gt(10),
			goqu.And(
				goqu.I("b").Lt(10),
				goqu.I("c").IsNull(),
			),
		),
	)
	sql, _, _ = ds.ClearWhere().ToSql()
	fmt.Println(sql)
	// CrossJoin
	sql, _, _ = db.From("test").CrossJoin(goqu.I("test2")).ToSql()
	fmt.Println(sql)
	sql, _, _ = db.From("test").CrossJoin(db.From("test2").Where(goqu.I("amount").Gt(0))).ToSql()
	fmt.Println(sql)
	sql, _, _ = db.From("test").CrossJoin(db.From("test2").Where(goqu.I("amount").Gt(0)).As("t")).ToSql()
	fmt.Println(sql)
	// From
	sql, _, _ = db.From("test").ToSql()
	fmt.Println(sql)

	ds = db.From("test")
	fromDs := ds.Where(goqu.I("age").Gt(10))
	sql, _, _ = ds.From(fromDs.As("test2")).ToSql()
	fmt.Println(sql)

	ds = db.From("test")
	fromDs = ds.Where(goqu.I("age").Gt(10))
	sql, _, _ = ds.From(fromDs).ToSql()
	fmt.Println(sql)

	// FromSelf
	sql, _, _ = db.From("test").FromSelf().ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").As("my_test_table").FromSelf().ToSql()
	fmt.Println(sql)

	// FullJoin
	sql, _, _ = db.From("test").FullJoin(goqu.I("test2"), goqu.On(goqu.Ex{"test.fkey": goqu.I("test2.Id")})).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").FullJoin(goqu.I("test2"), goqu.Using("common_column")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").FullJoin(db.From("test2").Where(goqu.I("amount").Gt(0)), goqu.On(goqu.I("test.fkey").Eq(goqu.I("test2.Id")))).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").FullJoin(db.From("test2").Where(goqu.I("amount").Gt(0)).As("t"), goqu.On(goqu.I("test.fkey").Eq(goqu.I("t.Id")))).ToSql()
	fmt.Println(sql)

	// FullOuterJoin
	sql, _, _ = db.From("test").FullOuterJoin(goqu.I("test2"), goqu.On(goqu.Ex{"test.fkey": goqu.I("test2.Id")})).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").FullOuterJoin(goqu.I("test2"), goqu.Using("common_column")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").FullOuterJoin(db.From("test2").Where(goqu.I("amount").Gt(0)), goqu.On(goqu.I("test.fkey").Eq(goqu.I("test2.Id")))).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").FullOuterJoin(db.From("test2").Where(goqu.I("amount").Gt(0)).As("t"), goqu.On(goqu.I("test.fkey").Eq(goqu.I("t.Id")))).ToSql()
	fmt.Println(sql)

	// Having
	sql, _, _ = db.From("test").Having(goqu.SUM("income").Gt(1000)).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").GroupBy("age").Having(goqu.SUM("income").Gt(1000)).ToSql()
	fmt.Println(sql)

	// InnerJoin
	sql, _, _ = db.From("test").InnerJoin(goqu.I("test2"), goqu.On(goqu.Ex{"test.fkey": goqu.I("test2.Id")})).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").InnerJoin(goqu.I("test2"), goqu.Using("common_column")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").InnerJoin(db.From("test2").Where(goqu.I("amount").Gt(0)), goqu.On(goqu.I("test.fkey").Eq(goqu.I("test2.Id")))).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").InnerJoin(db.From("test2").Where(goqu.I("amount").Gt(0)).As("t"), goqu.On(goqu.I("test.fkey").Eq(goqu.I("t.Id")))).ToSql()
	fmt.Println(sql)

	// Intersect
	sql, _, _ = db.From("test").Intersect(db.From("test2")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Limit(1).Intersect(db.From("test2")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Limit(1).Intersect(db.From("test2").Order(goqu.I("id").Desc())).ToSql()
	fmt.Println(sql)

	// Select
	sql, _, _ = db.From("test").Select("a", "b", "c").ToSql()
	fmt.Println(sql)

	ds = db.From("test")
	fromDs = ds.Select("age").Where(goqu.I("age").Gt(10))
	sql, _, _ = ds.From().Select(fromDs.As("ages")).ToSql()
	fmt.Println(sql)

	ds = db.From("test")
	fromDs = ds.Select("age").Where(goqu.I("age").Gt(10))
	sql, _, _ = ds.From().Select(fromDs).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Select(goqu.L("a + b").As("sum")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Select(
		goqu.COUNT("*").As("age_count"),
		goqu.MAX("age").As("max_age"),
		goqu.AVG("age").As("avg_age"),
	).ToSql()
	fmt.Println(sql)

	ds = db.From("test")
	type myStruct struct {
		Name         string
		Address      string `db:"address"`
		EmailAddress string `db:"email_address"`
	}
	sql, _, _ = ds.Select(&myStruct{}).ToSql()
	fmt.Println(sql)

	sql, _, _ = ds.Select(myStruct{}).ToSql()
	fmt.Println(sql)

	type myStruct2 struct {
		myStruct
		Zipcode string `db:"zipcode"`
	}

	sql, _, _ = ds.Select(&myStruct2{}).ToSql()
	fmt.Println(sql)

	sql, _, _ = ds.Select(myStruct2{}).ToSql()
	fmt.Println(sql)

	var myStructs []myStruct
	sql, _, _ = ds.Select(myStructs).ToSql()
	fmt.Println(sql)

	// type Ex
	//
	sql, args, _ := db.From("items").Where(
		goqu.Ex{
			"col1": "a",
			"col2": 1,
			"col3": true,
			"col4": false,
			"col5": nil,
			"col6": []string{"a", "b", "c"},
		}).ToSql()
	fmt.Println(sql, args)

	sql, args, _ = db.From("items").Prepared(true).Where(goqu.Ex{
		"col1": "a",
		"col2": 1,
		"col3": true,
		"col4": false,
		"col5": []string{"a", "b", "c"},
	}).ToSql()
	fmt.Println(sql, args)

	sql, _, _ = db.From("items").Where(goqu.Ex{
		"col1": goqu.Op{"neq": "a"},
		"col3": goqu.Op{"isNot": true},
		"col6": goqu.Op{"notIn": []string{"a", "b", "c"}},
	}).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("items").Where(goqu.Ex{
		"col1": goqu.Op{"gt": 1},
		"col2": goqu.Op{"gte": 1},
		"col3": goqu.Op{"lt": 1},
		"col4": goqu.Op{"lte": 1},
	}).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("items").Where(goqu.Ex{
		"col1": goqu.Op{"like": "a%"},
		"col2": goqu.Op{"notLike": "a%"},
		"col3": goqu.Op{"iLike": "a%"},
		"col4": goqu.Op{"notILike": "a%"},
	}).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("items").Where(goqu.Ex{
		"col1": goqu.Op{"like": regexp.MustCompile("^(a|b)")},
		"col2": goqu.Op{"notLike": regexp.MustCompile("^(a|b)")},
		"col3": goqu.Op{"iLike": regexp.MustCompile("^(a|b)")},
		"col4": goqu.Op{"notILike": regexp.MustCompile("^(a|b)")},
	}).ToSql()
	fmt.Println(sql)

	// type ExOr
	//
	sql, _, _ = db.From("items").Where(goqu.ExOr{
		"col1": "a",
		"col2": 1,
		"col3": true,
		"col4": false,
		"col5": nil,
		"col6": []string{"a", "b", "c"},
	}).ToSql()
	fmt.Println(sql)

	// func I
	sql, _, _ = db.From("test").Where(
		goqu.I("a").Eq(10),
		goqu.I("b").Lt(10),
		goqu.I("d").IsTrue(),
	).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From(goqu.I("test").Schema("my_schema")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From(goqu.I("mychema.test")).Where(
		goqu.I("my_schema.test.a").Eq(10),
	).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From(goqu.I("test")).Select(goqu.I("test.*")).ToSql()
	fmt.Println(sql)

	//InMethod
	sql, _, _ = db.From("test").Where(goqu.I("a").In("a", "b", "c")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Where(goqu.I("a").In([]string{"a", "b", "c"})).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Where(goqu.I("a").NotIn("a", "b", "c")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Where(goqu.I("a").NotIn([]string{"a", "b", "c"})).ToSql()
	fmt.Println(sql)

	// using an Ex expression map
	sql, _, _ = db.From("test").Where(goqu.Ex{
		"a": []string{"a", "b", "c"},
	}).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Where(goqu.Ex{
		"a": goqu.Op{"notIn": []string{"a", "b", "c"}},
	}).ToSql()
	fmt.Println(sql)

	// func L
	sql, _, _ = db.From("test").Where(goqu.L("a = 1")).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Where(goqu.L("a = 1 AND (b = ? OR ? = ?)", "a", goqu.I("c"), 0.01)).ToSql()
	fmt.Println(sql)

	sql, _, _ = db.From("test").Where(
		goqu.L(
			"(? AND ?) OR ?",
			goqu.I("a").Eq(1),
			goqu.I("b").Eq("b"),
			goqu.I("c").In([]string{"a", "b", "c"}),
		),
	).ToSql()
	fmt.Println(sql)

}
